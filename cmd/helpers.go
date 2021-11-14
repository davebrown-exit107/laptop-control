package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
	"mrogalski.eu/go/pulseaudio"
)

// Brightness is a struct to support simplifying communication of brightness
// information
type Brightness struct {
	notifyMsg     notify.Notification
	stringMsg     string
	percent       int
	maxBrightness int
	curBrightness int
}

// Volume is a struct to support simplifying communication of volume information
type Volume struct {
	notifyMsg notify.Notification
	stringMsg string
	muted     bool
	curVolume float32
}

// GetCurBrightness is a helper function that returns the current brighness in
// the form of a Brightness struct
func GetCurBrightness() Brightness {
	var current = Brightness{}
	var iconName = "/usr/share/icons/Adwaita/96x96/status/display-brightness-symbolic.symbolic.png"
	sysBasePath := "/sys/class/backlight/intel_backlight"

	curBrightness, err := ReadFileToInt(path.Join(sysBasePath, "brightness"))
	if err != nil {
		log.Println(err)
	}

	maxBrightness, err := ReadFileToInt(path.Join(sysBasePath, "max_brightness"))
	if err != nil {
		log.Println(err)
	}

	current.curBrightness = curBrightness
	current.maxBrightness = maxBrightness
	current.percent = int((float32(curBrightness) / float32(maxBrightness)) * 100)
	current.stringMsg = fmt.Sprintf("Brightness: %d%%", current.percent)

	current.notifyMsg = notify.Notification{
		AppName:       "local-control",
		ReplacesID:    uint32(0),
		AppIcon:       iconName,
		Summary:       "screen",
		Body:          fmt.Sprintf("%d%%", int(current.percent)),
		Actions:       []notify.Action{},
		Hints:         map[string]dbus.Variant{},
		ExpireTimeout: time.Second * 5,
	}

	return current
}

// ReadFileToInt is a helper function that takes a string argument of a file
// in the /sys path and returns the contents of it as an integer. This
// is mostly used to read in things like max_brightness or brightness
func ReadFileToInt(file string) (int, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}

	dataTrim := strings.TrimSpace(string(data))
	dataInt, err := strconv.Atoi(dataTrim)
	if err != nil {
		return 0, err
	}

	return dataInt, nil
}

// ConvertToFloat is a helper function that takes a string argument and
// converts it to a float between 0 and 1.
func ConvertToFloat(num string) (float32, error) {
	intVal, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		return 0.0, err
	}

	if intVal < 0 || intVal > 100 {
		err := fmt.Errorf("Please provide a number between 0 and 100")
		return 0.0, err
	}

	floatVal := float32(intVal) / 100
	return floatVal, nil

}

// GetCurrentVolume is a helper function that takes the current volume state
// and presents it in a number of helpful formats to consolidate formatting
// and presentation
func GetCurrentVolume(client *pulseaudio.Client) (notify.Notification, string, bool, float32, error) {
	curVolume, err := client.Volume()
	if err != nil {
		return notify.Notification{}, "", false, 0.0, err
	}

	muted, err := client.Mute()
	if err != nil {
		return notify.Notification{}, "", false, 0.0, err
	}

	var volumeStr string
	if muted {
		volumeStr = fmt.Sprintf("Muted: %d", int(curVolume*100))
	} else {
		volumeStr = fmt.Sprintf("Volume: %d", int(curVolume*100))
	}

	var icon []string
	icon = append(icon, "/usr/share/icons/Adwaita/96x96/status/audio-volume-")

	if muted {
		icon = append(icon, "muted")
	} else {
		switch {
		case curVolume == 0:
			icon = append(icon, "muted")
		case curVolume < 0.2:
			icon = append(icon, "low")
		case curVolume < 0.7:
			icon = append(icon, "medium")
		default:
			icon = append(icon, "high")
		}
	}

	icon = append(icon, "-symbolic.symbolic.png")

	iconName := strings.Join(icon, "")

	volumeNotify := notify.Notification{
		AppName:       "local-control",
		ReplacesID:    uint32(0),
		AppIcon:       iconName,
		Summary:       "volume",
		Body:          fmt.Sprintf("%d%%", int(curVolume*100)),
		Actions:       []notify.Action{},
		Hints:         map[string]dbus.Variant{},
		ExpireTimeout: time.Second * 5,
	}

	return volumeNotify, volumeStr, muted, curVolume, nil
}

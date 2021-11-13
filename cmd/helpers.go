package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
	"mrogalski.eu/go/pulseaudio"
)

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

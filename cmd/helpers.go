package cmd

import (
	"fmt"
	"strconv"

	"mrogalski.eu/go/pulseaudio"
)

// ConvertToFloat is a helper function that takes a string argument and
// converts it to a float between 0 and 1.
func ConvertToFloat(num string) (float32, error) {
	intVal, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		return 0.0, err
	}

	if intVal >= 0 && intVal <= 100 {
		floatVal := float32(intVal) / 100
		return floatVal, nil
	} else {
		err := fmt.Errorf("Please provide a number between 0 and 100")
		return 0.0, err
	}

}

func GetCurrentVolume(client *pulseaudio.Client) (string, bool, float32, error) {
	curVolume, err := client.Volume()
	if err != nil {
		return "", false, 0.0, err
	}

	muted, err := client.Mute()
	if err != nil {
		return "", false, 0.0, err
	}

	var volumeStr string
	if muted {
		volumeStr = fmt.Sprintf("Muted: %d", int(curVolume*100))
	} else {
		volumeStr = fmt.Sprintf("Volume: %d", int(curVolume*100))
	}

	return volumeStr, muted, curVolume, nil
}

package cmd

import (
	"fmt"
	"strconv"
)

// ConvertToFloat is a helper function that takes a string argument and
// converts it to a float between 0 and 1.
func ConvertToFloat(num string) (float32, error) {
	intVal, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		fmt.Errorf("Error converting string to integer")
		return 0.0, err
	}

	if intVal >= 0 && intVal <= 100 {
		floatVal := float32(intVal) / 100
		return floatVal, nil
	} else {
		fmt.Errorf("Please provide a number between 0 and 100")
		return 0.0, err
	}

}

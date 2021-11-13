/*
Copyright © 2021 David Brown <dave.brown@exit107.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"mrogalski.eu/go/pulseaudio"

	"github.com/spf13/cobra"
)

// incCmd represents the inc command
var incCmd = &cobra.Command{
	Use:   "inc",
	Short: "Increase the volume by a percentage",
	Long:  "Increase the volume by a percentage.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pulseaudio.NewClient()
		if err != nil {
			fmt.Println("Error encountered in building client")
		}

		curVolume, err := client.Volume()
		if err != nil {
			fmt.Printf("Error encountered: %v\n", err)
		}
		changeBy, err := ConvertToFloat(args[0])
		if err != nil {
			fmt.Printf("Error encountered: %v\n", err)
		}
		newVolume := curVolume + changeBy

		if newVolume <= 1 {
			client.SetVolume(newVolume)
		} else {
			client.SetVolume(1)
		}

		volumeStr, _, _, err := GetCurrentVolume(client)
		if err != nil {
			fmt.Printf("Error encountered: %v\n", err)
		}
		fmt.Println(volumeStr)
	},
}

func init() {
	volumeCmd.AddCommand(incCmd)
}

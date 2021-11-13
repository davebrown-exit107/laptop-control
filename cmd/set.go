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
	"log"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
	"mrogalski.eu/go/pulseaudio"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the volume",
	Long:  "Set the volume",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pulseaudio.NewClient()
		if err != nil {
			log.Println("Error encountered in building client")
		}

		newVolume, err := ConvertToFloat(args[0])
		if err != nil {
			log.Printf("Error encountered: %v\n", err)
		}

		client.SetVolume(newVolume)

		volumeNotify, _, _, _, err := GetCurrentVolume(client)
		conn, err := dbus.SessionBusPrivate()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		if err = conn.Auth(nil); err != nil {
			panic(err)
		}

		if err = conn.Hello(); err != nil {
			panic(err)
		}

		_, err = notify.SendNotification(conn, volumeNotify)
		if err != nil {
			log.Printf("error sending notification: %v", err.Error())
		}
	},
}

func init() {
	volumeCmd.AddCommand(setCmd)
}

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
	"github.com/spf13/cobra"
)

// screenCmd represents the screen command
var screenCmd = &cobra.Command{
	Use:   "screen",
	Short: "Control the laptop screen brightness",
	Long:  "Control the laptop screen brightness",
	Run: func(cmd *cobra.Command, args []string) {
		curBrightness := GetCurBrightness()

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

		_, err = notify.SendNotification(conn, curBrightness.notifyMsg)
		if err != nil {
			log.Printf("error sending notification: %v", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(screenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// screenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// screenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	"os"
	"path"
	"strconv"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
)

// screenCmd represents the screen command
var screenCmd = &cobra.Command{
	Use:   "screen",
	Short: "Control the laptop screen brightness",
	Long:  "Control the laptop screen brightness",
	Args:  cobra.NoArgs,
}

var incScreenCmd = &cobra.Command{
	Use:   "inc",
	Short: "Increase the laptop screen brightness",
	Long:  "Increase the laptop screen brightness",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sysBasePath := "/sys/class/backlight/intel_backlight"

		curBrightness := GetCurBrightness()

		changeBy, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		newBrightnessPercent := curBrightness.percent + changeBy + 1
		newBrightness := ((curBrightness.maxBrightness - 1) * newBrightnessPercent) / 100

		if newBrightnessPercent >= 100 {
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte(strconv.Itoa(curBrightness.maxBrightness)), 0644)
		} else {
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte(strconv.Itoa(newBrightness)), 0644)
		}

		curBrightness = GetCurBrightness()

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

var decScreenCmd = &cobra.Command{
	Use:   "dec",
	Short: "decrease the laptop screen brightness",
	Long:  "decrease the laptop screen brightness",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sysBasePath := "/sys/class/backlight/intel_backlight"

		curBrightness := GetCurBrightness()

		changeBy, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		newBrightnessPercent := curBrightness.percent - changeBy + 1
		newBrightness := ((curBrightness.maxBrightness - 1) * newBrightnessPercent) / 100

		if newBrightnessPercent <= 0 {
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte("0"), 0644)
		} else {
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte(strconv.Itoa(newBrightness)), 0644)
		}

		curBrightness = GetCurBrightness()

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

var getScreenCmd = &cobra.Command{
	Use:   "get",
	Short: "getrease the laptop screen brightness",
	Long:  "getrease the laptop screen brightness",
	Args:  cobra.NoArgs,
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

var setScreenCmd = &cobra.Command{
	Use:   "set",
	Short: "setrease the laptop screen brightness",
	Long:  "setrease the laptop screen brightness",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sysBasePath := "/sys/class/backlight/intel_backlight"

		curBrightness := GetCurBrightness()

		newBrightnessPercent, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		switch {
		case newBrightnessPercent >= 100:
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte(strconv.Itoa(curBrightness.maxBrightness)), 0644)
		case newBrightnessPercent <= 0:
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte("0"), 0644)
		default:
			newBrightness := ((curBrightness.maxBrightness - 1) * newBrightnessPercent) / 100
			os.WriteFile(path.Join(sysBasePath, "brightness"), []byte(strconv.Itoa(newBrightness)), 0644)
		}

		curBrightness = GetCurBrightness()

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
	screenCmd.AddCommand(incScreenCmd)
	screenCmd.AddCommand(decScreenCmd)
	screenCmd.AddCommand(getScreenCmd)
	screenCmd.AddCommand(setScreenCmd)
}

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
	"mrogalski.eu/go/pulseaudio"

	"log"

	"github.com/spf13/cobra"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
)

// volumeCmd represents the volume command
var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Control the laptop's volume.",
	Long:  `Control the laptop's volume.`,
	Args:  cobra.NoArgs,
}

// getVolumeCmd represents the get subcommand of the volume command
var getVolumeCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the current volume",
	Long:  "Get the current volume",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pulseaudio.NewClient()
		if err != nil {
			log.Println("Error encountered in building client")
		}

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

// decVolumeCmd represents the dec subcommand of the volume command
var decVolumeCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decrease the volume by a percentage",
	Long:  "Decrease the volume by a percentage",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pulseaudio.NewClient()
		if err != nil {
			log.Println("Error encountered in building client")
		}

		curVolume, err := client.Volume()
		if err != nil {
			log.Printf("Error encountered: %v\n", err)
		}
		changeBy, err := ConvertToFloat(args[0])
		if err != nil {
			log.Printf("Error encountered: %v\n", err)
		}
		newVolume := curVolume - changeBy

		if newVolume > 0 {
			client.SetVolume(newVolume)
		} else {
			client.SetVolume(0)
		}

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

// incVolumeCmd represents the inc subcommand of the volume command
var incVolumeCmd = &cobra.Command{
	Use:   "inc",
	Short: "Increase the volume by a percentage",
	Long:  "Increase the volume by a percentage.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pulseaudio.NewClient()
		if err != nil {
			log.Println("Error encountered in building client")
		}

		curVolume, err := client.Volume()
		if err != nil {
			log.Printf("Error encountered: %v\n", err)
		}
		changeBy, err := ConvertToFloat(args[0])
		if err != nil {
			log.Printf("Error encountered: %v\n", err)
		}
		newVolume := curVolume + changeBy

		if newVolume <= 1 {
			client.SetVolume(newVolume)
		} else {
			client.SetVolume(1)
		}

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

// setVolumeCmd represents the set subcommand of the volume command
var setVolumeCmd = &cobra.Command{
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

// muteVolumeCmd represents the mute subcommand of the volume command
var muteVolumeCmd = &cobra.Command{
	Use:   "mute",
	Short: "Toggle mute",
	Long:  "Toggle mute",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := pulseaudio.NewClient()
		if err != nil {
			log.Println("Error encountered in building client")
		}

		client.ToggleMute()

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
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(getVolumeCmd)
	volumeCmd.AddCommand(decVolumeCmd)
	volumeCmd.AddCommand(setVolumeCmd)
	volumeCmd.AddCommand(incVolumeCmd)
	volumeCmd.AddCommand(muteVolumeCmd)
}

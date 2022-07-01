/*

Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/adrianburgoscolas/drm/utils"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start the Drum Machine",
	Long: `Start the Drum Machine with the
default sound bank. Awesome!`,
	Run: func(cmd *cobra.Command, args []string) {

		exec.Command("stty", "-F", "/dev/tty", "-cooked").Run()
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		defer exec.Command("stty", "-F", "/dev/tty", "cooked").Run()
		key := make([]byte, 1)
		for {
			os.Stdin.Read(key)
			if key[0] == '0' {
				break
			}
			switch key[0] {
			case 'q':
				go utils.Play("kick.mp3")
			case 'w':
				go utils.Play("snr.mp3")
			}
		}
		fmt.Print("Bye")
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

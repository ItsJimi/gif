package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gif",
	Short: "gif convert videos to animated gifs",
	Long:  "gif convert and crop videos to animated gif using ffmpeg.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Type `gif help` to view commands available")
	},
}

// Execute cli init
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

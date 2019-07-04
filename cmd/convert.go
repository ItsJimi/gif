package cmd

import (
	"github.com/ItsJimi/gif-cli/pkg/convert"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert videos to animated gif",
	Long:  "Convert videos to animated gif.",
	Run: func(cmd *cobra.Command, args []string) {
		convert.Convert()
	},
}

package cmd

import (
	"fmt"
	"os"

	"github.com/ItsJimi/gif-cli/pkg/convert"
	"github.com/spf13/cobra"
)

func init() {
	convertCmd.Flags().StringVarP(&crop, "crop", "c", "", "Crop videos before conversion")
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:     "convert",
	Short:   "Convert videos to animated gif",
	Long:    "Convert videos to animated gif.",
	Example: `  gif convert my-video.mp4 my-gif.gif`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid arguments, type `gif convert -h` to see arguments available.")
			return
		}

		inputPath, err := os.Stat(args[0])
		if err != nil {
			fmt.Println(args[0] + ": no such file or directory")
			return
		}

		if inputPath.Mode().IsDir() == true {
			if len(args) != 2 {
				fmt.Println("Invalid arguments, type `gif convert -h` to see arguments available.")
				return
			}

			outputPath, err := os.Stat(args[1])
			if err != nil || outputPath.Mode().IsDir() == false {
				fmt.Println(args[1] + ": no such directory")
				return
			}

			convert.FromFolder(inputPath.Name(), outputPath.Name())
			return
		}

		crop := cmd.Flag("crop")
		if crop.Changed == false {
			// convert.FromFile()
		}
	},
}

var crop string

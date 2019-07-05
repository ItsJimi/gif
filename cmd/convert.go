package cmd

import (
	"fmt"
	"os"

	"github.com/ItsJimi/gif/pkg/convert"
	"github.com/spf13/cobra"
)

func init() {
	convertCmd.Flags().IntVarP(&fps, "fps", "f", 30, "gif convert my-video.mp4 my-gif.gif --fps 30")
	convertCmd.Flags().IntVarP(&scale, "scale", "s", -1, "gif convert my-video.mp4 my-gif.gif --scale -1")
	convertCmd.Flags().StringVarP(&crop, "crop", "c", "", "gif convert my-video.mp4 my-gif.gif --crop \"width:height:left:top\"")
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert videos to animated gif",
	Long:  "Convert videos to animated gif.",
	Example: `  gif convert my-video.mp4 my-gif.gif
  gif convert my-video.mp4 my-gif.gif --fps 60 --scale 1280
  gif convert ./folder ./another-folder --crop "1280:720:30:60"`,
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

			err = convert.FromFolder(inputPath.Name(), outputPath.Name(), convert.Options{
				FPS:     fps,
				Scale:   scale,
				Crop:    crop,
				Verbose: true,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}

		crop := cmd.Flag("crop")
		if crop.Changed == false {
			// convert.FromFile()
		}
	},
}

var fps int
var scale int
var crop string

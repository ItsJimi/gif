package convert

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FromFolder convert videos to gif from folder
func FromFolder(inputPath string, outputPath string) {
	var files []os.FileInfo

	err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(path)) != ".mov" {
			return nil
		}
		files = append(files, info)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(inputPath + " -> " + outputPath)
	for _, file := range files {
		fmt.Println("Converting " + file.Name() + " to " + file.Name() + ".gif")
		ffmpeg := exec.Command("ffmpeg", "-i", inputPath+"/"+file.Name(), "-vf", "fps=30,scale=1920:-1:flags=lanczos,palettegen", "-y", outputPath+"/palette.png")
		if err := ffmpeg.Run(); err != nil {
			fmt.Println(err)
		}
		ffmpeg = exec.Command("ffmpeg", "-i", inputPath+"/"+file.Name(), "-i", outputPath+"/palette.png", "-filter_complex", "fps=30,scale=500:-1:flags=lanczos[x];[x][1:v]paletteuse", "-y", outputPath+"/"+file.Name()+".gif")
		if err = ffmpeg.Run(); err != nil {
			fmt.Println(err)
		}
	}
}

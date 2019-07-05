package convert

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Options availables
type Options struct {
	FPS     int
	Scale   int
	Crop    string
	Verbose bool
}

func convertFiles(files []os.FileInfo, inputPath string, outputPath string, opt Options) error {
	cropFlag := ""
	if opt.Crop != "" {
		cropFlag = ",crop="
	}

	for _, file := range files {
		if opt.Verbose == true {
			fmt.Println("Converting " + file.Name() + " to " + file.Name() + ".gif")
		}
		ffmpeg := exec.Command("ffmpeg", "-i", inputPath+"/"+file.Name(), "-vf", "fps="+strconv.Itoa(opt.FPS)+",scale="+strconv.Itoa(opt.Scale)+":-1:flags=lanczos,palettegen", "-y", outputPath+"/"+file.Name()+".png")
		if err := ffmpeg.Run(); err != nil {
			return err
		}
		ffmpeg = exec.Command("ffmpeg", "-i", inputPath+"/"+file.Name(), "-i", outputPath+"/"+file.Name()+".png", "-filter_complex", "fps="+strconv.Itoa(opt.FPS)+",scale="+strconv.Itoa(opt.Scale)+":-1:flags=lanczos[x];[x][1:v]paletteuse"+cropFlag+opt.Crop, "-y", outputPath+"/"+file.Name()+".gif")
		if err := ffmpeg.Run(); err != nil {
			return err
		}
		if err := os.Remove(inputPath + "/" + file.Name() + ".png"); err != nil {
			return err
		}
	}
	return nil
}

// FromFolder convert videos to gif from folder
func FromFolder(inputPath string, outputPath string, opt Options) error {
	var files []os.FileInfo

	err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == true || strings.ToLower(filepath.Ext(path)) == ".gif" {
			return nil
		}
		files = append(files, info)
		return nil
	})
	if err != nil {
		return err
	}

	if opt.Verbose == true {
		fmt.Println(inputPath + " -> " + outputPath)
	}
	if err := convertFiles(files, inputPath, outputPath, opt); err != nil {
		return err
	}
	return nil
}

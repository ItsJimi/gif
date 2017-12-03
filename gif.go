package main

import (
	"fmt"
	"os"

	"github.com/ItsJimi/gif"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("To create a gif, run: gif file.[jpg|jpeg|png|gif]...")
		os.Exit(0)
	}
	fmt.Println("Runing ...")
	var images []string
	for i := 1; i < len(os.Args); i++ {
		images = append(images, os.Args[i])
	}
	gif.Create(images, "out.gif")
	fmt.Println("Done :)")
	os.Exit(0)
}

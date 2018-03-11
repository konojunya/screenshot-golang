package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/vova616/screenshot"
)

func main() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}

	myImg := image.Image(img)
	file, _ := os.Create("screen.jpg")
	defer file.Close()

	err = jpeg.Encode(file, myImg, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(err)
	}
}

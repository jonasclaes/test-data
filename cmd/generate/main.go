package main

import (
	"flag"
	"log"
)
import "github.com/fogleman/gg"

func main() {
	text := flag.String("text", "EN", "Text to overlay")
	baseImage := flag.String("base", ".assets/images/base-square.png", "Base image path")
	output := flag.String("output", "output.png", "Output path")
	fontPath := flag.String("font", "", "Path to TTF font (optional)")
	flag.Parse()

	img, err := gg.LoadImage(*baseImage)
	if err != nil {
		log.Fatalf("error loading image: %v", err)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	canvas := gg.NewContext(width, height)
	canvas.DrawImage(img, 0, 0)

	fontSize := float64(width) / 3

	if *fontPath != "" {
		if err := canvas.LoadFontFace(*fontPath, fontSize); err != nil {
			log.Fatalf("error loading font: %v", err)
		}
	} else {
		if err := canvas.LoadFontFace(".assets/fonts/WorkSans-Bold.ttf", fontSize); err != nil {
			log.Fatalf("error loading default font: %v", err)
		}
	}

	canvas.SetRGB(1, 1, 1)

	canvas.DrawStringAnchored(*text, float64(width)/2, float64(height)/2, 0.5, 0.5)

	if err := canvas.SavePNG(*output); err != nil {
		log.Fatalf("error saving image: %v", err)
	}
}

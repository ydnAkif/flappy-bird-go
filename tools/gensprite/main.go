package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// Create a 32x32 image
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with yellow color for bird body
	yellow := color.RGBA{255, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	// Draw bird body
	for y := 8; y < 24; y++ {
		for x := 8; x < 24; x++ {
			img.Set(x, y, yellow)
		}
	}

	// Draw eye
	img.Set(20, 12, black)

	// Draw beak
	for y := 16; y < 19; y++ {
		img.Set(24, y, black)
	}

	// Create output file
	f, err := os.Create("internal/game/assets/bird.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Encode as PNG
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

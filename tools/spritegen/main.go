package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

const (
	frameWidth  = 32
	frameHeight = 32
	numFrames   = 3
)

func main() {
	// Create a new RGBA image for the sprite sheet
	img := image.NewRGBA(image.Rect(0, 0, frameWidth*numFrames, frameHeight))

	// Colors for the bird
	bodyColor := color.RGBA{255, 217, 0, 255}   // Yellow
	wingColor := color.RGBA{255, 165, 0, 255}   // Orange
	eyeColor := color.RGBA{0, 0, 0, 255}        // Black
	beakColor := color.RGBA{255, 69, 0, 255}    // Red-Orange

	// Draw three different frames
	for frame := 0; frame < numFrames; frame++ {
		offsetX := frame * frameWidth

		// Draw body (circle)
		for y := 8; y < 24; y++ {
			for x := 8; x < 24; x++ {
				dx := float64(x - 16)
				dy := float64(y - 16)
				if dx*dx+dy*dy <= 64 { // radius of 8
					img.Set(x+offsetX, y, bodyColor)
				}
			}
		}

		// Draw wing (changes position in each frame)
		wingY := 16
		switch frame {
		case 0:
			wingY = 18 // Down position
		case 1:
			wingY = 16 // Middle position
		case 2:
			wingY = 14 // Up position
		}

		for y := wingY; y < wingY+6; y++ {
			for x := 10; x < 16; x++ {
				img.Set(x+offsetX, y, wingColor)
			}
		}

		// Draw eye
		img.Set(18+offsetX, 14, eyeColor)
		img.Set(19+offsetX, 14, eyeColor)

		// Draw beak
		for y := 15; y < 17; y++ {
			for x := 20; x < 24; x++ {
				img.Set(x+offsetX, y, beakColor)
			}
		}
	}

	// Create output directory if it doesn't exist
	outputDir := "../../internal/assets/sprites"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Save the sprite sheet
	outputPath := filepath.Join(outputDir, "bird.png")
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	log.Printf("Sprite sheet generated: %s\n", outputPath)
}

package spritegen

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// GenerateSprites generates bird sprite images
func GenerateSprites() {
	// Create the bird images directory
	err := os.MkdirAll("../assets/images", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Generate three frames of bird animation
	for i := 1; i <= 3; i++ {
		img := image.NewRGBA(image.Rect(0, 0, 30, 30))

		// Draw bird body (yellow circle)
		for y := 0; y < 30; y++ {
			for x := 0; x < 30; x++ {
				dx := float64(x - 15)
				dy := float64(y - 15)
				if dx*dx+dy*dy <= 225 { // 15^2 = 225
					img.Set(x, y, color.RGBA{255, 255, 0, 255}) // Yellow
				}
			}
		}

		// Draw eye (black dot)
		for y := 8; y < 12; y++ {
			for x := 20; x < 24; x++ {
				img.Set(x, y, color.Black)
			}
		}

		// Draw wing (changes position based on frame)
		wingY := 15
		switch i {
		case 1:
			wingY = 12
		case 2:
			wingY = 15
		case 3:
			wingY = 18
		}

		for y := wingY; y < wingY+8; y++ {
			for x := 5; x < 15; x++ {
				img.Set(x, y, color.RGBA{218, 165, 32, 255}) // Golden
			}
		}

		// Save the image
		f, err := os.Create("../assets/images/bird" + string(rune(i+'0')) + ".png")
		if err != nil {
			log.Fatal(err)
		}
		if err := png.Encode(f, img); err != nil {
			f.Close()
			log.Fatal(err)
		}
		f.Close()
	}
}

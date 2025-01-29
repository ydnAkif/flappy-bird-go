package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	// Create three different bird frames
	frames := []struct {
		wingOffset int
		filename   string
	}{
		{0, "../assets/images/bird1.png"},   // Wings level
		{-3, "../assets/images/bird2.png"},  // Wings up
		{3, "../assets/images/bird3.png"},   // Wings down
	}

	for _, frame := range frames {
		// Create a new RGBA image
		img := image.NewRGBA(image.Rect(0, 0, 30, 30))

		// Fill with transparent background
		draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

		// Bird body (yellow circle)
		drawCircle(img, 15, 15, 12, color.RGBA{0xFF, 0xFF, 0x00, 0xFF})

		// Bird eye (black dot)
		drawCircle(img, 22, 12, 3, color.Black)

		// Bird wing
		drawWing(img, 8, 15+frame.wingOffset, color.RGBA{0xCC, 0xCC, 0x00, 0xFF})

		// Save the image
		if err := saveImage(frame.filename, img); err != nil {
			log.Fatal(err)
		}
	}
}

func drawCircle(img *image.RGBA, x, y, r int, c color.Color) {
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r*r {
				img.Set(x+dx, y+dy, c)
			}
		}
	}
}

func drawWing(img *image.RGBA, x, y int, c color.Color) {
	points := []struct{ x, y int }{
		{0, 0},
		{8, -4},
		{8, 4},
	}

	// Draw a simple triangle for the wing
	for dy := -4; dy <= 4; dy++ {
		for dx := 0; dx <= 8; dx++ {
			inside := pointInTriangle(
				float64(dx), float64(dy),
				float64(points[0].x), float64(points[0].y),
				float64(points[1].x), float64(points[1].y),
				float64(points[2].x), float64(points[2].y),
			)
			if inside {
				img.Set(x+dx, y+dy, c)
			}
		}
	}
}

func pointInTriangle(px, py, x1, y1, x2, y2, x3, y3 float64) bool {
	area := 0.5 * (-y2*x3 + y1*(-x2 + x3) + x1*(y2 - y3) + x2*y3)
	s := 1 / (2 * area) * (y1*x3 - x1*y3 + (y3-y1)*px + (x1-x3)*py)
	t := 1 / (2 * area) * (x1*y2 - y1*x2 + (y1-y2)*px + (x2-x1)*py)
	return s > 0 && t > 0 && 1-s-t > 0
}

func saveImage(filename string, img image.Image) error {
	// Create directories if they don't exist
	dir := "../assets/images"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Create the file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode and save
	return png.Encode(f, img)
}

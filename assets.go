package main

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/images/*.png
var assets embed.FS

// BirdSprites holds all the animation frames for the bird
var BirdSprites []*ebiten.Image

// LoadAssets loads all game assets
func LoadAssets() {
	// Load bird sprites
	birdFrames := []string{
		"assets/images/bird1.png",
		"assets/images/bird2.png",
		"assets/images/bird3.png",
	}

	BirdSprites = make([]*ebiten.Image, len(birdFrames))
	for i, path := range birdFrames {
		data, err := assets.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to load sprite %s: %v", path, err)
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			log.Fatalf("failed to decode sprite %s: %v", path, err)
		}

		BirdSprites[i] = ebiten.NewImageFromImage(img)
	}
}

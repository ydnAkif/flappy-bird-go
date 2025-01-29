// Package game handles the game logic
package game

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/bird.png
var birdSprite []byte

// BirdSprites contains all bird animation frames
var BirdSprites []*ebiten.Image

// LoadAssets loads all game assets
func LoadAssets() error {
	// Load bird sprite
	img, _, err := image.Decode(bytes.NewReader(birdSprite))
	if err != nil {
		log.Printf("Error loading bird sprite: %v", err)
		return err
	}
	BirdSprites = []*ebiten.Image{ebiten.NewImageFromImage(img)}

	return nil
}

package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Bird represents the player character
type Bird struct {
	x, y        float64
	vy          float64
	gravity     float64
	frameCount  int
	frameIndex  int
	rotation    float64
}

// Update handles bird physics and animation
func (b *Bird) Update() {
	b.vy += b.gravity
	b.y += b.vy
	
	// Update animation frame every 5 game ticks
	b.frameCount++
	if b.frameCount >= 5 {
		b.frameCount = 0
		b.frameIndex = (b.frameIndex + 1) % len(BirdSprites)
	}

	// Calculate rotation based on velocity
	b.rotation = math.Atan2(b.vy*0.1, 1.0)
}

// Draw renders the bird
func (b *Bird) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	
	// Set the rotation origin to the center of the bird
	op.GeoM.Translate(-float64(birdSize)/2, -float64(birdSize)/2)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(b.x+float64(birdSize)/2, b.y+float64(birdSize)/2)

	// Draw the current animation frame
	if len(BirdSprites) > 0 {
		screen.DrawImage(BirdSprites[b.frameIndex], op)
	} else {
		// Fallback to rectangle if sprites aren't loaded
		ebitenutil.DrawRect(screen, b.x, b.y, birdSize, birdSize, color.RGBA{0xff, 0xff, 0x00, 0xff})
	}
}
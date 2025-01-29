package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Bird represents the player character
type Bird struct {
	x, y    float64
	vy      float64
	gravity float64
}

// Update handles bird physics
func (b *Bird) Update() {
	b.vy += b.gravity
	b.y += b.vy
}

// Draw renders the bird
func (b *Bird) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x, b.y, birdSize, birdSize, color.RGBA{0xff, 0xff, 0x00, 0xff})
}
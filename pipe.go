package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Pipe represents an obstacle in the game
type Pipe struct {
	x        float64
	topH     float64
	bottomH  float64
	passed   bool
}

// Draw renders the pipe
func (p *Pipe) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.x, 0, pipeWidth, p.topH, color.RGBA{0x00, 0x80, 0x00, 0xff})
	ebitenutil.DrawRect(screen, p.x, screenHeight-p.bottomH, pipeWidth, p.bottomH, color.RGBA{0x00, 0x80, 0x00, 0xff})
}
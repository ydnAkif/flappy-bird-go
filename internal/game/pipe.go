package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"flappy-bird-go/internal/config"
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
	ebitenutil.DrawRect(screen, p.x, 0, config.PipeWidth, p.topH, color.RGBA{0x00, 0x80, 0x00, 0xff})
	ebitenutil.DrawRect(screen, p.x, config.ScreenHeight-p.bottomH, config.PipeWidth, p.bottomH, color.RGBA{0x00, 0x80, 0x00, 0xff})
}
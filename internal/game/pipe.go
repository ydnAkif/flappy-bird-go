package game

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/ydnAkif/flappy-bird-go/internal/config"
)

// Pipe represents an obstacle in the game
type Pipe struct {
	X       float64
	TopH    float64
	BottomH float64
	passed  bool
	imd     *imdraw.IMDraw
}

// NewPipe creates a new pipe with random gap position
func NewPipe() *Pipe {
	gapY := rand.Float64()*(float64(config.ScreenHeight)-200) + 100
	return &Pipe{
		X:       float64(config.ScreenWidth),
		TopH:    gapY - float64(config.PipeGap)/2,
		BottomH: float64(config.ScreenHeight) - (gapY + float64(config.PipeGap)/2),
		passed:  false,
		imd:     imdraw.New(nil),
	}
}

// Update moves the pipe to the left
func (p *Pipe) Update(dt float64) {
	p.X -= float64(config.PipeSpeed) * dt * 60
}

// Draw draws the pipe
func (p *Pipe) Draw(win *pixelgl.Window) {
	p.imd.Clear()
	
	// Set pipe color (green)
	p.imd.Color = pixel.RGB(0, 0.8, 0)

	// Draw top pipe
	topRect := pixel.R(p.X, float64(config.ScreenHeight)-p.TopH, p.X+float64(config.PipeWidth), float64(config.ScreenHeight))
	p.imd.Push(topRect.Min, topRect.Max)
	p.imd.Rectangle(0)

	// Draw bottom pipe
	bottomRect := pixel.R(p.X, 0, p.X+float64(config.PipeWidth), p.BottomH)
	p.imd.Push(bottomRect.Min, bottomRect.Max)
	p.imd.Rectangle(0)

	p.imd.Draw(win)
}

// IsOffscreen returns true if the pipe is off the left side of the screen
func (p *Pipe) IsOffscreen() bool {
	return p.X+float64(config.PipeWidth) < 0
}

// Collides checks if the pipe collides with a given rectangle
func (p *Pipe) Collides(rect []float64) bool {
	birdX, birdY, birdW, birdH := rect[0], rect[1], rect[2], rect[3]

	// Check collision with top pipe
	if birdX < p.X+float64(config.PipeWidth) &&
		birdX+birdW > p.X &&
		birdY < float64(config.ScreenHeight)-p.TopH {
		return true
	}

	// Check collision with bottom pipe
	if birdX < p.X+float64(config.PipeWidth) &&
		birdX+birdW > p.X &&
		birdY+birdH > p.BottomH {
		return true
	}

	return false
}

// PassedBy returns true if the bird has passed this pipe and updates the passed flag
func (p *Pipe) PassedBy(rect []float64) bool {
	birdX := rect[0]
	if !p.passed && p.X+float64(config.PipeWidth) < birdX {
		p.passed = true
		return true
	}
	return false
}
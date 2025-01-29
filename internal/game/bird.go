package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/ydnAkif/flappy-bird-go/internal/config"
)

// Bird states
const (
	BirdStateIdle = iota
	BirdStateFlying
	BirdStateFalling
)

// Bird represents the player character
type Bird struct {
	pos    pixel.Vec
	vel    pixel.Vec
	state  int
	imd    *imdraw.IMDraw
}

// NewBird creates a new bird instance
func NewBird() (*Bird, error) {
	bird := &Bird{
		pos:   pixel.V(50, float64(config.ScreenHeight)/2),
		vel:   pixel.V(0, 0),
		state: BirdStateIdle,
		imd:   imdraw.New(nil),
	}
	return bird, nil
}

// Update updates the bird's position and state
func (b *Bird) Update(dt float64) {
	// Update velocity and position
	b.vel.Y += config.BirdGravity * dt * 60
	b.pos = b.pos.Add(b.vel.Scaled(dt * 60))

	// Update bird state based on velocity
	if b.vel.Y < 0 {
		b.state = BirdStateFlying
	} else if b.vel.Y > 2 {
		b.state = BirdStateFalling
	} else {
		b.state = BirdStateIdle
	}

	// Keep bird within screen bounds
	if b.pos.Y < 0 {
		b.pos.Y = 0
		b.vel.Y = 0
	} else if b.pos.Y > float64(config.ScreenHeight) {
		b.pos.Y = float64(config.ScreenHeight)
		b.vel.Y = 0
	}
}

// Draw draws the bird
func (b *Bird) Draw(win *pixelgl.Window) {
	b.imd.Clear()
	
	// Draw bird body (yellow circle)
	b.imd.Color = pixel.RGB(1, 1, 0)
	b.imd.Push(b.pos)
	b.imd.Circle(float64(config.BirdSize)/2, 0)

	// Draw eye (black dot)
	b.imd.Color = pixel.RGB(0, 0, 0)
	eyePos := b.pos.Add(pixel.V(float64(config.BirdSize)/4, float64(config.BirdSize)/4))
	b.imd.Push(eyePos)
	b.imd.Circle(2, 0)

	b.imd.Draw(win)
}

// Jump makes the bird jump
func (b *Bird) Jump() {
	b.vel.Y = config.BirdJumpVel
	b.state = BirdStateFlying
}

// GetBounds returns the bird's bounding box for collision detection
func (b *Bird) GetBounds() []float64 {
	size := float64(config.BirdSize)
	return []float64{
		b.pos.X - size/2,
		b.pos.Y - size/2,
		size,
		size,
	}
}
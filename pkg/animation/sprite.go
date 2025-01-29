package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image        *ebiten.Image
	frameWidth   int
	frameHeight  int
	frames       int
	currentFrame int
}

func NewSprite(image *ebiten.Image, frameWidth, frameHeight, frames int) *Sprite {
	return &Sprite{
		image:        image,
		frameWidth:   frameWidth,
		frameHeight:  frameHeight,
		frames:       frames,
		currentFrame: 0,
	}
}

func (s *Sprite) Update() {
	s.currentFrame = (s.currentFrame + 1) % s.frames
}

func (s *Sprite) Draw(screen *ebiten.Image, x, y float64) {
	// TODO: Implement sprite drawing
}

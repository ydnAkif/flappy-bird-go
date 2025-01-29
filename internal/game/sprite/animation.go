package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Animation represents a sprite animation
type Animation struct {
	spriteSheet *ebiten.Image
	frames      []image.Rectangle
	frameTime   float64
	currentTime float64
	currentFrame int
	isLoop      bool
}

// NewAnimation creates a new animation from a sprite sheet
func NewAnimation(spriteSheet *ebiten.Image, frameWidth, frameHeight int, frameTime float64, isLoop bool) *Animation {
	frames := make([]image.Rectangle, 0)
	sheetWidth := spriteSheet.Bounds().Dx()
	
	// Calculate frame rectangles
	for x := 0; x < sheetWidth; x += frameWidth {
		frames = append(frames, image.Rect(x, 0, x+frameWidth, frameHeight))
	}

	return &Animation{
		spriteSheet:  spriteSheet,
		frames:       frames,
		frameTime:    frameTime,
		currentTime:  0,
		currentFrame: 0,
		isLoop:       isLoop,
	}
}

// Update updates the animation state
func (a *Animation) Update(dt float64) {
	a.currentTime += dt
	if a.currentTime >= a.frameTime {
		a.currentTime = 0
		a.currentFrame++
		if a.currentFrame >= len(a.frames) {
			if a.isLoop {
				a.currentFrame = 0
			} else {
				a.currentFrame = len(a.frames) - 1
			}
		}
	}
}

// Draw draws the current frame
func (a *Animation) Draw(screen *ebiten.Image, x, y float64, scale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	
	screen.DrawImage(a.spriteSheet.SubImage(a.frames[a.currentFrame]).(*ebiten.Image), op)
}

// Reset resets the animation to the first frame
func (a *Animation) Reset() {
	a.currentTime = 0
	a.currentFrame = 0
}

// IsFinished returns true if the animation has finished (only for non-looping animations)
func (a *Animation) IsFinished() bool {
	return !a.isLoop && a.currentFrame == len(a.frames)-1
}

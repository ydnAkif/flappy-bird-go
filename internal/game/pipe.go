package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Pipe struct {
	x      float64
	y      float64
	width  float64
	height float64
	image  *ebiten.Image
	speed  float64
}

func NewPipe(x, y, width, height float64, image *ebiten.Image) *Pipe {
	return &Pipe{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		image:  image,
		speed:  PipeSpeed,
	}
}

func (p *Pipe) Update() {
	p.x -= p.speed
}

func (p *Pipe) Draw(screen *ebiten.Image) {
	if p.image == nil {
		return
	}

	// Draw top pipe (inverted)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, -1) // Flip vertically
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, op)

	// Draw bottom pipe
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y+PipeGap)
	screen.DrawImage(p.image, op)
}

// GetTopBoundingBox returns the collision box for the top pipe
func (p *Pipe) GetTopBoundingBox() BoundingBox {
	return BoundingBox{
		X:      p.x,
		Y:      0,
		Width:  p.width,
		Height: p.y, // Top pipe extends from 0 to gap position
	}
}

// GetBottomBoundingBox returns the collision box for the bottom pipe
func (p *Pipe) GetBottomBoundingBox() BoundingBox {
	gapBottom := p.y + PipeGap
	return BoundingBox{
		X:      p.x,
		Y:      gapBottom,
		Width:  p.width,
		Height: float64(ScreenHeight) - gapBottom,
	}
}

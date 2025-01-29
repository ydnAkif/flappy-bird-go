package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bird struct {
	x, y    float64
	dy      float64
	gravity float64
	image   *ebiten.Image
}

func NewBird(x, y float64, image *ebiten.Image) *Bird {
	return &Bird{
		x:       x,
		y:       y,
		dy:      0,
		gravity: 0.5,
		image:   image,
	}
}
func (b *Bird) Update() {
	b.dy += b.gravity
	b.y += b.dy

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		b.Jump()
	}
}

func (b *Bird) Draw(screen *ebiten.Image) {
	if b.image == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	// Calculate rotation based on velocity
	rotation := b.dy * 0.05
	op.GeoM.Translate(-float64(b.image.Bounds().Dx())/2, -float64(b.image.Bounds().Dy())/2)
	op.GeoM.Rotate(rotation)
	op.GeoM.Translate(float64(b.image.Bounds().Dx())/2, float64(b.image.Bounds().Dy())/2)

	// Position the bird
	op.GeoM.Translate(b.x, b.y)

	screen.DrawImage(b.image, op)
}

func (b *Bird) Reset() {
	b.x = float64(ScreenWidth) / 3
	b.y = float64(ScreenHeight) / 2
	b.dy = 0
}

func (b *Bird) Jump() {
	b.dy = -8
}

// GetBoundingBox returns the bird's collision box
func (b *Bird) GetBoundingBox() BoundingBox {
	return BoundingBox{
		X:      b.x,
		Y:      b.y,
		Width:  float64(b.image.Bounds().Dx()),
		Height: float64(b.image.Bounds().Dy()),
	}
}

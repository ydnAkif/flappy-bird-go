package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	
	"github.com/ydnAkif/flappy-bird-go/internal/config"
	"github.com/ydnAkif/flappy-bird-go/internal/game/sprite"
)

// Bird states
const (
	BirdStateIdle = iota
	BirdStateFlying
	BirdStateFalling
)

// Bird represents the player character
type Bird struct {
	x, y    float64
	vy      float64
	state   int
	animations map[int]*sprite.Animation
	currentAnim *sprite.Animation
}

// NewBird creates a new bird instance
func NewBird() (*Bird, error) {
	// Load bird sprite sheet
	birdImg, _, err := ebitenutil.NewImageFromFile("internal/assets/sprites/bird.png")
	if err != nil {
		return nil, err
	}

	bird := &Bird{
		x:       50,
		y:       config.ScreenHeight / 2,
		state:   BirdStateIdle,
		animations: make(map[int]*sprite.Animation),
	}

	// Create animations for different states
	bird.animations[BirdStateIdle] = sprite.NewAnimation(birdImg, 32, 32, 0.1, true)    // Idle animation
	bird.animations[BirdStateFlying] = sprite.NewAnimation(birdImg, 32, 32, 0.08, true) // Flying animation
	bird.animations[BirdStateFalling] = sprite.NewAnimation(birdImg, 32, 32, 0.1, false) // Falling animation
	
	bird.currentAnim = bird.animations[BirdStateIdle]
	return bird, nil
}

// Update updates the bird's position and state
func (b *Bird) Update() {
	// Update velocity and position
	b.vy += config.BirdGravity
	b.y += b.vy

	// Update bird state based on velocity
	newState := b.state
	if b.vy < 0 {
		newState = BirdStateFlying
	} else if b.vy > 2 {
		newState = BirdStateFalling
	} else {
		newState = BirdStateIdle
	}

	// Change animation if state changed
	if newState != b.state {
		b.state = newState
		b.currentAnim = b.animations[b.state]
		if b.state == BirdStateFalling {
			b.currentAnim.Reset()
		}
	}

	// Update current animation
	b.currentAnim.Update(1.0 / 60.0) // Assuming 60 FPS
}

// Draw draws the bird
func (b *Bird) Draw(screen *ebiten.Image) {
	b.currentAnim.Draw(screen, b.x, b.y, 1.0)
}

// Jump makes the bird jump
func (b *Bird) Jump() {
	b.vy = config.BirdJumpVel
	b.state = BirdStateFlying
	b.currentAnim = b.animations[b.state]
	b.currentAnim.Reset()
}

// GetBounds returns the bird's bounding box for collision detection
func (b *Bird) GetBounds() []float64 {
	return []float64{
		b.x,
		b.y,
		float64(config.BirdSize),
		float64(config.BirdSize),
	}
}
package game

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type PipeManager struct {
	pipes       []*Pipe
	spawnTime   float64
	lastSpawn   float64
	scoredPipes map[*Pipe]bool
	pipeImage   *ebiten.Image
}

// CheckScore checks if the bird has passed a pipe and should score a point
func (pm *PipeManager) CheckScore(bird *Bird) bool {
	for _, pipe := range pm.pipes {
		// If we haven't scored this pipe yet and bird has passed it
		if !pm.scoredPipes[pipe] && bird.x > pipe.x+pipe.width {
			pm.scoredPipes[pipe] = true
			return true
		}
	}
	return false
}

// CheckCollision checks if the bird has collided with any pipe
func (pm *PipeManager) CheckCollision(bird *Bird) bool {
	birdBox := bird.GetBoundingBox()

	for _, pipe := range pm.pipes {
		topPipeBox := pipe.GetTopBoundingBox()
		bottomPipeBox := pipe.GetBottomBoundingBox()

		if birdBox.Intersects(topPipeBox) || birdBox.Intersects(bottomPipeBox) {
			return true
		}
	}
	return false
}

func (pm *PipeManager) Update() {
	pm.lastSpawn += 1.0 / 60.0 // Convert frames to seconds

	// Spawn new pipe
	if pm.lastSpawn >= pm.spawnTime {
		pipeY := rand.Float64()*(float64(ScreenHeight)-PipeGap-100) + 50

		newPipe := NewPipe(
			float64(ScreenWidth), // x
			pipeY,                // y
			PipeWidth,            // width
			PipeHeight,           // height
			pm.pipeImage,         // image
		)
		pm.pipes = append(pm.pipes, newPipe)
		pm.lastSpawn = 0
	}

	// Update existing pipes
	for _, pipe := range pm.pipes {
		pipe.Update()
	}

	pm.pipes = filterPipes(pm.pipes)
}

func (pm *PipeManager) Draw(screen *ebiten.Image) {
	for _, pipe := range pm.pipes {
		pipe.Draw(screen)
	}
}

// Reset resets the pipe manager state
func (pm *PipeManager) Reset() {
	pm.pipes = make([]*Pipe, 0)
	pm.scoredPipes = make(map[*Pipe]bool)
	pm.lastSpawn = 0
}

func filterPipes(pipes []*Pipe) []*Pipe {
	filtered := pipes[:0]
	for _, pipe := range pipes {
		if pipe.x > -pipe.width {
			filtered = append(filtered, pipe)
		}
	}
	return filtered
}

func NewPipeManager(pipeImage *ebiten.Image) *PipeManager {
	return &PipeManager{
		pipes:       make([]*Pipe, 0),
		spawnTime:   2.0,
		lastSpawn:   0,
		scoredPipes: make(map[*Pipe]bool),
		pipeImage:   pipeImage,
	}
}

package game

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"flappy-bird-go/internal/config"
)

// Game represents the game state
type Game struct {
	bird      *Bird
	pipes     []Pipe
	score     int
	gameOver  bool
	spaceJust bool
}

// NewGame creates a new instance of the game
func NewGame() (*Game, error) {
	bird, err := NewBird()
	if err != nil {
		return nil, err
	}

	g := &Game{
		bird:  bird,
		pipes: make([]Pipe, 0),
	}
	g.addPipe()
	return g, nil
}

func (g *Game) addPipe() {
	gapY := rand.Float64()*(config.ScreenHeight-200) + 100
	pipe := Pipe{
		x:       config.ScreenWidth,
		topH:    gapY - config.PipeGap/2,
		bottomH: config.ScreenHeight - (gapY + config.PipeGap/2),
		passed:  false,
	}
	g.pipes = append(g.pipes, pipe)
}

// Update handles game logic updates
func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			newGame, err := NewGame()
			if err != nil {
				return err
			}
			*g = *newGame
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.spaceJust {
			g.bird.Jump()
			g.spaceJust = true
		}
	} else {
		g.spaceJust = false
	}

	g.bird.Update()

	// Update pipes and check collisions
	for i := range g.pipes {
		g.pipes[i].x -= 2

		if g.checkCollisionWithPipe(i) {
			g.gameOver = true
			PlayHitSound()
		}

		if !g.pipes[i].passed && g.pipes[i].x+config.PipeWidth < g.bird.x {
			g.score++
			g.pipes[i].passed = true
			PlayScoreSound()
		}
	}

	// Remove off-screen pipes
	if len(g.pipes) > 0 && g.pipes[0].x < -config.PipeWidth {
		g.pipes = g.pipes[1:]
	}

	// Add new pipes
	if len(g.pipes) == 0 || g.pipes[len(g.pipes)-1].x < config.ScreenWidth-200 {
		g.addPipe()
	}

	// Check ceiling and floor collisions
	if g.bird.y < 0 || g.bird.y > config.ScreenHeight-config.BirdSize {
		g.gameOver = true
		PlayHitSound()
	}

	return nil
}

// Draw renders the game
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	
	g.bird.Draw(screen)
	
	for _, pipe := range g.pipes {
		pipe.Draw(screen)
	}

	scoreStr := fmt.Sprintf("Score: %d", g.score)
	ebitenutil.DebugPrint(screen, scoreStr)

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "\n\nGame Over!\nPress SPACE to restart")
	}
}

// Layout implements ebiten.Game's Layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) checkCollisionWithPipe(i int) bool {
	birdBounds := g.bird.GetBounds()
	pipeBounds := []float64{
		g.pipes[i].x, 0,
		config.PipeWidth, g.pipes[i].topH,
	}
	bottomPipeBounds := []float64{
		g.pipes[i].x, config.ScreenHeight - g.pipes[i].bottomH,
		config.PipeWidth, g.pipes[i].bottomH,
	}

	return checkCollision(birdBounds, pipeBounds) || checkCollision(birdBounds, bottomPipeBounds)
}

func checkCollision(rect1, rect2 []float64) bool {
	return rect1[0] < rect2[0]+rect2[2] &&
		rect1[0]+rect1[2] > rect2[0] &&
		rect1[1] < rect2[1]+rect2[3] &&
		rect1[1]+rect1[3] > rect2[1]
}

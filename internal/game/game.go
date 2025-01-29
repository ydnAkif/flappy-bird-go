package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game represents the game state
type Game struct {
	bird      Bird
	pipes     []Pipe
	score     int
	gameOver  bool
	spaceJust bool
}

// NewGame creates a new instance of the game
func NewGame() *Game {
	g := &Game{
		bird: Bird{
			x:       50,
			y:       screenHeight / 2,
			gravity: 0.5,
		},
		pipes: make([]Pipe, 0),
	}
	g.addPipe()
	return g
}

func (g *Game) addPipe() {
	gapY := rand.Float64()*(screenHeight-200) + 100
	pipe := Pipe{
		x:       screenWidth,
		topH:    gapY - pipeGap/2,
		bottomH: screenHeight - (gapY + pipeGap/2),
		passed:  false,
	}
	g.pipes = append(g.pipes, pipe)
}

// Update handles game logic updates
func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			*g = *NewGame()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.spaceJust {
			g.bird.vy = -8
			g.spaceJust = true
			PlayJumpSound()
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

		if !g.pipes[i].passed && g.pipes[i].x+pipeWidth < g.bird.x {
			g.score++
			g.pipes[i].passed = true
			PlayScoreSound()
		}
	}

	// Remove off-screen pipes
	if len(g.pipes) > 0 && g.pipes[0].x < -pipeWidth {
		g.pipes = g.pipes[1:]
	}

	// Add new pipes
	if len(g.pipes) == 0 || g.pipes[len(g.pipes)-1].x < screenWidth-200 {
		g.addPipe()
	}

	// Check ceiling and floor collisions
	if g.bird.y < 0 || g.bird.y > screenHeight-birdSize {
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

	ebitenutil.DebugPrint(screen, "Score: "+string(rune(g.score+'0')))

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "\n\nGame Over!\nPress SPACE to restart")
	}
}

// Layout implements ebiten.Game's Layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) checkCollisionWithPipe(i int) bool {
	birdRect := []float64{g.bird.x, g.bird.y, birdSize, birdSize}
	topPipeRect := []float64{g.pipes[i].x, 0, pipeWidth, g.pipes[i].topH}
	bottomPipeRect := []float64{g.pipes[i].x, screenHeight - g.pipes[i].bottomH, pipeWidth, g.pipes[i].bottomH}

	return checkCollision(birdRect, topPipeRect) || checkCollision(birdRect, bottomPipeRect)
}

func checkCollision(rect1, rect2 []float64) bool {
	return rect1[0] < rect2[0]+rect2[2] &&
		rect1[0]+rect1[2] > rect2[0] &&
		rect1[1] < rect2[1]+rect2[3] &&
		rect1[1]+rect1[3] > rect2[1]
}

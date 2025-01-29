package game

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"

	"github.com/ydnAkif/flappy-bird-go/internal/config"
)

// Game states
const (
	StateTitle = iota
	StatePlaying
	StateGameOver
)

// Game represents the game state
type Game struct {
	state  int
	bird   *Bird
	pipes  []*Pipe
	score  int
	paused bool
	atlas  *text.Atlas
}

// NewGame creates a new game instance
func NewGame() (*Game, error) {
	// Initialize game
	g := &Game{
		state: StateTitle,
		pipes: make([]*Pipe, 0),
		score: 0,
		atlas: text.NewAtlas(basicfont.Face7x13, text.ASCII),
	}

	// Initialize bird
	bird, err := NewBird()
	if err != nil {
		return nil, fmt.Errorf("failed to create bird: %v", err)
	}
	g.bird = bird

	return g, nil
}

// Update updates the game state
func (g *Game) Update(win *pixelgl.Window, dt float64) {
	if g.paused {
		return
	}

	switch g.state {
	case StateTitle:
		if win.JustPressed(pixelgl.KeySpace) {
			g.state = StatePlaying
			g.bird.Jump()
		}
	case StatePlaying:
		g.bird.Update(dt)

		// Update pipes
		for i := len(g.pipes) - 1; i >= 0; i-- {
			g.pipes[i].Update(dt)
			if g.pipes[i].IsOffscreen() {
				g.pipes = append(g.pipes[:i], g.pipes[i+1:]...)
			}
		}

		// Add new pipe
		if len(g.pipes) == 0 || g.pipes[len(g.pipes)-1].X < float64(config.ScreenWidth-config.PipeInterval) {
			g.pipes = append(g.pipes, NewPipe())
		}

		// Handle input
		if win.JustPressed(pixelgl.KeySpace) {
			g.bird.Jump()
		}

		// Check collisions
		for _, pipe := range g.pipes {
			if pipe.Collides(g.bird.GetBounds()) {
				g.state = StateGameOver
				break
			}
			if pipe.PassedBy(g.bird.GetBounds()) {
				g.score++
			}
		}

	case StateGameOver:
		if win.JustPressed(pixelgl.KeyR) {
			g.Reset()
		}
	}
}

// Draw draws the game screen
func (g *Game) Draw(win *pixelgl.Window) {
	// Clear screen
	win.Clear(color.RGBA{135, 206, 235, 255}) // Sky blue background

	// Draw game elements
	g.bird.Draw(win)
	for _, pipe := range g.pipes {
		pipe.Draw(win)
	}

	// Draw score
	txt := text.New(pixel.V(10, float64(config.ScreenHeight-30)), g.atlas)
	txt.Color = color.Black
	fmt.Fprintf(txt, "Score: %d", g.score)
	txt.Draw(win, pixel.IM)

	// Draw game state text
	switch g.state {
	case StateTitle:
		txt := text.New(pixel.V(float64(config.ScreenWidth)/2-50, float64(config.ScreenHeight)/2), g.atlas)
		txt.Color = color.Black
		fmt.Fprintf(txt, "Press SPACE to start")
		txt.Draw(win, pixel.IM)
	case StateGameOver:
		txt := text.New(pixel.V(float64(config.ScreenWidth)/2-70, float64(config.ScreenHeight)/2), g.atlas)
		txt.Color = color.Black
		fmt.Fprintf(txt, "Game Over! Press R to restart")
		txt.Draw(win, pixel.IM)
	}
}

// Reset resets the game state
func (g *Game) Reset() {
	g.state = StateTitle
	g.score = 0
	g.pipes = make([]*Pipe, 0)
	g.bird, _ = NewBird()
}

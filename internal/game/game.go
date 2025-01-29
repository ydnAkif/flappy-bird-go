package game

import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"  // Updated import
	"github.com/ydnakif/flappy-bird-go/internal/assets"
)

type Game struct {
	bird        *Bird
	pipeManager *PipeManager
	state       GameState
	assetLoader *assets.AssetLoader
	background  *ebiten.Image
	bgX         float64
	score       int
}

func NewGame() (*Game, error) {
	assetLoader := assets.NewAssetLoader()

	// Create default background first
	bgImage := ebiten.NewImage(ScreenWidth, ScreenHeight)
	bgImage.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	// Try loading background image asynchronously
	go func() {
		if loaded, err := assetLoader.LoadImage("background.png"); err == nil {
			bgImage = loaded
		}
	}()

	// Load bird image
	birdImage, err := assetLoader.LoadImage("bird.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load bird image: %v", err)
	}

	// Load pipe image
	pipeImage, err := assetLoader.LoadImage("pipe.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load pipe image: %v", err)
	}

	game := &Game{
		bird:        NewBird(float64(ScreenWidth)/3, float64(ScreenHeight)/2, birdImage),
		pipeManager: NewPipeManager(pipeImage),
		state:       StateMenu,
		assetLoader: assetLoader,
		background:  bgImage,
		bgX:         0,
	}

	return game, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background with optimized scrolling
	if g.background != nil {
		bgWidth := float64(g.background.Bounds().Dx())

		// Draw multiple backgrounds to ensure smooth scrolling
		for i := 0; i < 3; i++ {
			op := &ebiten.DrawImageOptions{}
			x := g.bgX + (bgWidth * float64(i))
			if x < ScreenWidth {
				op.GeoM.Translate(x, 0)
				screen.DrawImage(g.background, op)
			}
		}
	}

	// Draw game elements
	if g.state == StatePlaying || g.state == StateGameOver {
		g.pipeManager.Draw(screen)
		g.bird.Draw(screen)
	}

	msg := fmt.Sprintf("FPS: %0.2f\nState: %d", ebiten.ActualFPS(), g.state)
	if g.state == StateGameOver {
		msg += "\nGame Over!\nPress SPACE to restart"
	}
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Update() error {
	// Smoother background scrolling
	if g.background != nil {
		g.bgX -= BackgroundScrollSpeed
		bgWidth := float64(g.background.Bounds().Dx())
		if g.bgX <= -bgWidth {
			g.bgX += bgWidth
		}
	}

	switch g.state {
	case StateMenu:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.state = StatePlaying
		}
	case StatePlaying:
		if g.bird != nil {
			g.bird.Update()
		}
		if g.pipeManager != nil {
			g.pipeManager.Update()
		}
	}
	return nil
}

// Layout implements ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Reset resets the game state for a new game
func (g *Game) Reset() {
	g.state = StatePlaying
	g.score = 0
	g.bird.Reset()
	g.pipeManager.Reset()
}

// checkCollisions checks for collisions between the bird and pipes or screen boundaries
func (g *Game) checkCollisions() bool {
	// Check screen boundaries
	if g.bird.y < 0 || g.bird.y > float64(ScreenHeight) {
		return true
	}

	// Check pipe collisions
	return g.pipeManager.CheckCollision(g.bird)
}

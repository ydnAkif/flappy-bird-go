package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"flappy-bird-go/internal/config"
	"flappy-bird-go/internal/game"
)

// Game constants
const (
	screenWidth  = 320
	screenHeight = 480
	birdSize     = 30
	pipeWidth    = 60
	pipeGap      = 100
)

func init() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle(config.GameTitle)

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

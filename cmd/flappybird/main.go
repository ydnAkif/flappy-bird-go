package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ydnAkif/flappy-bird-go/internal/config"
	"github.com/ydnAkif/flappy-bird-go/internal/game"
)

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

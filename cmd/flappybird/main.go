package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ydnakif/flappy-bird-go/internal/game"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle(game.GameTitle)

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

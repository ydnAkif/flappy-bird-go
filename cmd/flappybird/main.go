package main

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/ydnAkif/flappy-bird-go/internal/config"
	"github.com/ydnAkif/flappy-bird-go/internal/game"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  config.GameTitle,
		Bounds: pixel.R(0, 0, float64(config.ScreenWidth), float64(config.ScreenHeight)),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		g.Update(win, dt)
		g.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

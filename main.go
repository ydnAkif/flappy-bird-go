package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	// Load game assets
	LoadAssets()
	
	// Initialize audio
	InitAudio()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Flappy Bird")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

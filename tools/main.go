package main

import (
	"log"

	"github.com/ydnAkif/flappy-bird-go/tools/soundgen"
	"github.com/ydnAkif/flappy-bird-go/tools/spritegen"
)

func main() {
	log.Println("Generating sprites...")
	spritegen.GenerateSprites()
	
	log.Println("Generating sound effects...")
	soundgen.GenerateSounds()
	
	log.Println("Done!")
}

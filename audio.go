package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	sampleRate = 44100
)

var (
	audioContext *audio.Context
	jumpPlayer   *audio.Player
	scorePlayer  *audio.Player
	hitPlayer    *audio.Player
)

//go:embed assets/audio/jump.wav
var jumpSound []byte

//go:embed assets/audio/score.wav
var scoreSound []byte

//go:embed assets/audio/hit.wav
var hitSound []byte

// InitAudio initializes the audio system and loads sound effects
func InitAudio() {
	audioContext = audio.NewContext(sampleRate)

	// Initialize jump sound
	jumpSound, err := wav.Decode(audioContext, bytes.NewReader(jumpSound))
	if err != nil {
		log.Printf("Warning: could not load jump sound: %v", err)
		return
	}
	jumpPlayer, err = audio.NewPlayer(audioContext, jumpSound)
	if err != nil {
		log.Printf("Warning: could not create jump player: %v", err)
		return
	}

	// Initialize score sound
	scoreSound, err := wav.Decode(audioContext, bytes.NewReader(scoreSound))
	if err != nil {
		log.Printf("Warning: could not load score sound: %v", err)
		return
	}
	scorePlayer, err = audio.NewPlayer(audioContext, scoreSound)
	if err != nil {
		log.Printf("Warning: could not create score player: %v", err)
		return
	}

	// Initialize hit sound
	hitSound, err := wav.Decode(audioContext, bytes.NewReader(hitSound))
	if err != nil {
		log.Printf("Warning: could not load hit sound: %v", err)
		return
	}
	hitPlayer, err = audio.NewPlayer(audioContext, hitSound)
	if err != nil {
		log.Printf("Warning: could not create hit player: %v", err)
		return
	}
}

// PlayJumpSound plays the jump sound effect
func PlayJumpSound() {
	if jumpPlayer != nil {
		jumpPlayer.Rewind()
		jumpPlayer.Play()
	}
}

// PlayScoreSound plays the score sound effect
func PlayScoreSound() {
	if scorePlayer != nil {
		scorePlayer.Rewind()
		scorePlayer.Play()
	}
}

// PlayHitSound plays the hit sound effect
func PlayHitSound() {
	if hitPlayer != nil {
		hitPlayer.Rewind()
		hitPlayer.Play()
	}
}

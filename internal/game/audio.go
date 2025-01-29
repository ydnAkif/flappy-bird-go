package game

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

func init() {
	audioContext = audio.NewContext(sampleRate)
}

//go:embed internal/assets/audio/jump.wav
var jumpSound []byte

//go:embed internal/assets/audio/score.wav
var scoreSound []byte

//go:embed internal/assets/audio/hit.wav
var hitSound []byte

// LoadAudio loads all audio files
func LoadAudio() error {
	// Load jump sound
	jumpData, err := loadWavFile(jumpSound)
	if err != nil {
		return err
	}
	jumpPlayer, err = audio.NewPlayer(audioContext, jumpData)
	if err != nil {
		return err
	}

	// Load hit sound
	hitData, err := loadWavFile(hitSound)
	if err != nil {
		return err
	}
	hitPlayer, err = audio.NewPlayer(audioContext, hitData)
	if err != nil {
		return err
	}

	// Load score sound
	scoreData, err := loadWavFile(scoreSound)
	if err != nil {
		return err
	}
	scorePlayer, err = audio.NewPlayer(audioContext, scoreData)
	if err != nil {
		return err
	}

	return nil
}

// PlayJumpSound plays the jump sound effect
func PlayJumpSound() {
	if jumpPlayer != nil {
		jumpPlayer.Rewind()
		jumpPlayer.Play()
	}
}

// PlayHitSound plays the hit sound effect
func PlayHitSound() {
	if hitPlayer != nil {
		hitPlayer.Rewind()
		hitPlayer.Play()
	}
}

// PlayScoreSound plays the score sound effect
func PlayScoreSound() {
	if scorePlayer != nil {
		scorePlayer.Rewind()
		scorePlayer.Play()
	}
}

func loadWavFile(data []byte) (*wav.Stream, error) {
	wavReader, err := wav.Decode(audioContext, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return wavReader, nil
}

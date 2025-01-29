package game

import (
	"bytes"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	sampleRate = 44100
)

var (
	audioContext *audio.Context
	jumpPlayer   *audio.Player
	hitPlayer    *audio.Player
	scorePlayer  *audio.Player
)

//go:embed assets/jump.wav
var jumpSound []byte

//go:embed assets/score.wav
var scoreSound []byte

//go:embed assets/hit.wav
var hitSound []byte

func init() {
	audioContext = audio.NewContext(sampleRate)
}

// LoadAudio loads all audio files
func LoadAudio() error {
	var err error

	// Load jump sound
	jumpPlayer, err = loadWavPlayer(jumpSound)
	if err != nil {
		return err
	}

	// Load hit sound
	hitPlayer, err = loadWavPlayer(hitSound)
	if err != nil {
		return err
	}

	// Load score sound
	scorePlayer, err = loadWavPlayer(scoreSound)
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

func loadWavPlayer(data []byte) (*audio.Player, error) {
	d, err := wav.Decode(audioContext, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	player, err := audio.NewPlayer(audioContext, d)
	if err != nil {
		return nil, err
	}

	return player, nil
}

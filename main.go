package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var assets embed.FS

const (
	screenWidth  = 320
	screenHeight = 480
	birdSize     = 30
	pipeWidth    = 60
	pipeGap      = 100
	gameTitle    = "Flappy Bird"
	sampleRate   = 44100
	audioSampleRate = 44100
)

// Global audio context
var globalAudioContext *audio.Context

type Game struct {
	bird       Bird
	pipes      []Pipe
	score      int
	gameOver   bool
	spaceJust  bool
	difficulty float64

	// Görsel varlıklar
	birdImg    *ebiten.Image
	pipeImg    *ebiten.Image
	bgImg      *ebiten.Image
	
	// Ses varlıkları
	audioContext *audio.Context
	jumpSound    *audio.Player
	hitSound     *audio.Player
	scoreSound   *audio.Player
}

type Bird struct {
	x, y    float64
	vy      float64
	gravity float64
}

type Pipe struct {
	x        float64
	topH     float64
	bottomH  float64
	passed   bool
}

func loadImage(name string) *ebiten.Image {
	imgFile, err := assets.ReadFile("assets/" + name)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgFile))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func loadSound(audioContext *audio.Context, name string) *audio.Player {
	soundFile, err := assets.ReadFile("assets/" + name)
	if err != nil {
		log.Fatal(err)
	}

	decoded, err := wav.Decode(audioContext, bytes.NewReader(soundFile))
	if err != nil {
		log.Fatal(err)
	}

	player, err := audioContext.NewPlayer(decoded)
	if err != nil {
		log.Fatal(err)
	}

	return player
}

func NewGame() *Game {
	// Audio context'i bir kere oluştur
	if globalAudioContext == nil {
		globalAudioContext = audio.NewContext(audioSampleRate)
	}

	g := &Game{
		bird: Bird{
			x:       50,
			y:       screenHeight / 2,
			gravity: 0.5,
		},
		pipes:        make([]Pipe, 0),
		difficulty:   1.0,
		audioContext: globalAudioContext,
	}

	// Geçici olarak basit görüntüler oluştur
	g.birdImg = ebiten.NewImage(birdSize, birdSize)
	g.birdImg.Fill(color.RGBA{0xff, 0xff, 0x00, 0xff})

	g.pipeImg = ebiten.NewImage(pipeWidth, screenHeight)
	g.pipeImg.Fill(color.RGBA{0x00, 0x80, 0x00, 0xff})

	g.bgImg = ebiten.NewImage(screenWidth, screenHeight)
	g.bgImg.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	// Ses özelliklerini geçici olarak devre dışı bırak
	g.jumpSound = nil
	g.hitSound = nil
	g.scoreSound = nil

	g.addPipe()
	return g
}

func (g *Game) addPipe() {
	gapY := rand.Float64()*(screenHeight-200) + 100
	pipe := Pipe{
		x:        screenWidth,
		topH:     gapY - pipeGap/2,
		bottomH:  screenHeight - (gapY + pipeGap/2),
		passed:   false,
	}
	g.pipes = append(g.pipes, pipe)
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			newGame := NewGame()
			// Ses bağlamını kopyala
			newGame.audioContext = g.audioContext
			*g = *newGame
		}
		return nil
	}

	// Increase difficulty based on score
	g.difficulty = 1.0 + float64(g.score)/20.0

	// Bird control with updated velocity
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.spaceJust {
			g.bird.vy = -8 * g.difficulty
			g.spaceJust = true
			// Ses çalma kısmını kontrol et
			if g.jumpSound != nil {
				g.jumpSound.Rewind()
				g.jumpSound.Play()
			}
		}
	} else {
		g.spaceJust = false
	}

	g.bird.vy += g.bird.gravity
	g.bird.y += g.bird.vy

	// Updated pipe speed based on difficulty
	for i := range g.pipes {
		g.pipes[i].x -= 2 * g.difficulty

		// Çarpışma kontrolü
		birdRect := []float64{g.bird.x, g.bird.y, birdSize, birdSize}
		topPipeRect := []float64{g.pipes[i].x, 0, pipeWidth, g.pipes[i].topH}
		bottomPipeRect := []float64{g.pipes[i].x, screenHeight - g.pipes[i].bottomH, pipeWidth, g.pipes[i].bottomH}

		if checkCollision(birdRect, topPipeRect) || checkCollision(birdRect, bottomPipeRect) {
			g.gameOver = true
			if g.hitSound != nil {
				g.hitSound.Rewind()
				g.hitSound.Play()
			}
		}

		// Skor kontrolü
		if !g.pipes[i].passed && g.pipes[i].x+pipeWidth < g.bird.x {
			g.score++
			g.pipes[i].passed = true
			if g.scoreSound != nil {
				g.scoreSound.Rewind()
				g.scoreSound.Play()
			}
		}
	}

	// Ekrandan çıkan boruları sil
	if len(g.pipes) > 0 && g.pipes[0].x < -pipeWidth {
		g.pipes = g.pipes[1:]
	}

	// Yeni boru ekle
	if len(g.pipes) == 0 || g.pipes[len(g.pipes)-1].x < screenWidth-200 {
		g.addPipe()
	}

	// Zemin ve tavan çarpışma kontrolü
	if g.bird.y < 0 || g.bird.y > screenHeight-birdSize {
		g.gameOver = true
	}

	return nil
}

func checkCollision(rect1, rect2 []float64) bool {
	return rect1[0] < rect2[0]+rect2[2] &&
		rect1[0]+rect1[2] > rect2[0] &&
		rect1[1] < rect2[1]+rect2[3] &&
		rect1[1]+rect1[3] > rect2[1]
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Arkaplanı çiz
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.bgImg, op)

	// Kuşu çiz
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bird.x, g.bird.y)
	screen.DrawImage(g.birdImg, op)

	// Boruları çiz
	for _, pipe := range g.pipes {
		// Üst boru
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, -1) // Üst boru için ters çevir
		op.GeoM.Translate(pipe.x, pipe.topH)
		screen.DrawImage(g.pipeImg, op)

		// Alt boru
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pipe.x, screenHeight-pipe.bottomH)
		screen.DrawImage(g.pipeImg, op)
	}

	// Improved score display
	scoreText := fmt.Sprintf("Score: %d", g.score)
	ebitenutil.DebugPrint(screen, scoreText)

	if g.gameOver {
		gameOverText := fmt.Sprintf("\n\nGame Over!\nFinal Score: %d\nPress SPACE to restart", g.score)
		ebitenutil.DebugPrint(screen, gameOverText)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(gameTitle)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

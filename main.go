package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 480
	birdSize     = 30
	pipeWidth    = 60
	pipeGap      = 100
)

type Game struct {
	bird      Bird
	pipes     []Pipe
	score     int
	gameOver  bool
	spaceJust bool
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

func NewGame() *Game {
	g := &Game{
		bird: Bird{
			x:       50,
			y:       screenHeight / 2,
			gravity: 0.5,
		},
		pipes: make([]Pipe, 0),
	}
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
			*g = *NewGame()
		}
		return nil
	}

	// Kuş kontrolü
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.spaceJust {
			g.bird.vy = -8
			g.spaceJust = true
		}
	} else {
		g.spaceJust = false
	}

	g.bird.vy += g.bird.gravity
	g.bird.y += g.bird.vy

	// Boru hareketi ve çarpışma kontrolü
	for i := range g.pipes {
		g.pipes[i].x -= 2

		// Çarpışma kontrolü
		birdRect := []float64{g.bird.x, g.bird.y, birdSize, birdSize}
		topPipeRect := []float64{g.pipes[i].x, 0, pipeWidth, g.pipes[i].topH}
		bottomPipeRect := []float64{g.pipes[i].x, screenHeight - g.pipes[i].bottomH, pipeWidth, g.pipes[i].bottomH}

		if checkCollision(birdRect, topPipeRect) || checkCollision(birdRect, bottomPipeRect) {
			g.gameOver = true
		}

		// Skor kontrolü
		if !g.pipes[i].passed && g.pipes[i].x+pipeWidth < g.bird.x {
			g.score++
			g.pipes[i].passed = true
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
	// Arkaplan
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	// Kuş
	ebitenutil.DrawRect(screen, g.bird.x, g.bird.y, birdSize, birdSize, color.RGBA{0xff, 0xff, 0x00, 0xff})

	// Borular
	for _, pipe := range g.pipes {
		ebitenutil.DrawRect(screen, pipe.x, 0, pipeWidth, pipe.topH, color.RGBA{0x00, 0x80, 0x00, 0xff})
		ebitenutil.DrawRect(screen, pipe.x, screenHeight-pipe.bottomH, pipeWidth, pipe.bottomH, color.RGBA{0x00, 0x80, 0x00, 0xff})
	}

	// Skor
	ebitenutil.DebugPrint(screen, "Score: "+string(rune(g.score+'0')))

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "\n\nGame Over!\nPress SPACE to restart")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Flappy Bird")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

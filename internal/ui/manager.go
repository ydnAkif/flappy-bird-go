package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	score int
}

func NewManager() *Manager {
	return &Manager{
		score: 0,
	}
}

func (m *Manager) Draw(screen *ebiten.Image) {
	// TODO: Implement UI drawing
}

func (m *Manager) UpdateScore(score int) {
	m.score = score
}

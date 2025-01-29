package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	images map[string]*ebiten.Image
}

func NewManager() *Manager {
	return &Manager{
		images: make(map[string]*ebiten.Image),
	}
}

func (m *Manager) LoadImage(path string) error {
	// TODO: Implement image loading
	return nil
}

func (m *Manager) GetImage(name string) *ebiten.Image {
	return m.images[name]
}

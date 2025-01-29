package assets

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// AssetLoader handles loading and caching game assets
type AssetLoader struct {
	imageCache map[string]*ebiten.Image
}

// NewAssetLoader creates a new asset loader instance
func NewAssetLoader() *AssetLoader {
	return &AssetLoader{
		imageCache: make(map[string]*ebiten.Image),
	}
}

// LoadImage loads an image from the assets directory
func (l *AssetLoader) LoadImage(name string) (*ebiten.Image, error) {
	if img, exists := l.imageCache[name]; exists {
		return img, nil
	}

	path := filepath.Join("assets", "images", name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("image file not found: %s", path)
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load image %s: %v", path, err)
	}

	l.imageCache[name] = img
	return img, nil
}

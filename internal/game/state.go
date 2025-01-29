package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// GameState represents different states of the game
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
)

// StateManager handles game state transitions and rendering
type StateManager struct {
	currentState GameState
	score        int
	highScore    int
}

// NewStateManager creates a new state manager instance
func NewStateManager() *StateManager {
	return &StateManager{
		currentState: StateMenu,
		score:        0,
		highScore:    0,
	}
}

// Update handles state-specific logic updates
func (s *StateManager) Update() error {
	switch s.currentState {
	case StateMenu:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			s.currentState = StatePlaying
		}
	case StatePlaying:
		if ebiten.IsKeyPressed(ebiten.KeyP) {
			s.currentState = StatePaused
		}
	case StatePaused:
		if ebiten.IsKeyPressed(ebiten.KeyP) {
			s.currentState = StatePlaying
		}
	case StateGameOver:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			s.Reset()
		}
	}
	return nil
}

// Draw handles state-specific rendering
func (s *StateManager) Draw(screen *ebiten.Image) {
	switch s.currentState {
	case StateMenu:
		// TODO: Draw menu screen
	case StatePlaying:
		// TODO: Draw game elements
	case StatePaused:
		// TODO: Draw pause screen
	case StateGameOver:
		// TODO: Draw game over screen
	}
}

// Reset resets the game state for a new game
func (s *StateManager) Reset() {
	s.currentState = StatePlaying
	if s.score > s.highScore {
		s.highScore = s.score
	}
	s.score = 0
}

// GetState returns the current game state
func (s *StateManager) GetState() GameState {
	return s.currentState
}

// SetState changes the current game state
func (s *StateManager) SetState(state GameState) {
	s.currentState = state
}

// UpdateScore updates the current score
func (s *StateManager) UpdateScore(points int) {
	s.score += points
}

// GetScore returns the current score
func (s *StateManager) GetScore() int {
	return s.score
}

// GetHighScore returns the high score
func (s *StateManager) GetHighScore() int {
	return s.highScore
}

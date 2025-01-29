package audio

type Manager struct {
	sounds map[string][]byte
}

func NewManager() *Manager {
	return &Manager{
		sounds: make(map[string][]byte),
	}
}

func (m *Manager) LoadSound(name string, data []byte) {
	m.sounds[name] = data
}

func (m *Manager) PlaySound(name string) error {
	// TODO: Implement sound playing
	return nil
}

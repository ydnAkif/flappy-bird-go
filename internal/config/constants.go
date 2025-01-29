package config

// Window configuration
const (
	ScreenWidth  = 800
	ScreenHeight = 600
	GameTitle    = "Flappy Bird Go"
)

// Game configuration
const (
	BirdSize     = 30
	PipeWidth    = 50
	PipeGap      = 150
	BirdGravity  = 0.5
	BirdJumpVel  = -8
	PipeSpeed    = 2
	PipeInterval = 200
)

// Asset paths
const (
	SpritePath    = "internal/assets/sprites"
	AudioPath     = "internal/assets/audio"
	BirdSprite    = "bird.png"
	PipeSprite    = "pipe.png"
	BackgroundImg = "background.png"
)

package game

// Window constants
const (
	ScreenWidth  = 800
	ScreenHeight = 600
	GameTitle    = "Flappy Bird Go"
)

// Pipe constants
const (
	PipeWidth  = 52  // Width of pipe sprite
	PipeHeight = 320 // Height of pipe sprite
	PipeGap    = 150 // Gap between top and bottom pipes
	PipeSpeed  = 2.0 // Horizontal speed of pipes
)

// Game physics constants
const (
	Gravity   = 0.5
	JumpForce = -8.0
)

// Asset file names
const (
	BackgroundImagePath = "background.png"
	PipeImagePath       = "pipe.png"
)

// Background constants
const (
	BackgroundScrollSpeed = 0.5  // Reduced from default value for smoother scrolling
	BackgroundWidth      = 800   // Match screen width
	BackgroundHeight     = 600   // Match screen height
)

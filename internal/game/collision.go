package game

// BoundingBox represents a rectangular collision box
type BoundingBox struct {
	X, Y          float64
	Width, Height float64
}

// Intersects checks if this box intersects with another box
func (b BoundingBox) Intersects(other BoundingBox) bool {
	return b.X < other.X+other.Width &&
		b.X+b.Width > other.X &&
		b.Y < other.Y+other.Height &&
		b.Y+b.Height > other.Y
}

package types

type Vector2 struct {
	X, Y float64
}

// Add adds the other vector to this vector.
func (v *Vector2) Add(other *Vector2) {
	v.X += other.X
	v.Y += other.Y
}

package cross

// Vec is a 3D vector.
type Vec struct {
	X, Y, Z float64
}

// Cross returns a cross b.
func Cross(a, b Vec) Vec {
	return Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Cross2 returns a cross b.
func Cross2(a, b *Vec) *Vec {
	return &Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Cross3 sets v to a cross b and returns v.
func (v *Vec) Cross3(a, b *Vec) *Vec {
	v.X = a.Y*b.Z - a.Z*b.Y
	v.Y = a.Z*b.X - a.X*b.Z
	v.Z = a.X*b.Y - a.Y*b.X
	return v
}

// Cross4 sets v to a cross b.
func (v *Vec) Cross4(a, b *Vec) {
	v.X = a.Y*b.Z - a.Z*b.Y
	v.Y = a.Z*b.X - a.X*b.Z
	v.Z = a.X*b.Y - a.Y*b.X
}

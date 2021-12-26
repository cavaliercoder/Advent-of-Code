package geo3d

import "fmt"

type Cube struct {
	A Pos // close, bottom, left
	B Pos // far, top, right
}

func (c Cube) Width() int               { return c.B.X - c.A.X }
func (c Cube) Height() int              { return c.B.Y - c.A.Y }
func (c Cube) Depth() int               { return c.B.Z - c.A.Z }
func (c Cube) Volume() int              { return c.Width() * c.Height() * c.Depth() }
func (c Cube) In(other Cube) bool       { return c.A.In(other) && c.B.In(other) }
func (c Cube) Contains(other Cube) bool { return other.In(c) }

func (c Cube) Intersects(other Cube) bool {
	a, b := c, other
	if a.B.X < b.A.X || b.B.X < a.A.X {
		return false
	}
	if a.B.Y < b.A.Y || b.B.Y < a.A.Y {
		return false
	}
	if a.B.Z < b.A.Z || b.B.Z < a.A.Z {
		return false
	}
	return true
}

// Split divide two cubes into non-intersecting sub-cubes.
func (c Cube) Split(other Cube) []Cube {
	a := make([]Cube, 1, 8)
	a[0] = c
	if !c.Intersects(other) {
		return a
	}
	a = split(a, other.A)
	a = split(a, other.B)
	return a
}

func split(a []Cube, p Pos) []Cube {
	for i, c := range a {
		if p.X > c.A.X && p.X < c.B.X {
			c2 := c
			c.B.X, c2.A.X = p.X, p.X
			a[i] = c
			a = append(a, c2)
		}
	}
	for i, c := range a {
		if p.Y > c.A.Y && p.Y < c.B.Y {
			c2 := c
			c.B.Y, c2.A.Y = p.Y, p.Y
			a[i] = c
			a = append(a, c2)
		}
	}
	for i, c := range a {
		if p.Z > c.A.Z && p.Z < c.B.Z {
			c2 := c
			c.B.Z, c2.A.Z = p.Z, p.Z
			a[i] = c
			a = append(a, c2)
		}
	}
	return a
}

func (c Cube) String() string {
	return fmt.Sprintf(
		"x=%d..%d,y=%d..%d,z=%d..%d",
		c.A.X, c.B.X-1,
		c.A.Y, c.B.Y-1,
		c.A.Z, c.B.Z-1,
	)
}

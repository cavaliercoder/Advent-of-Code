package day10

import (
	"fmt"
	"math"
)

type Coord struct {
	x, y float64
}

func (c Coord) IsZero() bool {
	return c.x == 0 && c.y == 0
}

func (c Coord) Subtract(v Coord) Coord {
	return Coord{
		x: c.x - v.x,
		y: c.y - v.y,
	}
}

// Degrees returns the angle of these coordinates from the positive x axis.
func (c Coord) Degrees() float64 {
	Θ := math.Atan2(c.y, c.x)
	if Θ < 0 {
		Θ += 2 * math.Pi // normalize to range [0, 2π)
	}
	return (Θ * 360) / (2 * math.Pi)
}

// Distance returns the line-of-sight distance between this point and (0, 0).
func (c Coord) Distance() float64 {
	return math.Sqrt(c.x*c.x + c.y*c.y)
}

func (c Coord) String() string {
	return fmt.Sprintf("{%.0f, %.0f}", c.x, c.y)
}

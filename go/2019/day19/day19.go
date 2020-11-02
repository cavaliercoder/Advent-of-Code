package main

import (
	. "aoc/2019/common"
	"aoc/2019/intcode"
	"math"
)

type Drone struct {
	data intcode.Data
}

func NewDrone() (*Drone, error) {
	data, err := intcode.OpenData(Fixture("day19"))
	if err != nil {
		return nil, err
	}
	return &Drone{
		data: data,
	}, nil
}

// Scan runs the Drone Intcode program and returns 1 if the given coordinate
// is affected by the tractor beam.
func (c *Drone) Scan(x, y int) (v int, err error) {
	return intcode.Run(c.data, x, y)
}

// GetBeamAngle returns the angle of the left and right boundaries of the beam
// by scanning the beam at y.
func (c *Drone) GetBeamAngle(y int) (l, r float64) {
	var x, v int
	// find left side of the beam
	for ; v == 0; x++ {
		v, _ = c.Scan(x, y)
		if x > y {
			panic("no beam")
		}
	}
	l = math.Atan2(float64(y), float64(x))

	// find right side of the beam
	for ; v == 1; x++ {
		v, _ = c.Scan(x, y)
		if x > y {
			panic("no beam")
		}
	}
	r = math.Atan2(float64(y), float64(x))
	return
}

// IsBoxInBeam returns true if a box with dimensions n*n is covered within the
// tractor beam.
func (c *Drone) IsBoxInBeam(x, y, n int) bool {
	if v, _ := c.Scan(x, y); v == 0 {
		return false
	}
	if v, _ := c.Scan(x+n-1, y); v == 0 {
		return false
	}
	if v, _ := c.Scan(x, y+n-1); v == 0 {
		return false
	}
	if v, _ := c.Scan(x+n-1, y+n-1); v == 0 {
		return false
	}
	return true
}

// AdjustBox tries to nudge a box of n*n dimensions at x, y as far up and to the
// left as it can while remaining inside the beam. The purpose is to compensate
// for rounding errors between the quantization of the beam and the box to whole
// numbers.
func (c *Drone) AdjustBox(x, y, n int) (int, int) {
	moved := true
	for moved {
		moved = false
		for i := 1; i < 10; i++ {
			for j := 1; j < 10; j++ {
				if c.IsBoxInBeam(x-i, y-j, n) {
					x -= i
					y -= j
					moved = true
				}
			}
		}
	}
	return x, y
}

// DistanceToBox returns the X, Y coordinate of the top left of a box with n*n
// dimensions at the minimum distance to the tractor beam where it is completely
// covered by the beam.
func (c *Drone) DistanceToBox(n int) (x int, y int) {
	N := float64(n)
	Θ1, Θ2 := c.GetBeamAngle(10000)

	// Computing X is "simple" trignometry. We know there is some value of x at
	// the left of the beam where its y value less n is on the right side of the
	// beam.
	// We know that x₀ = y / tan(Θ) and y = x.tan(Θ).
	// If x₀, y₀ is the top-left of the box, x₁, y₁ is the bottom left and x₂, y₂
	// is the top right then:
	//
	// x₀ = x₁ = x₂ - n
	// y₀ = y₁ - n = y₂
	//
	// We can convert an x to its y and vice-versa with:
	//
	// x = y/tan(Θ)
	// y = x.tan(Θ)
	//
	// Let's solve for x₁:
	// x₁ = x₂ - n
	// x₁ = y₂/tan(Θ₂) - n
	// x₁ = y₁/tan(Θ₂) - n/tan(Θ₂) - n
	// x₁ = x₁/tan(Θ₂) * tan(Θ₁)/tan(Θ₂) - n/tan(Θ₂) - n
	// x₁.tan(Θ₂) = x₁.tan(Θ₁) - n.tan(Θ₂) - n
	// x₁.tan(Θ₁) - x₁.tan(Θ₂) = n.tan(Θ₂) + n
	// x₁(x₁.tan(Θ₁) - tan(Θ₂)) = n.tan(Θ₂) + n
	// x₁ = n.tan(Θ₂) + n / (x₁.tan(Θ₁) - tan(Θ₂))
	x = int((N*math.Tan(Θ2) + N) / (math.Tan(Θ1) - math.Tan(Θ2)))

	// now that we know x₁, simply y₂ = y₁ - n = x₁.tan(Θ₁) - n
	y = int((float64(x) * math.Tan(Θ1)) - N)
	return c.AdjustBox(x, y, n)
}

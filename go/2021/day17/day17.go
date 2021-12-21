package day17

import (
	"math"

	. "aoc2021"
)

type Rect struct {
	A, B Pos
}

func (r Rect) ContainsX(x int) bool { return r.A.X <= x && x <= r.B.X }
func (r Rect) ContainsY(y int) bool { return r.A.Y >= y && y >= r.B.Y }
func (r Rect) Contains(p Pos) bool  { return r.ContainsX(p.X) && r.ContainsY(p.Y) }

// Limit computes the highest positive value for initial velocity u, assuming
// constant acceleration of -1.
func Limit(u int) int {
	// Since acceleration is a constant of -1, lim = (Δx² + Δx) / 2
	return ((u * u) + u) / 2
}

// VelocityWithLimit returns the highest initial velocity allowed to stay below
// the given positive limit.
func VelocityWithLimit(lim int) int {
	// Factor the quadratic and return the positive solution.
	//
	// In 0 = ax² + bx + c, we known a and b are 1, so the factors of c must be
	// a positive and a negative with a sum of 1. We also know the function to
	// compute the limit uses a square, so we can just find the ceiling and
	// floor of the square root of c to find its factors. Their product will be
	// c and sum will be 1.
	//
	// lim = (Δx² + Δx) / 2
	//   0 = Δx² + Δx - 2lim
	//   0 = (Δx - floor(sqrt(2lim))) (Δx + ceil(sqrt(2lim)))
	return int(math.Floor(math.Sqrt(float64(lim * 2))))
}

// TrickShot returns the y position of the highest possible azimuth of any shot
// that can land in the target.
func TrickShot(target Rect) (y int) {
	// We recognize that all positive values of Y-velocity will eventually
	// result in a y-value of 0 before plunging below 0 with a velocity that is
	// the inverse of the initial velocity less 1. The best possible velocity to
	// land in the target bounds will move from y=0 to the lowest edge of the
	// target in one move. In other words, the best initial velocity of y is
	// -(min-y(Target)) - 1. Once known, just find the azimuth (limit).
	return Limit(-target.B.Y - 1)
}

// EnumerateTrajectories returns every possible initial trajectory that will
// land in r at any T.
func EnumerateTrajectories(r Rect) map[Pos]struct{} {
	m := make(map[Pos]struct{}, 64)

	// Lowest X-velocity allowed is the first with a limit that lands in bounds.
	// Any lower won't reach the target. Highest X-velocity allowed lands at the
	// farthest edge of the target at T1.
	minΔX, maxΔX := VelocityWithLimit(r.A.X), r.B.X

	// Lowest Y-velocity allowed will land at the lowest edge of the target at
	// T1. Highest Y-velocity allowed will be our best trick shot. Any higher
	// and we overshoot the target on descent.
	minΔY, maxΔY := r.B.Y, -r.B.Y-1

	// Maximum time we can take is the time it takes for our trick shot to
	// ascend, descend, land on 0 and then one plunge below 0.
	maxT := maxΔY*2 + 2

	// enumerate all possible initial y-velocities (Δy0)
	for Δy0 := minΔY; Δy0 <= maxΔY; Δy0++ {
		y := 0
		Δy := Δy0
		for T := 0; T < maxT; T++ {
			y += Δy
			Δy--
			if y < r.B.Y {
				break // overshot
			}
			if y > r.A.Y {
				continue // undershot
			}

			// try all possible initial x-velocities (Δx0) at time T
			for Δx0 := minΔX; Δx0 <= maxΔX; Δx0++ {
				x := 0
				Δx := Δx0
				for i := 0; i < T+1; i++ {
					x += Δx
					Δx--
					if Δx < 0 {
						Δx = 0
					}
				}
				if r.ContainsX(x) {
					// Δx0 lands in the target at time T
					m[Pos{X: Δx0, Y: Δy0}] = struct{}{}
				}
			}
		}
	}
	return m
}

package day17

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/geo"
)

var FixtureExample = geo.Rect{
	A: geo.Pos{X: 20, Y: -5},
	B: geo.Pos{X: 30, Y: -10},
}

// From day17.txt
var FixtureInput = geo.Rect{
	A: geo.Pos{X: 235, Y: -62},
	B: geo.Pos{X: 259, Y: -118},
}

func TestVelocityWithLimit(t *testing.T) {
	for u := 0; u < 10000; u++ {
		// find maximum x value
		lim := Limit(u)

		// validate limit by hand
		x := 0
		Δx := u
		for {
			if x == lim {
				break
			}
			x += Δx
			Δx--
		}
		assert.Int(t, lim, x, "bad final x position")

		// ensure we can solve back to the same initial velocity
		assert.Int(t, u, VelocityWithLimit(lim), "bad initial velocity from limit")
	}
}

func TestExample(t *testing.T) {
	assert.Int(t, 45, TrickShot(FixtureExample), "bad azimuth")
	assert.Int(
		t,
		112,
		len(EnumerateTrajectories(FixtureExample)),
		"bad initial trajectory count",
	)
}

func TestPart1(t *testing.T) {
	assert.Int(t, 6903, TrickShot(FixtureInput), "bad azimuth")
}

func TestPart2(t *testing.T) {
	assert.Int(
		t,
		2351,
		len(EnumerateTrajectories(FixtureInput)),
		"bad initial trajectory count",
	)
}

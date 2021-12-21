package day17

import (
	"testing"

	. "aoc2021"
)

var FixtureExample = Rect{
	A: Pos{X: 20, Y: -5},
	B: Pos{X: 30, Y: -10},
}

// From day17.txt
var FixtureInput = Rect{
	A: Pos{X: 235, Y: -62},
	B: Pos{X: 259, Y: -118},
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
		AssertInt(t, lim, x, "bad final x position")

		// ensure we can solve back to the same initial velocity
		AssertInt(t, u, VelocityWithLimit(lim), "bad initial velocity from limit")
	}
}

func TestExample(t *testing.T) {
	AssertInt(t, 45, TrickShot(FixtureExample), "bad azimuth")
	AssertInt(
		t,
		112,
		len(EnumerateTrajectories(FixtureExample)),
		"bad initial trajectory count",
	)
}

func TestPart1(t *testing.T) {
	AssertInt(t, 6903, TrickShot(FixtureInput), "bad azimuth")
}

func TestPart2(t *testing.T) {
	AssertInt(
		t,
		2351,
		len(EnumerateTrajectories(FixtureInput)),
		"bad initial trajectory count",
	)
}

package day24

import (
	"aoc"
	"strings"
	"testing"
)

func TestPath1(t *testing.T) {
	f, err := aoc.OpenFixture("day24")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	eris := ReadEris(f, false)
	seen := make(map[string]int)
	for {
		eris.Step()
		h := eris.SHA256()
		if _, ok := seen[h]; ok {
			aoc.AssertInt(t, 32573535, seen[h], "bad biodiversity rating")
			return
		}
		seen[h] = eris.BiodiversityRating()
	}
}

func TestPart2Example1(t *testing.T) {
	example := `....#
#..#.
#.?##
..#..
#....
`
	eris := ReadEris(strings.NewReader(example), true)
	for i := 0; i < 10; i++ {
		eris.Step()
	}
	aoc.AssertInt(t, 99, eris.RecursiveBugCount(), "bad bug count")
}

func TestPath2(t *testing.T) {
	f, err := aoc.OpenFixture("day24")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	eris := ReadEris(f, true)
	for i := 0; i < 200; i++ {
		eris.Step()
	}
	aoc.AssertInt(t, 1951, eris.RecursiveBugCount(), "bad bug count")
}

package day19

import (
	"testing"

	. "aoc"
)

func TestPart1(t *testing.T) {
	d, err := NewDrone()
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	v := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			v, err = d.Scan(x, y)
			if err != nil {
				panic(err)
			}
			n += v
		}
	}
	AssertInt(t, 169, n, "bad scan result")
}

func TestPart2(t *testing.T) {
	d, err := NewDrone()
	if err != nil {
		t.Fatal(err)
	}
	x, y := d.DistanceToBox(100)
	v := x*10000 + y
	AssertInt(t, 7001134, v, "bad x, y")
}

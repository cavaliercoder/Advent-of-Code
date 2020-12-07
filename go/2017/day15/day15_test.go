package day15

import (
	"testing"
)

const (
	MultiplierA uint64 = 4
	MultiplierB uint64 = 8
)

func TestGenerators(t *testing.T) {
	tests := []struct {
		generator Generator
		values    []uint64
	}{
		{&generator{16807, 65}, []uint64{1092455, 1181022009, 245556042, 1744312007, 1352636452}},
		{&generator{48271, 8921}, []uint64{430625591, 1233683848, 1431495498, 137874439, 285222916}},
		{&multipleGenerator{16807, 4, 65}, []uint64{1352636452, 1992081072, 530830436, 1980017072, 740335192}},
		{&multipleGenerator{48271, 8, 8921}, []uint64{1233683848, 862516352, 1159784568, 1616057672, 412269392}},
	}
	for _, test := range tests {
		for i := 0; i < len(test.values); i++ {
			actual := test.generator.Next()
			if actual != test.values[i] {
				t.Errorf("expected %d, got %d", test.values[i], actual)
			}
		}
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		SeedA, SeedB uint64
		Expect       int
	}{
		{65, 8921, 588},
		{873, 583, 631},
	}

	for _, test := range tests {
		a := &generator{16807, test.SeedA}
		b := &generator{48271, test.SeedB}
		actual := CountMatches(40000000, a, b)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		SeedA, SeedB uint64
		Expect       int
	}{
		{65, 8921, 309},
		{873, 583, 279},
	}

	for _, test := range tests {
		a := &multipleGenerator{16807, 4, test.SeedA}
		b := &multipleGenerator{48271, 8, test.SeedB}
		actual := CountMatches(5000000, a, b)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

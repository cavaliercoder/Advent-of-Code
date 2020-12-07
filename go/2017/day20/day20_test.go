package day20

import (
	"bytes"
	"testing"

	. "aoc"
)

var (
	example1 = []byte(`p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>
`)

	example2 = []byte(`p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>
`)
)

func TestOne(t *testing.T) {
	tests := []struct {
		Input  []byte
		Expect int
	}{
		{example1, 0},
		{MustReadFixture("day20"), 150},
	}
	for _, test := range tests {
		actual := NearestParticle(ParseParticles(bytes.NewReader(test.Input)))
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		Input  []byte
		Expect int
	}{
		{example2, 1},
		{MustReadFixture("day20"), 657},
	}
	for _, test := range tests {
		actual := CountSurvivors(ParseParticles(bytes.NewReader(test.Input)))
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

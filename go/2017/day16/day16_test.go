package day16

import (
	"bytes"
	"testing"

	. "aoc"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		d      *DanceFloor
		input  string
		expect []byte
	}{
		{NewDanceFloor(5), "s1,x3/4,pe/b", []byte("baedc")},
		{NewDanceFloor(16), string(MustReadFixture("day16")), []byte("kbednhopmfcjilag")},
	}
	for _, test := range tests {
		moves := ParseDanceMoves(test.input)
		test.d.Do(moves, 1)
		if !bytes.Equal(test.d.b, test.expect) {
			t.Errorf("expected %s, got %s", test.expect, test.d.b)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		d      *DanceFloor
		input  string
		expect []byte
	}{
		{NewDanceFloor(5), "s1,x3/4,pe/b", []byte("abcde")},
		{NewDanceFloor(16), string(MustReadFixture("day16")), []byte("fbmcgdnjakpioelh")},
	}
	for _, test := range tests {
		moves := ParseDanceMoves(test.input)
		test.d.Do(moves, 1e9)
		if !bytes.Equal(test.d.b, test.expect) {
			t.Errorf("expected %s, got %s", test.expect, test.d.b)
		}
	}
}

package main

import (
	. "aoc"
	"aoc/intcode"

	"testing"
)

func TestPart2(t *testing.T) {
	solution := []string{
		"A,B,A,C,A,B,A,C,B,C", // main()
		"R,4,L,12,L,8,R,4",    // A()
		"L,8,R,10,R,10,R,6",   // B()
		"R,4,R,10,L,12",       // C()
		"n",                   // no feed
	}
	data, err := intcode.OpenData(Fixture("day17"))
	if err != nil {
		t.Fatal(err)
	}
	data[0] = 2
	v, err := intcode.RunASCII(data, solution...)
	if err != nil {
		t.Fatal(err)
	}
	AssertInt(t, 785733, v, "bad dust count")
}

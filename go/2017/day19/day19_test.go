package day19

import (
	"bytes"
	"testing"

	. "aoc"
)

var example = []byte(`     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`)

func TestPartOne(t *testing.T) {
	tests := []struct {
		Input  []byte
		Expect []byte
	}{
		{example, []byte("ABCDEF")},
		{MustReadFixture("day19"), []byte("RUEDAHWKSM")},
	}
	for _, test := range tests {
		grid := ReadGrid(bytes.NewReader(test.Input))
		_, actual := Route(grid)
		if !bytes.Equal(actual, test.Expect) {
			t.Errorf("expected '%s', got '%s'", test.Expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  []byte
		Expect int
	}{
		{example, 38},
		{MustReadFixture("day19"), 17264},
	}
	for _, test := range tests {
		grid := ReadGrid(bytes.NewReader(test.Input))
		actual, _ := Route(grid)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

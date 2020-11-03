package day22

import (
	"strings"
	"testing"

	. "aoc/2019/common"
)

func TestPart1Example1(t *testing.T) {
	expect := "Deck{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}"
	input := `deal with increment 7
deal into new stack
deal into new stack
`
	d := NewDeck(10)
	d.Run(strings.NewReader(input))
	AssertString(t, expect, d.String(), "bad deck result")
}

func TestPart1Example2(t *testing.T) {
	expect := "Deck{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}"
	input := `cut 6
deal with increment 7
deal into new stack
`
	d := NewDeck(10)
	d.Run(strings.NewReader(input))
	AssertString(t, expect, d.String(), "bad deck result")
}
func TestPart1Example3(t *testing.T) {
	expect := "Deck{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}"
	input := `deal with increment 7
deal with increment 9
cut -2
`
	d := NewDeck(10)
	d.Run(strings.NewReader(input))
	AssertString(t, expect, d.String(), "bad deck result")
}

func TestPart1Example4(t *testing.T) {
	expect := "Deck{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}"
	input := `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1
`
	d := NewDeck(10)
	d.Run(strings.NewReader(input))
	AssertString(t, expect, d.String(), "bad deck result")
}

func TestPart1(t *testing.T) {
	f, err := OpenFixture("day22")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	d := NewDeck(10007)
	d.Run(f)
	AssertInt(t, 7860, d.IndexOf(2019), "wrong index of card 2019")
}

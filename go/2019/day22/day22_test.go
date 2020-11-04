package day22

import (
	"fmt"
	"strings"
	"testing"

	. "aoc"
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

func TestUndoDealIntoNewDeck(t *testing.T) {
	n := 1024
	f := UndoDealIntoNewDeck()
	for i := 0; i < n; i++ {
		d := NewDeck(n)
		d.DealIntoNewStack()
		sd := NewSuperDeck(int64(n), int64(i))
		f.Do(sd)
		AssertInt(t, int(d.Get(i)), int(sd.Index()), "bad rewind at index %d", i)
	}
}

func TestUndoCut(t *testing.T) {
	n := 1024
	for cut := -n + 1; cut < n; cut++ { // TODO: negatives!
		d := NewDeck(n)
		d.Cut(cut)
		for i := 0; i < n; i++ {
			sd := NewSuperDeck(int64(n), int64(i))
			UndoCut(int64(cut)).Do(sd)
			AssertInt(t, int(d.Get(i)), int(sd.Index()), "bad UndoCut(%d) @ %d", cut, i)
		}
	}
}

func TestUndoDealWithIncrement(t *testing.T) {
	n := 271 // must be prime
	for incr := 2; incr < n; incr++ {
		t.Run(fmt.Sprintf("WithIncrement%000d", incr), func(t *testing.T) {
			d := NewDeck(n)
			d.DealWithIncrement(incr)
			for i := 0; i < n; i++ {
				t.Run(fmt.Sprintf("TrackingIndex%00d", i), func(t *testing.T) {
					sd := NewSuperDeck(int64(n), int64(i))
					UndoDealWithIncrement(int64(incr)).Do(sd)
					AssertInt(t, int(d.Get(i)), int(sd.Index()), "bad card in %v", d)
				})
			}
		})
	}
}

func TestUndoShuffle(t *testing.T) {
	n := 10007

	// run standard deck shuffle once
	d := NewDeck(10007)
	func() {
		f, err := OpenFixture("day22")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		d.Run(f)
	}()

	// load super shuffles
	var shuffles []SuperDeckShuffle
	func() {
		f, err := OpenFixture("day22")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		shuffles = ReadSuperDeckShuffles(f)
	}()

	// test result for each i in n
	for i := 0; i < n; i++ {
		t.Run(fmt.Sprintf("WithIndex%0000d", i), func(t *testing.T) {
			sd := NewSuperDeck(int64(n), int64(i))
			v := sd.UndoShuffle(shuffles...)
			AssertInt(t, int(d.Get(i)), int(v), "bad card in %v", d)
		})
	}
}

func TestFindLoops(t *testing.T) {
	// load super shuffles
	var shuffles []SuperDeckShuffle
	func() {
		f, err := OpenFixture("day22")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		shuffles = ReadSuperDeckShuffles(f)
	}()

	var v int64
	sd := NewSuperDeck(int64(119315717514047), int64(2020))
	for i := int64(0); i < 101741582076661; i++ {
		v = sd.UndoShuffle(shuffles...)
	}
	t.Fatal(v)
}

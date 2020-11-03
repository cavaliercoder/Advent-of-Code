package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Card int64

type Deck struct {
	cards []Card
	buf   []Card
}

func NewDeck(n int) *Deck {
	d := &Deck{
		cards: make([]Card, n),
		buf:   make([]Card, n),
	}
	for i := 0; i < n; i++ {
		d.cards[i] = Card(i)
	}
	return d
}

func (d *Deck) String() string {
	b := &bytes.Buffer{}
	b.WriteString("Deck{")
	for i := 0; i < len(d.cards); i++ {
		if i == 0 {
			fmt.Fprintf(b, "%d", d.cards[i])
		} else {
			fmt.Fprintf(b, ", %d", d.cards[i])
		}
	}
	b.WriteString("}")
	return b.String()
}

func (d *Deck) IndexOf(card Card) int {
	for i := 0; i < len(d.cards); i++ {
		if d.cards[i] == card {
			return i
		}
	}
	return -1
}

func (d *Deck) Get(i int) Card {
	return d.cards[i]
}

func (d *Deck) DealIntoNewStack() {
	for i := 0; i < len(d.cards); i++ {
		d.buf[len(d.cards)-1-i] = d.cards[i]
	}
	d.cards, d.buf = d.buf, d.cards
}

func (d *Deck) Cut(n int) {
	if n < 0 {
		n = len(d.cards) + n
	}
	for i := 0; i < len(d.cards); i++ {
		j := (i + n) % len(d.cards)
		d.buf[i] = d.cards[j]
	}
	d.cards, d.buf = d.buf, d.cards
}

func (d *Deck) DealWithIncrement(n int) {
	for i := 0; i < len(d.cards); i++ {
		j := (i * n) % len(d.cards)
		d.buf[j] = d.cards[i]
	}
	d.cards, d.buf = d.buf, d.cards
}

func (d *Deck) Run(r io.Reader) {
	var v, n int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()

		// deal into new stack
		if s == "deal into new stack" {
			d.DealIntoNewStack()
			continue
		}

		// cut
		if n, _ = fmt.Sscanf(s, "cut %d", &v); n == 1 {
			d.Cut(v)
			continue
		}

		// deal with increment
		if n, _ = fmt.Sscanf(s, "deal with increment %d", &v); n == 1 {
			d.DealWithIncrement(v)
			continue
		}

		panic(s)
	}
}

type SuperDeck struct {
	size       int64
	trackIndex int64
}

func NewSuperDeck(size int64, trackIndex int64) *SuperDeck {
	return &SuperDeck{
		size:       size,
		trackIndex: trackIndex,
	}
}

func (d *SuperDeck) Index() int64 {
	return d.trackIndex
}

func (d *SuperDeck) UndoShuffle(shuffles ...SuperDeckShuffle) int64 {
	for i := 0; i < len(shuffles); i++ {
		shuffles[len(shuffles)-1-i].Do(d)
	}
	return d.Index()
}

type SuperDeckShuffle interface {
	Do(d *SuperDeck)
}

func ReadSuperDeckShuffles(r io.Reader) []SuperDeckShuffle {
	shuffles := make([]SuperDeckShuffle, 0, 64)
	var v, n int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()

		// deal into new stack
		if s == "deal into new stack" {
			shuffles = append(shuffles, UndoDealIntoNewDeck())
			continue
		}

		// cut
		if n, _ = fmt.Sscanf(s, "cut %d", &v); n == 1 {
			shuffles = append(shuffles, UndoCut(int64(v)))
			continue
		}

		// deal with increment
		if n, _ = fmt.Sscanf(s, "deal with increment %d", &v); n == 1 {
			shuffles = append(shuffles, UndoDealWithIncrement(int64(v)))
			continue
		}

		panic(s)
	}
	return shuffles
}

type SuperDeckShuffleFunc func(*SuperDeck)

func (f SuperDeckShuffleFunc) Do(d *SuperDeck) {
	f(d)
}

func UndoDealIntoNewDeck() SuperDeckShuffle {
	return SuperDeckShuffleFunc(func(d *SuperDeck) {
		d.trackIndex = d.size - 1 - d.trackIndex
	})
}

func UndoCut(n int64) SuperDeckShuffle {
	return SuperDeckShuffleFunc(func(d *SuperDeck) {
		if n < 0 {
			n = d.size + n
		}
		d.trackIndex = (n + d.trackIndex) % d.size
	})
}

func UndoDealWithIncrement(n int64) SuperDeckShuffle {
	return SuperDeckShuffleFunc(func(d *SuperDeck) {
		A := make([]int64, n)
		var v, j int64
		for i := int64(1); i < n; i++ {
			v += (d.size-j)/n + 1
			j = (v * n) % d.size
			A[j] = v
		}
		d.trackIndex = A[d.trackIndex%n] + d.trackIndex/n
	})
}

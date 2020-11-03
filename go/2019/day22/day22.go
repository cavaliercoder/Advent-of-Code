package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Card uint64

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

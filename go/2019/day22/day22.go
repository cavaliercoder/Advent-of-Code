package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/big"
)

var (
	Zero = big.NewInt(0)
	One  = big.NewInt(1)
	Two  = big.NewInt(2)
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

// SuperDeck implements arbitrarily large deck shuffles using linear
// polynomials/group theory. I know very little about these so I cargo-culted
// the following impressive solution:
// https://github.com/metalim/metalim.adventofcode.2019.python/blob/master/22_cards_shuffle.ipynb
type SuperDeck struct {
	L, a, b *big.Int
}

func ReadSuperDeck(r io.Reader, size int64) *SuperDeck {
	var n int64
	var tokens int
	L, a, b := big.NewInt(size), big.NewInt(1), big.NewInt(0)
	rules := make([]string, 0, 64)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}
	for i := 0; i < len(rules); i++ {
		s := rules[len(rules)-1-i]

		// deal into new stack
		if s == "deal into new stack" {
			// a = -a
			a.Neg(a)

			// b = L - b - 1
			b.Sub(L, b)
			b.Sub(b, One)
			continue
		}

		// cut
		if tokens, _ = fmt.Sscanf(s, "cut %d", &n); tokens == 1 {
			// b = (b + n) % size
			b.Add(b, big.NewInt(n))
			b.Mod(b, L)
			continue
		}

		// deal with increment
		if tokens, _ = fmt.Sscanf(s, "deal with increment %d", &n); tokens == 1 {
			// z = n^(L-2) % L
			var z, exp big.Int
			exp.Sub(L, Two)
			z.Exp(big.NewInt(n), &exp, L)

			// a = a * z % L
			a.Mul(a, &z)
			a.Mod(a, L)

			// b = b * z % L
			b.Mul(b, &z)
			b.Mod(b, L)
			continue
		}

		// unsupported
		panic(s)
	}
	return &SuperDeck{L: L, a: a, b: b}
}

func polyPow(a, b, n, L *big.Int) {
	if n.Cmp(Zero) == 0 {
		a.Set(One)
		b.Set(Zero)
		return
	}

	// if n%2 == 0
	var nMod2 big.Int
	nMod2.Mod(n, Two)
	if nMod2.Cmp(Zero) == 0 {
		// b = (a*b + b) % L
		var x big.Int
		x.Mul(a, b)
		x.Add(&x, b)
		x.Mod(&x, L)
		b.Set(&x)

		// a = a^2 % L
		a.Exp(a, Two, L)

		// n = n / 2
		n.Div(n, Two)
		polyPow(a, b, n, L)
		return
	}

	// c, d = polyPow(a, b, n-1, L)
	var c, d, nn big.Int
	c.Set(a)
	d.Set(b)
	nn.Sub(n, One)
	polyPow(&c, &d, &nn, L)

	// a = a*c % L
	c.Mul(&c, a)
	c.Mod(&c, L)

	// b = (a*d + b) % L
	d.Mul(&d, a)
	d.Add(&d, b)
	d.Mod(&d, L)

	a.Set(&c)
	b.Set(&d)
}

func (c *SuperDeck) Shuffle(n int64) {
	polyPow(c.a, c.b, big.NewInt(n), c.L)
}

func (c *SuperDeck) Get(i int64) int64 {
	// x = (i*a + b) % L
	v := big.NewInt(i)
	v.Mul(v, c.a)
	v.Add(v, c.b)
	v.Mod(v, c.L)
	return v.Int64()
}

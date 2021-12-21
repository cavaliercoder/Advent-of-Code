package day18

import (
	"bytes"
	"fmt"
	"io"
)

// Number is a snailfish number.
type Number struct {
	N    int
	L, R *Number
}

func (n *Number) Copy() *Number {
	if n == nil {
		return nil
	}
	return &Number{
		N: n.N,
		L: n.L.Copy(),
		R: n.R.Copy(),
	}
}

func (n *Number) IsRegular() bool {
	return n.L == nil && n.R == nil
}

// M returns the magnitude of the snailfish number.
func (n *Number) M() int {
	if n.IsRegular() {
		return n.N
	}
	return 3*n.L.M() + 2*n.R.M()
}

func (n *Number) reduce() {
	for {
		if n.explode(0, nil, nil) {
			continue
		}
		if n.split() {
			continue
		}
		return
	}
}

func (n *Number) explode(depth int, l, r *Number) (ok bool) {
	if n.IsRegular() {
		return false
	}
	if depth == 4 {
		if !n.L.IsRegular() || !n.R.IsRegular() {
			panic("why?")
		}
		if l != nil {
			for !l.IsRegular() {
				l = l.R
			}
			l.N += n.L.N
		}
		if r != nil {
			for !r.IsRegular() {
				r = r.L
			}
			r.N += n.R.N
		}
		n.N, n.L, n.R = 0, nil, nil
		return true
	}
	if ok := n.L.explode(depth+1, l, n.R); ok {
		return true
	}
	if ok := n.R.explode(depth+1, n.L, r); ok {
		return true
	}
	return false
}

func (n *Number) split() (ok bool) {
	if n.IsRegular() {
		if n.N > 9 {
			n.L = &Number{
				N: n.N / 2,
			}
			n.R = &Number{
				N: n.N/2 + n.N%2,
			}
			return true
		}
		return false
	}
	if ok := n.L.split(); ok {
		return true
	}
	if ok := n.R.split(); ok {
		return true
	}
	return false
}

func (n *Number) Format(w io.Writer) {
	if n.IsRegular() {
		fmt.Fprintf(w, "%d", n.N)
	} else {
		fmt.Fprint(w, "[")
		n.L.Format(w)
		fmt.Fprint(w, ",")
		n.R.Format(w)
		fmt.Fprint(w, "]")
	}
}

func (n *Number) String() string {
	b := new(bytes.Buffer)
	n.Format(b)
	return b.String()
}

// Add returns a new number that is the sum of multiple snailfish numbers.
func Add(A ...*Number) *Number {
	if len(A) == 0 {
		return nil
	}
	a := A[0].Copy()
	for _, b := range A[1:] {
		a = &Number{L: a, R: b.Copy()}
		a.reduce()
	}
	return a
}

// MaxMagnitude returns the maximum magnitude for the sum of any two pairs of
// snailfish numbers in O(n²).
func MaxMagnitude(A ...*Number) int {
	max := 0
	for _, a := range A {
		for _, b := range A {
			if a == b {
				continue
			}
			m := Add(a, b).M()
			if m > max {
				max = m
			}
		}
	}
	return max
}

func Parse(b []byte) *Number {
	c := &parser{
		data: b,
	}
	return c.Parse()
}

type parser struct {
	data []byte
	ptr  int
}

func (c *parser) Next() {
	if !c.EOF() {
		c.ptr++
	}
}

func (c *parser) Expect(b byte) {
	if c.B() != b {
		panic(fmt.Sprintf("expected: '%c', got '%c' at %d", b, c.B(), c.ptr))
	}
	c.Next()
}

func (c *parser) EOF() bool { return c.ptr >= len(c.data) }

func (c *parser) B() byte {
	if c.EOF() {
		return 0
	}
	return c.data[c.ptr]
}

func (c *parser) IsNumeric() bool {
	if c.EOF() {
		return false
	}
	return c.B() >= '0' && c.B() <= '9'
}

func (c *parser) Parse() *Number {
	n := &Number{}
	if c.IsNumeric() {
		for c.IsNumeric() {
			n.N *= 10
			n.N += int(c.data[c.ptr] - '0')
			c.Next()
		}
		return n
	}
	c.Expect('[')
	n.L = c.Parse()
	c.Expect(',')
	n.R = c.Parse()
	c.Expect(']')
	return n
}

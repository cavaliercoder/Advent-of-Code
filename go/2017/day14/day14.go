package day14

import (
	"fmt"
)

// KnotHash is adapted from the day 10 challenge and returns a hash of the
// given input.
func KnotHash(input []byte) []byte {
	input = append(input, 17, 31, 73, 47, 23)
	A := make([]byte, 256)
	for i := 0; i < len(A); i++ {
		A[i] = byte(i)
	}
	var pos, skip, a, b int
	for r := 0; r < 64; r++ {
		for _, l := range input {
			for i := byte(0); i < l/2; i++ {
				a = (pos + int(i)) % len(A)
				b = (pos + int(l-i) - 1) % len(A)
				A[a], A[b] = A[b], A[a]
			}
			pos = (pos + int(l) + skip) % len(A)
			skip++
		}
	}
	h := make([]byte, 16)
	for i := 0; i < 16; i++ {
		var b byte
		for j := 0; j < 16; j++ {
			b ^= A[(16*i)+j]
		}
		h[i] = b
	}
	return h
}

// A Grid is a bitmap representation of a disk which requires defragmentation.
type Grid struct {
	Width int
	Data  []byte
}

// NewGrid returns a grid of the given width and populates it with a hash of the
// given input string.
func NewGrid(n int, s string) *Grid {
	g := &Grid{
		Width: n,
		Data:  make([]byte, n*n/8),
	}
	for i := 0; i < n; i++ {
		b := KnotHash([]byte(fmt.Sprintf("%s-%d", s, i)))
		copy(g.Data[i*16:], b)
	}
	return g
}

// BitCount is my solution to Part One and returns the number of bits set in the
// Grid.
func (g *Grid) BitCount() int {
	count := 0
	for i := 0; i < len(g.Data); i++ {
		for j := byte(0); j < 8; j++ {
			if g.Data[i]&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
}

// GroupCount is my solution to Part Two and returns the number of bit groups,
// where all bits are adjacent to eachother.
func (g *Grid) GroupCount() int {
	count := 0
	for i := 0; i < len(g.Data)*8; i++ {
		count += g.squash(i)
	}
	return count
}

// squash returns 1 if the value at the given bit offset is not 0. The bit at
// the given offset will be set to zero and this function applied recursively to
// all of its adjacent neighbours, like a depth-first search.
func (g *Grid) squash(i int) int {
	var b, m byte
	var o int
	b = g.Data[i/8]
	m = byte(1 << uint(7-(i%8)))
	if b&m == 0 {
		return 0
	}
	g.Data[i/8] = b &^ m

	// above
	o = i - g.Width
	if o >= 0 {
		g.squash(o)
	}

	// right
	if i%g.Width < g.Width-1 {
		g.squash(i + 1)
	}

	// below
	o = i + g.Width
	if o < g.Width*g.Width {
		g.squash(o)
	}

	// left
	if i%g.Width > 0 {
		g.squash(i - 1)
	}

	return 1
}

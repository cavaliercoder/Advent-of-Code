package day19

import (
	"io"
	"io/ioutil"
)

type Grid struct {
	Width int
	Start int
	Data  []byte
}

func ReadGrid(r io.Reader) *Grid {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	g := &Grid{Data: b}
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '|':
			g.Start = i
		case '\n':
			g.Width = i + 1 // includes \n
			return g
		}
	}
	panic("bad grid")
}

// Route steps through the given Grid and returns both the number of steps taken
// and each of the letters encountered along the path. Its serves as my solution
// to Part One and Two, and completes in linear O(n) time.
func Route(g *Grid) (int, []byte) {
	b := make([]byte, 0)
	v := 0
	i, x, y := g.Start, 0, 1
	for {
		switch g.Data[i] {
		case ' ':
			return v, b

		case '|', '-':
			// keep moving

		case '+':
			x, y = y, x
			peek := i + (y * g.Width) + x
			if peek < 0 || peek >= len(g.Data) || g.Data[peek] == ' ' {
				x, y = -x, -y
			}

		default:
			b = append(b, g.Data[i])
		}

		i += (y * g.Width) + x
		v++
	}
}

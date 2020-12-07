package day22

import (
	"io"
	"io/ioutil"
)

type Grid struct {
	b          []byte
	p, w, x, y int
}

func ParseGrid(r io.Reader) *Grid {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	g := &Grid{
		b: make([]byte, 0, len(b)),
		y: -1,
	}
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '\n':
			if g.w == 0 {
				g.w = i
			}
		default:
			g.b = append(g.b, b[i])
		}
	}
	g.p = (g.w * (g.w / 2)) + g.w/2
	return g
}

// Grow allocates a new Grid, three times larger than the current Grid, with the
// original Grid copied into the exact centre of the new Grid. The Cursor
// position is also transposed onto the new Grid.
//
// A more elegant solution might first crop unused sides of the grid and only
// preallocate extra capacity in the direction of the cursor.
func (g *Grid) Grow() {
	n := g.w * 3
	b := make([]byte, n*n)
	for i := 0; i < len(b); i++ {
		b[i] = '.'
	}
	j := g.w*g.w*3 + g.w
	for i := 0; i < len(g.b); i++ {
		b[j] = g.b[i]
		if i%g.w == g.w-1 {
			j += 2 * g.w
		}
		j++
	}
	g.p = (g.w * g.w * 3) + (g.p/g.w)*(g.w*3) + g.w + (g.p % g.w)
	g.w = n
	g.b = b
}

// Infect is my solution to Part One. It takes n steps over the infinite Grid,
// toggling nodes between clean and infected. The number of infections that are
// created during these steps is returned.
func (g *Grid) Infect(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		switch g.b[g.p] {
		case '.':
			g.b[g.p] = '#'
			count++

			// turn left
			switch g.x {
			case 1:
				g.x, g.y = 0, -1
			case -1:
				g.x, g.y = 0, 1
			case 0:
				switch g.y {
				case 1:
					g.x, g.y = 1, 0
				case -1:
					g.x, g.y = -1, 0
				}
			}

		case '#':
			g.b[g.p] = '.'

			// turn right
			switch g.x {
			case 1:
				g.x, g.y = 0, 1
			case -1:
				g.x, g.y = 0, -1
			case 0:
				switch g.y {
				case 1:
					g.x, g.y = -1, 0
				case -1:
					g.x, g.y = 1, 0
				}
			}
		}

		// grow if cursor on boundary
		if g.p < g.w ||
			g.p/g.w == g.w-1 ||
			g.p%g.w == 0 ||
			g.p%g.w == g.w-1 {
			g.Grow()
		}

		// move cursor
		g.p += g.y*g.w + g.x
	}
	return count
}

// InfectEvolved is my solution to Part Two. It takes n steps over the infinite
// Grid, toggling nodes between clean, weakened, infected and flagged. The
// number of infections that are created during these steps is returned.
func (g *Grid) InfectEvolved(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		switch g.b[g.p] {
		case '.':
			g.b[g.p] = 'W'

			// turn left
			switch g.x {
			case 1:
				g.x, g.y = 0, -1
			case -1:
				g.x, g.y = 0, 1
			case 0:
				switch g.y {
				case 1:
					g.x, g.y = 1, 0
				case -1:
					g.x, g.y = -1, 0
				}
			}

		case '#':
			g.b[g.p] = 'F'

			// turn right
			switch g.x {
			case 1:
				g.x, g.y = 0, 1
			case -1:
				g.x, g.y = 0, -1
			case 0:
				switch g.y {
				case 1:
					g.x, g.y = -1, 0
				case -1:
					g.x, g.y = 1, 0
				}
			}

		case 'W':
			g.b[g.p] = '#'
			count++

		case 'F':
			g.b[g.p] = '.'
			g.x = -g.x
			g.y = -g.y
		}

		// grow if cursor on boundary
		if g.p < g.w ||
			g.p/g.w == g.w-1 ||
			g.p%g.w == 0 ||
			g.p%g.w == g.w-1 {
			g.Grow()
		}

		// move cursor
		g.p += g.y*g.w + g.x
	}
	return count
}

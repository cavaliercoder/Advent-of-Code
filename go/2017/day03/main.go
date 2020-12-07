/*
This is my solution for Day 3 of the Advent of Code 2017.

Input is read - one per line - from standard input. The shortest path solution
for each problem is printed to standard output.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer '%v': %v\n", scanner.Text(), err)
			os.Exit(1)
		}

		fmt.Printf("%v\n", computeDistance(n))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
}

// computeDistance is my solution to Part One and solves the shortest path from
// any position in a matrix to position 1, the center of the matrix. The
// position identifiers are assigned from 1 to infinity, starting in the
// center-most position, spiraling out counter-clockwise starting from the
// immediate right of the center.
//
// Example:
//
// ╔═══╦═══╦═══╗
// ║ 5 ║ 4 ║ 3 ║
// ╠═══╬═══╬═══╣
// ║ 6 ║ 1 ║ 2 ║
// ╠═══╬═══╬═══╣
// ║ 7 ║ 8 ║ 9 ║
// ╚═══╩═══╩═══╝
//
// This function returns in O(1) constant time.
func computeDistance(n int) int {
	if n < 1 {
		panic("invalid spiral memory address")
	}
	if n == 1 {
		return 0
	}
	w := int(math.Ceil(math.Sqrt(float64(n))))
	if w%2 == 0 {
		w++
	}
	r := (w - 1) / 2
	l := (w - 2) * (w - 2)
	x := ((n - l - 1) % (r * 2)) - r + 1
	if x < 0 {
		x = -x
	}
	return r + x
}

type Matrix struct {
	width int
	data  []int
}

// findNextInMatrix is my solution to Part Two. It's too complicated to explain
// right now... Essentially it builds the spiral matric with dimensions w * w
// until is encounters a cell whose value exceeds n.
func findNextInMatrix(n, w int) int {
	m := &Matrix{
		width: w,
		data:  make([]int, w*w),
	}
	x := w / 2
	y := w / 2
	m.data[y*w+x] = 1
	for r := 1; r <= (w / 2); r++ {
		x++
		y++
		for u := 0; u < r*2; u++ {
			y--
			if v := m.set(x, y); v > n {
				return v
			}
		}
		for l := 0; l < r*2; l++ {
			x--
			if v := m.set(x, y); v > n {
				return v
			}
		}
		for d := 0; d < r*2; d++ {
			y++
			if v := m.set(x, y); v > n {
				return v
			}
		}
		for rt := 0; rt < r*2; rt++ {
			x++
			if v := m.set(x, y); v > n {
				return v
			}
		}
	}
	return -1
}

// set computes and stores the value of the given Matrix cell by summing the
// value of all surrounding cells in their current state.
func (m Matrix) set(x, y int) int {
	v := 0
	if x > 0 {
		v += m.data[y*m.width+x-1]
	}
	if x < m.width-1 {
		v += m.data[y*m.width+x+1]
	}
	if y > 0 {
		v += m.data[(y-1)*m.width+x]
		if x != 0 {
			v += m.data[(y-1)*m.width+x-1]
		}
		if x < m.width-1 {
			v += m.data[(y-1)*m.width+x+1]
		}
	}
	if y < m.width-1 {
		v += m.data[(y+1)*m.width+x]
		if x != 0 {
			v += m.data[(y+1)*m.width+x-1]
		}
		if x < m.width-1 {
			v += m.data[(y+1)*m.width+x+1]
		}
	}
	m.data[y*m.width+x] = v
	return v
}

// Print prints the Matrix to the given writer for debug purposes.
func (m *Matrix) Print(w io.Writer) {
	for j := 0; j < m.width; j++ {
		switch j {
		case 0:
			fmt.Fprintf(w, "╔═════════")
		case m.width - 1:
			fmt.Fprintf(w, "╦═════════╗\n")
		default:
			fmt.Fprintf(w, "╦═════════")
		}
	}
	for i := 0; i < len(m.data); i++ {
		if i%m.width == 0 && i > 0 {
			for j := 0; j < m.width; j++ {
				switch j {
				case 0:
					fmt.Fprintf(w, "║\n╠═════════")
				case m.width - 1:
					fmt.Fprintf(w, "╬═════════╣\n")
				default:
					fmt.Fprintf(w, "╬═════════")
				}
			}
		}
		fmt.Fprintf(w, "║ %7d ", m.data[i])
	}
	for j := 0; j < m.width; j++ {
		switch j {
		case 0:
			fmt.Fprintf(w, "║\n╚═════════")
		case m.width - 1:
			fmt.Fprintf(w, "╩═════════╝\n")
		default:
			fmt.Fprintf(w, "╩═════════")
		}
	}
}

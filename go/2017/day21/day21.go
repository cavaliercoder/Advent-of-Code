package day21

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Block []byte

func ParseBlock(s string) Block {
	b := make(Block, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			b = append(b, 0x00)

		case '#':
			b = append(b, 0x01)
		}
	}
	return b
}

// Transforms returns mutations of the given block.
func (b Block) Transforms() []Block {
	var transforms [][]int
	switch len(b) {
	case 4:
		transforms = [][]int{
			[]int{1, 0, 3, 2}, // flip horizontal
			[]int{2, 3, 0, 1}, // flip vertical
			[]int{2, 0, 3, 1}, // rotate 90°
			[]int{3, 2, 1, 0}, // rotate 180°
			[]int{1, 3, 0, 2}, // rotate 270°
		}
	case 9:
		transforms = [][]int{
			[]int{2, 1, 0, 5, 4, 3, 8, 7, 6}, // flip horizontal
			[]int{6, 7, 8, 3, 4, 5, 0, 1, 2}, // flip vertical
			[]int{6, 3, 0, 7, 4, 1, 8, 5, 2}, // rotate 90°
			[]int{8, 7, 6, 5, 4, 3, 2, 1, 0}, // rotate 180°
			[]int{2, 5, 8, 1, 4, 7, 0, 3, 6}, // rotate 270°

			// other are needed - the following suffices for my inputs
			[]int{8, 5, 2, 7, 4, 1, 6, 3, 0}, // flip horizontal and rotate 90°
		}
	default:
		panic(fmt.Sprintf("bad block size: %d", len(b)))
	}

	out := []Block{b}
	for t := 0; t < len(transforms); t++ {
		v := make(Block, len(b))
		for i := 0; i < len(b); i++ {
			v[i] = b[transforms[t][i]]
		}
		out = append(out, v)
	}
	return out
}

func (b Block) String() string {
	s := bytes.Buffer{}
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case 0:
			s.WriteByte('.')
		default:
			s.WriteByte('#')
		}
	}
	return s.String()
}

type EnhancementRule struct {
	Input, Output Block
}

func ParseEnhancementRule(s string) EnhancementRule {
	tkns := strings.Split(s, " => ")
	v := EnhancementRule{
		Input:  ParseBlock(tkns[0]),
		Output: ParseBlock(tkns[1]),
	}
	return v
}

// EnhancementMap maps rule inputs (hashed by string) to output Blocks.
type EnhancementRuleList map[string]Block

// Add stores an enhancement rule
func (l EnhancementRuleList) Add(e EnhancementRule) {
	et := e.Input.Transforms()
	for i := 0; i < len(et); i++ {
		l[et[i].String()] = e.Output
	}
}

// Get returns the output block for the given input block.
func (l EnhancementRuleList) Get(b Block) Block {
	if out, ok := l[b.String()]; ok {
		return out
	}
	panic(fmt.Sprintf("no rule found for: %v", b))
}

type Matrix struct {
	Size, BlockSize int
	Blocks          []Block
}

func NewMatrix(size, blocksize int) *Matrix {
	if size%blocksize != 0 {
		panic("invalid size/blocksize")
	}
	blockcount := (size * size) / (blocksize * blocksize)
	m := &Matrix{
		Size:      size,
		Blocks:    make([]Block, blockcount),
		BlockSize: blocksize,
	}
	for i := 0; i < blockcount; i++ {
		m.Blocks[i] = make(Block, blocksize*blocksize)
	}
	return m
}

// InitialMatrix returns a matrix initialized according to the challenge input:
//
// .#.
// ..#
// ###
//
func InitialMatrix() *Matrix {
	m := NewMatrix(3, 3)
	copy(m.Blocks[0], []byte{0, 1, 0, 0, 0, 1, 1, 1, 1})
	return m
}

// CountOnPixels returns the number of pixels in all blocks in a matrix that are
// in the "on" ('#') state.
func (m *Matrix) CountOnPixels() int {
	count := 0
	for i := 0; i < len(m.Blocks); i++ {
		for j := 0; j < len(m.Blocks[i]); j++ {
			if m.Blocks[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

// Realign returns a new Matrix with all data redistributed between all blocks
// such that they are aligned to the given block size n.
func (m *Matrix) Realign(size, blocksize int) *Matrix {
	mm := NewMatrix(size, blocksize)
	for i := 0; i < m.Size*m.Size; i++ {
		/*
			block row = i / mw*bw
			block col = i/bw % mw/bw
			block idx = (i/(mw*bw))*(mw/bw)) + ((i/bw)%(mw/bw))

			internal row = (i / mw) % bw
			internal col = i % bw
			internal idx = (((i / mw) % bw) * bw) + (i % bw)
		*/

		// TODO: reduce math
		bi0 := ((i / (m.Size * m.BlockSize)) * (m.Size / m.BlockSize)) + ((i / m.BlockSize) % (m.Size / m.BlockSize))
		ii0 := (((i / m.Size) % m.BlockSize) * m.BlockSize) + (i % m.BlockSize)
		bi1 := ((i / (mm.Size * mm.BlockSize)) * (mm.Size / mm.BlockSize)) + ((i / mm.BlockSize) % (mm.Size / mm.BlockSize))
		ii1 := (((i / mm.Size) % mm.BlockSize) * mm.BlockSize) + (i % mm.BlockSize)
		mm.Blocks[bi1][ii1] = m.Blocks[bi0][ii0]
	}
	return mm
}

// Generate is my solution for Part One and Two. It steps through n generations
// of the martix, using the given ruleset to mutate each generation.
func Generate(n int, ruleset EnhancementRuleList) int {
	m := InitialMatrix()
	for i := 0; i < n; i++ {
		for j := 0; j < len(m.Blocks); j++ {
			m.Blocks[j] = ruleset.Get(m.Blocks[j])
		}
		m.BlockSize = int(math.Sqrt(float64(len(m.Blocks[0]))))
		m.Size = m.BlockSize * int(math.Sqrt(float64(len(m.Blocks))))
		switch {
		case m.Size%2 == 0:
			m = m.Realign(m.Size, 2)
		case m.Size%3 == 0:
			m = m.Realign(m.Size, 3)
		default:
			panic(fmt.Sprintf("bad matrix size: %d", m.Size))
		}
	}
	return m.CountOnPixels()
}

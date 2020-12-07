package day09

import (
	"fmt"
)

// A Parser reads a stream of characters and extracts the number of garbage
// characters and the "group score".
type Parser struct {
	in             []byte
	n              int
	curr, next     byte
	score, garbage int
}

func NewParser(b []byte) *Parser {
	return &Parser{in: b}
}

// readByte returns the next byte in the parser's input or zero once all of the
// input has been read.
func (p *Parser) readByte() byte {
	p.curr = p.next
	if p.n < len(p.in) {
		p.next = p.in[p.n]
		p.n++
	} else {
		p.next = 0
	}
	return p.curr
}

// readUnignoredByte returns the next byte in the parser's input, ignoring any !
// characters and their subsequent byte.
func (p *Parser) readUnignoredByte() byte {
	p.readByte()
	if p.curr == '!' {
		p.readByte()
		return p.readUnignoredByte()
	}
	return p.curr
}

// readNonGarbageByte returns the next byte in the parser's input that is not
// part of a garbage sequence (including the surrounding < >) or ignored by a !.
func (p *Parser) readNonGarbageByte() byte {
	p.readUnignoredByte()
	if p.curr == '<' {
		for {
			switch p.readUnignoredByte() {
			case '>':
				return p.readNonGarbageByte()
			case 0:
				return 0
			default:
				p.garbage++
			}
		}
	}
	return p.curr
}

// Parse reads all the parser's input, keeping track of the number of garbage
// characters that have been discarded, and the input's "group score".
func (p *Parser) Parse() {
	p.readByte() // seed p.next

	var d int
	for {
		p.readNonGarbageByte()
		switch p.curr {
		case '{':
			d++
			p.score += d
		case '}':
			d--
		case ',':
			// do nothing
		case 0:
		case '\n':
			return
		default:
			panic(fmt.Sprintf("invalid character: %q at %d", p.curr, p.n))
		}
	}
}

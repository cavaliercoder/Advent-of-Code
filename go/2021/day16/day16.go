package day16

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type Type int

const (
	TypeSum         Type = 0
	TypeProduct     Type = 1
	TypeMin         Type = 2
	TypeMax         Type = 3
	TypeLiteral     Type = 4
	TypeGreaterThan Type = 5
	TypeLessThan    Type = 6
	TypeEqual       Type = 7
)

// Expr is an expression that may be a literal or an operator and its operands.
type Expr struct {
	Offset   int
	Version  int
	Type     Type
	Len      int
	Size     int
	Operands []*Expr
	V        int
}

// Eval the expression.
func (c *Expr) Eval() int {
	switch c.Type {
	case TypeSum:
		n := 0
		for _, expr := range c.Operands {
			n += expr.Eval()
		}
		return n

	case TypeProduct:
		n := 1
		for _, expr := range c.Operands {
			n *= expr.Eval()
		}
		return n

	case TypeMin:
		min := c.Operands[0].Eval()
		for _, expr := range c.Operands[1:] {
			if n := expr.Eval(); n < min {
				min = n
			}
		}
		return min

	case TypeMax:
		max := c.Operands[0].Eval()
		for _, expr := range c.Operands[1:] {
			if n := expr.Eval(); n > max {
				max = n
			}
		}
		return max

	case TypeLiteral:
		return c.V

	case TypeGreaterThan:
		a, b := c.Operands[0], c.Operands[1]
		if a.Eval() > b.Eval() {
			return 1
		}
		return 0

	case TypeLessThan:
		a, b := c.Operands[0], c.Operands[1]
		if a.Eval() < b.Eval() {
			return 1
		}
		return 0

	case TypeEqual:
		a, b := c.Operands[0], c.Operands[1]
		if a.Eval() == b.Eval() {
			return 1
		}
		return 0

	}
	panic("unsupported operator")
}

func (c *Expr) String() string {
	b := new(bytes.Buffer)
	formatExpr(b, c)
	return b.String()
}

func formatExpr(w io.Writer, expr *Expr) {
	switch expr.Type {
	case TypeSum:
		fmt.Fprintf(w, "sum(")
	case TypeProduct:
		fmt.Fprintf(w, "mult(")
	case TypeMin:
		fmt.Fprintf(w, "min(")
	case TypeMax:
		fmt.Fprintf(w, "max(")
	case TypeLiteral:
		fmt.Fprintf(w, "%d", expr.Eval())
		return
	case TypeGreaterThan:
		fmt.Fprintf(w, "gt(")
	case TypeLessThan:
		fmt.Fprintf(w, "lt(")
	case TypeEqual:
		fmt.Fprintf(w, "eq(")
	}
	for i, op := range expr.Operands {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		formatExpr(w, op)
	}
	fmt.Fprintf(w, ")")
}

type lexer struct {
	Data   []byte
	ptr    int
	offset int
}

func Lex(b []byte) ([]*Expr, error) {
	c := &lexer{}
	c.ptr = 0
	c.Data = make([]byte, len(b)/2)
	_, err := hex.Decode(c.Data, b)
	if err != nil {
		return nil, err
	}
	return c.Tokenize(), nil
}

// Read a single bit.
func (c *lexer) Bit() (b byte, ok bool) {
	// Parsing one bit at a time sucks but is semi-necessary given that packets
	// are not aligned to any notable boundary.
	if c.ptr >= len(c.Data) {
		return
	}
	ok = true
	b = (c.Data[c.ptr] >> byte(7-c.offset)) & 0x01
	c.offset = (c.offset + 1) % 8
	if c.offset == 0 {
		c.ptr++
	}
	return
}

// Read a byte with sz bits.
func (c *lexer) Byte(sz int) (n byte, ok bool) {
	if sz < 0 || sz > 8 {
		panic("invalid size")
	}
	for i := 0; i < sz; i++ {
		n <<= 1
		b, ok := c.Bit()
		if !ok {
			return 0, false
		}
		n |= b
	}
	ok = true
	return
}

// Read a uint16 with sz bits.
func (c *lexer) Uint16(sz int) (n uint16, ok bool) {
	if sz < 0 || sz > 16 {
		panic("invalid size")
	}
	for i := 0; i < sz; i++ {
		n <<= 1
		b, ok := c.Bit()
		if !ok {
			return 0, false
		}
		n |= uint16(b)
	}
	ok = true
	return
}

// Read an expression from a single packet (not including operands).
func (c *lexer) Expr() *Expr {
	offset := 8*c.ptr + c.offset
	b, ok := c.Byte(6)
	if !ok {
		return nil
	}
	expr := &Expr{
		Offset:  offset,
		Version: int((b >> 3) & 0x07),
		Type:    Type(b & 0x07),
	}

	// literal expr
	if expr.Type == TypeLiteral {
		for {
			b, ok = c.Byte(5)
			if !ok {
				return nil
			}
			expr.V <<= 4
			expr.V |= int(b & 0x0F)
			if b&0x10 == 0 {
				break
			}
		}
		return expr
	}

	// operator expr
	b, ok = c.Bit()
	if !ok {
		return nil
	}
	switch b {
	case 0: // 15-bit operands size
		n, ok := c.Uint16(15)
		if !ok {
			return nil
		}
		expr.Size = int(n)
	case 1: // 11-bit operands count
		n, ok := c.Uint16(11)
		if !ok {
			return nil
		}
		expr.Len = int(n)
	}
	return expr
}

// Read all expressions and return them in lexical order.
func (c *lexer) Tokenize() []*Expr {
	a := make([]*Expr, 0, 64)
	for {
		expr := c.Expr()
		if expr == nil {
			break
		}
		a = append(a, expr)
	}
	return a
}

// Parse all expressions into an AST.
func Parse(b []byte) (*Expr, error) {
	tokens, err := Lex(b)
	if err != nil {
		return nil, err
	}
	p := &parser{
		tokens: tokens,
	}
	return p.parse(), nil
}

type parser struct {
	tokens []*Expr
	ptr    int
}

func (c *parser) parse() *Expr {
	expr := c.tokens[c.ptr]
	c.ptr++
	if expr.Type == TypeLiteral {
		return expr
	}
	if sz := expr.Size; sz > 0 {
		start := c.tokens[c.ptr].Offset
		for {
			if c.ptr >= len(c.tokens) {
				break
			}
			if c.tokens[c.ptr].Offset >= start+expr.Size {
				break
			}
			expr.Operands = append(expr.Operands, c.parse())
		}
	} else {
		for i := 0; i < expr.Len; i++ {
			expr.Operands = append(expr.Operands, c.parse())
		}
	}
	return expr
}

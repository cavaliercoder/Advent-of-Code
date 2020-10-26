package common

import (
	"strings"
	"testing"
)

func TestGrid(t *testing.T) {
	g, err := ReadGrid(strings.NewReader("123\n456\n789\n"))
	if err != nil {
		t.Error(err)
	}
	AssertInt(t, 3, g.Width, "bad grid width")
	AssertInt(t, 3, g.Height, "bad grid height")
	AssertBytes(t, []byte("123456789"), g.Data, "bad grid data")
	i := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			pos := Pos{x, y}
			AssertByte(t, '1'+byte(i), g.Get(pos), "bad grid get")
			AssertPos(t, pos, g.Pos(i), "bad position lookup")
			AssertInt(t, i, g.Index(pos), "bad index lookup")
			i++
		}
	}
}

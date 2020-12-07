package aoc

import (
	"bytes"
	"fmt"
	"testing"
)

func AssertString(t *testing.T, expect, actual string, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%s', got: '%s'", s, expect, actual)
	return false
}

func AssertInt(t *testing.T, expect, actual int, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

func AssertByte(t *testing.T, expect, actual byte, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%0X', got: '%0X'", s, expect, actual)
	return false
}

func AssertBytes(t *testing.T, expect, actual []byte, format string, a ...interface{}) bool {
	if bytes.Equal(expect, actual) {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

func AssertPos(t *testing.T, expect, actual Pos, format string, a ...interface{}) bool {
	if expect == actual {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%v', got: '%v'", s, expect, actual)
	return false
}

func AssertGridCell(t *testing.T, grid *Grid, p Pos, expect byte, format string, a ...interface{}) bool {
	s := fmt.Sprintf(format, a...)
	if !grid.Contains(p) {
		t.Errorf("%s. Expected '%c', got: out of range at %v", s, expect, p)
		return false
	}
	actual := grid.Get(p)
	if actual == expect {
		return true
	}
	t.Errorf("%s. Expected '%c', got '%c' at %v", s, expect, actual, p)
	return false
}

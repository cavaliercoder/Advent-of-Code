package aoc2021

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

func AssertInt64(t *testing.T, expect, actual int64, format string, a ...interface{}) bool {
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

package assert

import (
	"bytes"
	"fmt"
	"testing"
)

func String(t *testing.T, expect, actual string, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%s', got: '%s'", s, expect, actual)
	return false
}

func Int(t *testing.T, expect, actual int, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

func Int64(t *testing.T, expect, actual int64, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

func Bool(t *testing.T, expect, actual bool, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: %v, got: %v", s, expect, actual)
	return false
}

func Byte(t *testing.T, expect, actual byte, format string, a ...interface{}) bool {
	if actual == expect {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%0X', got: '%0X'", s, expect, actual)
	return false
}

func Bytes(t *testing.T, expect, actual []byte, format string, a ...interface{}) bool {
	if bytes.Equal(expect, actual) {
		return true
	}
	s := fmt.Sprintf(format, a...)
	t.Errorf("%s. Expected: '%d', got: '%d'", s, expect, actual)
	return false
}

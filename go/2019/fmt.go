package aoc

import (
	"fmt"
	"os"
)

var (
	enableDebug = os.Getenv("DEBUG") != ""
)

func Debugf(format string, a ...interface{}) {
	if !enableDebug {
		return
	}
	fmt.Fprintf(os.Stderr, format, a...)
}

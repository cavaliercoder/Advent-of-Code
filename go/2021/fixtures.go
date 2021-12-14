package aoc2021

import (
	"os"
	"path/filepath"
)

func Fixture(name string) string {
	return filepath.Join("..", "..", "..", "inputs", "2021", name+".txt")
}

func OpenFixture(name string) (*os.File, error) {
	return os.Open(Fixture(name))
}

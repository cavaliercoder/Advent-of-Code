package aoc

import (
	"os"
	"path/filepath"
)

func Fixture(name string) string {
	return filepath.Join("..", "..", "..", "inputs", "2019", name+".dat")
}

func OpenFixture(name string) (*os.File, error) {
	return os.Open(Fixture(name))
}

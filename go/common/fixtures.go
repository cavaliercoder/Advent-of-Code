package common

import (
	"os"
	"path/filepath"
)

func Fixture(name string) string {
	return filepath.Join("..", "..", "inputs", name+".dat")
}

func OpenFixture(name string) (*os.File, error) {
	return os.Open(Fixture(name))
}

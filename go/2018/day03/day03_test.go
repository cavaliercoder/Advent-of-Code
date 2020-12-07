package day03

import (
	"bufio"
	"fmt"

	. "aoc"
)

func ParseClaims(name string) []*Claim {
	f := MustOpenFixture("day03")
	defer f.Close()

	claims := make([]*Claim, 0, 64)
	s := bufio.NewScanner(f)
	for s.Scan() {
		c := &Claim{}
		n, err := fmt.Sscanf(
			s.Text(),
			"#%d @ %d,%d: %dx%d",
			&c.ID,
			&c.X,
			&c.Y,
			&c.W,
			&c.H,
		)
		if err != nil {
			panic(err)
		}
		if n != 5 {
			panic(s.Text())
		}
		claims = append(claims, c)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return claims
}

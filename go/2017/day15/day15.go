package day15

type Generator interface {
	Next() uint64
}

type generator struct {
	factor, value uint64
}

func (g *generator) Next() uint64 {
	g.value = (g.value * g.factor) % 2147483647
	return g.value
}

type multipleGenerator struct {
	factor, multiplier, value uint64
}

func (g *multipleGenerator) Next() uint64 {
	v := g.value
	for {
		v = (v * g.factor) % 2147483647
		if v%g.multiplier == 0 {
			g.value = v
			return v
		}
	}
}

func CountMatches(n int, a, b Generator) int {
	var count int
	for i := 0; i < n; i++ {
		if a.Next()&0xFFFF == b.Next()&0xFFFF {
			count++
		}
	}
	return count
}

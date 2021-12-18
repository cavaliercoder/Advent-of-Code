package day14

import "fmt"

type Pair [2]byte

type Template struct {
	countByPair    map[Pair]int
	countByPolymer map[byte]int
}

func NewTemplate(elem ...byte) *Template {
	tmpl := &Template{
		countByPair:    make(map[Pair]int, 256),
		countByPolymer: make(map[byte]int, 256),
	}
	if len(elem) == 0 {
		return tmpl
	}
	for i := 0; i < len(elem); i++ {
		if i < len(elem)-1 {
			k := Pair{elem[i], elem[i+1]}
			tmpl.countByPair[k]++
		}
		tmpl.countByPolymer[elem[i]]++
	}
	return tmpl
}

func (c *Template) Step(rules map[Pair]byte) {
	newPairs := make(map[Pair]int, len(c.countByPair))
	for k, n := range c.countByPair {
		b, ok := rules[k]
		if !ok {
			panic(fmt.Sprintf("bad pair: %s", k))
		}
		c.countByPolymer[b] += n
		c.countByPair[k] -= n
		newPairs[Pair{k[0], b}] += n
		newPairs[Pair{b, k[1]}] += n
	}
	for k, v := range newPairs {
		c.countByPair[k] += v
	}
}

func (c *Template) Hash() int {
	min, max := -1, -1
	for _, v := range c.countByPolymer {
		if min < 0 || v < min {
			min = v
		}
		if max < 0 || v > max {
			max = v
		}
	}
	return max - min
}

package day22

import "aoc/internal/geo3d"

type Op struct {
	On   bool
	Cube geo3d.Cube
}

type Reactor struct {
	m map[geo3d.Cube]struct{}
}

func NewReactor() *Reactor { return &Reactor{} }

func (r *Reactor) do(op Op) {
	newCubes := make(map[geo3d.Cube]struct{})
	for c := range r.m {
		for _, child := range c.Split(op.Cube) {
			if child.In(op.Cube) {
				continue
			}
			newCubes[child] = struct{}{}
		}
	}
	r.m = newCubes
	if op.On {
		r.m[op.Cube] = struct{}{}
	}
}

var initBounds = geo3d.Cube{
	A: geo3d.Pos{X: -50, Y: -50, Z: -50},
	B: geo3d.Pos{X: 51, Y: 51, Z: 51},
}

func (r *Reactor) Init(ops ...Op) int {
	for _, op := range ops {
		if op.Cube.In(initBounds) {
			r.do(op)
		}
	}
	n := 0
	for cube := range r.m {
		n += cube.Volume()
	}
	return n
}

func (r *Reactor) Reboot(ops ...Op) int {
	for _, op := range ops {
		r.do(op)
	}
	n := 0
	for cube := range r.m {
		n += cube.Volume()
	}
	return n
}

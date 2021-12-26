package day19

import (
	"fmt"

	"aoc/internal/geo3d"
)

type Scanner struct {
	Position geo3d.Pos
	Beacons  map[geo3d.Pos]struct{}

	rotations []*Scanner
	moves     map[geo3d.Pos]*Scanner
	noFit     map[geo3d.Pos]struct{}
}

func NewScanner(position geo3d.Pos, beacons ...geo3d.Pos) *Scanner {
	v := &Scanner{
		Position: position,
		Beacons:  make(map[geo3d.Pos]struct{}, 32),
		moves:    make(map[geo3d.Pos]*Scanner),
		noFit:    make(map[geo3d.Pos]struct{}),
	}
	v.Add(beacons...)
	return v
}

func (c *Scanner) Add(beacons ...geo3d.Pos) {
	for _, obj := range beacons {
		c.Beacons[obj] = struct{}{}
	}
}

func (c *Scanner) IsInRange(p geo3d.Pos) bool {
	if p.X < c.Position.X-1000 ||
		p.Y < c.Position.Y-1000 ||
		p.Z < c.Position.Z-1000 {
		return false
	}
	if p.X > c.Position.X+1000 ||
		p.Y > c.Position.Y+1000 ||
		p.Z > c.Position.Z+1000 {
		return false
	}
	return true
}

func (c *Scanner) HasBeacon(p geo3d.Pos) bool {
	_, ok := c.Beacons[p]
	return ok
}

func (c *Scanner) Transform(t geo3d.Transform) *Scanner {
	v := NewScanner(c.Position.Transform(t))
	for c := range c.Beacons {
		v.Add(c.Transform(t))
	}
	return v
}

func (c *Scanner) Rotations() []*Scanner {
	if c.rotations != nil {
		return c.rotations
	}
	c.rotations = make([]*Scanner, len(geo3d.Rotations))
	for i, t := range geo3d.Rotations {
		c.rotations[i] = c.Transform(t)
	}
	return c.rotations
}

func (c *Scanner) Move(offset geo3d.Pos) *Scanner {
	if v, ok := c.moves[offset]; ok {
		return v
	}
	v := NewScanner(c.Position.Add(offset))
	for obj := range c.Beacons {
		v.Add(obj.Add(offset))
	}
	c.moves[offset] = v
	return v
}

type Field struct {
	Scanners []*Scanner
	Beacons  map[geo3d.Pos]struct{}
}

func NewField(scanners ...*Scanner) *Field {
	v := &Field{
		Scanners: make([]*Scanner, 0, 64),
		Beacons:  make(map[geo3d.Pos]struct{}, 4096),
	}
	v.Add(scanners...)
	return v
}

func (c *Field) Add(scanners ...*Scanner) {
	c.Scanners = append(c.Scanners, scanners...)
	for _, s := range scanners {
		for obj := range s.Beacons {
			c.Beacons[obj] = struct{}{}
		}
	}
}

func (c *Field) HasBeacon(beacon geo3d.Pos) bool {
	_, ok := c.Beacons[beacon]
	return ok
}

func (c *Field) CanFit(s *Scanner) bool {
	n := 0
	for beacon := range s.Beacons {
		if c.HasBeacon(beacon) {
			n++
			if n == 12 {
				break
			}
		}
	}
	if n < 12 {
		return false
	}
	for beacon := range c.Beacons {
		if !s.IsInRange(beacon) {
			continue
		}
		if !s.HasBeacon(beacon) {
			return false
		}
	}
	return true
}

func (c *Field) MaxManhattan() int {
	max := 0
	for i := 0; i < len(c.Scanners)-1; i++ {
		for j := 1; j < len(c.Scanners); j++ {
			n := c.Scanners[i].Position.Sub(c.Scanners[j].Position).Manhattan()
			if n > max {
				max = n
			}
		}
	}
	return max
}

func Merge(scanners ...*Scanner) (field *Field, ok bool) {
	// TODO: I'm very slow...
	field = NewField()
	if len(scanners) == 0 {
		ok = true
		return
	}
	field.Add(scanners[0])
	scanners[0] = nil
	merges, needMerges := 1, len(scanners)
	for merges < needMerges {
		merged := false
		for i, scanner := range scanners {
			if scanner == nil {
				continue
			}
			r := tryRotations(field, scanner)
			if r == nil {
				continue
			}
			field.Add(r)
			scanners[i] = nil
			merged = true
			merges++
			fmt.Printf("merged %d/%d...\n", merges, needMerges)
		}
		if !merged {
			break
		}
	}
	ok = merges == needMerges
	return
}

func tryRotations(field *Field, s *Scanner) *Scanner {
	rotations := s.Rotations()
	for b := range field.Beacons {
		if _, ok := s.noFit[b]; ok {
			continue
		}
		for _, r := range rotations {
			for a := range r.Beacons {
				offset := b.Sub(a)
				m := r.Move(offset)
				if field.CanFit(m) {
					return m
				}
			}
		}
		s.noFit[b] = struct{}{}
	}
	return nil
}

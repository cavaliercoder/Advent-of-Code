package day19

import (
	"math/rand"

	"aoc/internal/geo3d"
)

type Scanner struct {
	Position geo3d.Pos
	Beacons  []geo3d.Pos

	rotations []*Scanner
}

func NewScanner(position geo3d.Pos, beacons ...geo3d.Pos) *Scanner {
	v := &Scanner{
		Position: position,
		Beacons:  make([]geo3d.Pos, 0, 32),
	}
	v.Add(beacons...)
	return v
}

func (c *Scanner) Add(beacons ...geo3d.Pos) {
	c.Beacons = append(c.Beacons, beacons...)
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
	for _, b := range c.Beacons {
		if b == p {
			return true
		}
	}
	return false
}

func (c *Scanner) Transform(t geo3d.Transform) *Scanner {
	v := NewScanner(c.Position.Transform(t))
	for _, obj := range c.Beacons {
		v.Add(obj.Transform(t))
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
	v := NewScanner(c.Position.Add(offset))
	for _, obj := range c.Beacons {
		v.Add(obj.Add(offset))
	}
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
		for _, obj := range s.Beacons {
			c.Beacons[obj] = struct{}{}
		}
	}
}

func (c *Field) HasBeacon(beacon geo3d.Pos) bool {
	_, ok := c.Beacons[beacon]
	return ok
}

// CanFit returns true if s has at least 12 beacons shared with the field and no
// conflicting beacons.
func (c *Field) CanFit(s *Scanner) bool {
	n := 0
	for _, beacon := range s.Beacons {
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
	field = NewField()
	if len(scanners) == 0 {
		ok = true
		return
	}
	field.Add(scanners[0])
	scanners[0] = nil
	merges, needMerges := 1, len(scanners)
	for merges < needMerges {
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
			merges++
		}
	}
	ok = merges == needMerges
	return
}

func tryRotations(field *Field, s *Scanner) *Scanner {
	rotations := s.Rotations()
	for b := range field.Beacons {
		for _, r := range rotations {
			beacon := r.Beacons[rand.Int()%len(r.Beacons)]
			offset := b.Sub(beacon)
			m := r.Move(offset)
			if field.CanFit(m) {
				return m
			}
		}
	}
	return nil
}

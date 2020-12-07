package day03

import "sort"

type Pos struct{ X, Y int }

type Claim struct {
	ID, X, Y, W, H int

	NextX, NextY *Claim
}

type ClaimsByX []*Claim

func (c ClaimsByX) Len() int           { return len(c) }
func (c ClaimsByX) Less(i, j int) bool { return c[i].X < c[j].X }
func (c ClaimsByX) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func NewClaimsByX(claims []*Claim) ClaimsByX {
	v := make(ClaimsByX, len(claims))
	copy(v, claims)
	sort.Sort(v)
	for i := 0; i < len(v); i++ {
		if i < len(v)-1 {
			v[i].NextX = v[i+1]
		}
	}
	return v
}

type ClaimsByY []*Claim

func (c ClaimsByY) Len() int           { return len(c) }
func (c ClaimsByY) Less(i, j int) bool { return c[i].Y < c[j].Y }
func (c ClaimsByY) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func NewClaimsByY(claims []*Claim) ClaimsByY {
	v := make(ClaimsByY, len(claims))
	copy(v, claims)
	sort.Sort(v)
	for i := 0; i < len(v); i++ {
		if i < len(v)-1 {
			v[i].NextY = v[i+1]
		}
	}
	return v
}

func Intersect(a, b *Claim) *Claim {
	c := &Claim{}
	c.X = b.X
	c.W = a.X + a.W - b.X + 1
	return c
}

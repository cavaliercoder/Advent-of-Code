package geo

type Rect struct {
	A, B Pos
}

func (r Rect) ContainsX(x int) bool { return r.A.X <= x && x <= r.B.X }
func (r Rect) ContainsY(y int) bool { return r.A.Y >= y && y >= r.B.Y }
func (r Rect) Contains(p Pos) bool  { return r.ContainsX(p.X) && r.ContainsY(p.Y) }

func (r Rect) Expand(n int) Rect {
	return Rect{
		A: Pos{X: r.A.X - n, Y: r.A.Y + n},
		B: Pos{X: r.B.X + n, Y: r.B.Y - n},
	}
}

func (r Rect) Fit(p Pos) Rect {
	if p.X < r.A.X {
		r.A.X = p.X
	}
	if p.X > r.B.X {
		r.B.X = p.X
	}
	if p.Y > r.A.Y {
		r.A.Y = p.Y
	}
	if p.Y < r.B.Y {
		r.B.Y = p.Y
	}
	return r
}

func (r Rect) Iter() *RectIter {
	return &RectIter{r: r}
}

type RectIter struct {
	r Rect
	p Pos
	i int
}

func (iter *RectIter) Next() bool {
	if iter.i == 0 {
		iter.p = iter.r.A
	} else {
		iter.p.X++
		if iter.p.X > iter.r.B.X {
			iter.p.X = iter.r.A.X
			iter.p.Y--
		}
	}
	iter.i++
	return iter.p.Y >= iter.r.B.Y
}

func (iter *RectIter) Pos() Pos { return iter.p }

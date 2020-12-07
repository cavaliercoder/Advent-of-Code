package day17

type SpinLock struct {
	Value int
	Next  *SpinLock
}

func NewSpinLock() *SpinLock {
	n := &SpinLock{Value: 0}
	n.Next = n
	return n
}

// Insert inserts n values into the spinlock ring buffer, stepping forward by
// the given step count each time. The return value is the value following the
// last inserted value.
//
// This function runs in linear O(n) time where n is the number of insertions
// to be made.
func (s *SpinLock) Insert(n, step int) int {
	v := s
	for i := 1; i <= n; i++ {
		for j := 0; j < step; j++ {
			v = v.Next
		}
		ins := &SpinLock{
			Value: i,
			Next:  v.Next,
		}
		v.Next = ins
		v = ins
	}
	return v.Next.Value
}

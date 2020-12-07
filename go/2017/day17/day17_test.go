package day17

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		n, step, expect int
	}{
		{2017, 3, 638},
		{2017, 363, 136},
	}
	for _, test := range tests {
		s := NewSpinLock()
		actual := s.Insert(test.n, test.step)
		if actual != test.expect {
			t.Errorf("expected %d, got %d", test.expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		n, step, expect int
	}{
		{5e7, 363, 1080289},
	}
	for _, test := range tests {
		s := NewSpinLock()
		s.Insert(test.n, test.step)
		actual := s.Next.Value
		if actual != test.expect {
			t.Errorf("expected %d, got %d", test.expect, actual)
		}
	}
}

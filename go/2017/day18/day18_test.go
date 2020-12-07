package day18

import "testing"

const (
	input1 = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

	input2 = `set i 31
set a 1
mul p 17
jgz p p
mul a 2
add i -1
jgz i -2
add a -1
set i 127
set p 680
mul p 8505
mod p a
mul p 129749
add p 12345
mod p a
set b p
mod b 10000
snd b
add i -1
jgz i -9
jgz a 3
rcv b
jgz b -1
set f 0
set i 126
rcv a
rcv b
set p a
mul p -1
add p b
jgz p 4
snd a
set a b
jgz 1 3
snd b
set f 1
add i -1
jgz i -11
snd a
jgz f -16
jgz a -19`
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		Input  string
		Expect int
	}{
		{
			Input:  input1,
			Expect: 4,
		},
		{
			Input:  input2,
			Expect: 3188,
		},
	}

	for _, test := range tests {
		actual := LastSnd(test.Input)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  string
		Expect int
	}{
		{
			Input:  input1,
			Expect: 1,
		},
		{
			Input:  input2,
			Expect: 7112,
		},
	}

	for _, test := range tests {
		actual := Duet(test.Input)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

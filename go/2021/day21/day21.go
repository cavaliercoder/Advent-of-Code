package day21

// Die is a deterministic 100-sided die.
type Die struct {
	v, n int
}

// N returns the number of times the die has been rolled.
func (d *Die) N() int { return d.n }

// RollN rolls the die n times and returns the sum of the results.
func (d *Die) RollN(n int) (v int) {
	for i := 0; i < n; i++ {
		d.n++
		v += d.v + 1
		d.v = (d.v + 1) % 100
	}
	return
}

type Player struct {
	Position int // 0-indexed
	Score    int
}

func (p Player) Move(roll int) Player {
	p.Position = (p.Position + roll) % 10
	p.Score += p.Position + 1
	return p
}

func Play(pos1, pos2 int) int {
	d := &Die{}
	p1, p2 := Player{Position: pos1 - 1}, Player{Position: pos2 - 1}
	for {
		p1 = p1.Move(d.RollN(3))
		if p1.Score >= 1000 {
			return p2.Score * d.N()
		}
		p2 = p2.Move(d.RollN(3))
		if p2.Score >= 1000 {
			return p1.Score * d.N()
		}
	}
}

// diracX3 shows how many universes are created per cumulative score after three
// rolls of the Dirac Dice.
//
// Maps roll score -> new universes.
var diracX3 = map[int]int64{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

type DiracState struct {
	P1, P2 Player
}

func PlayDirac(pos1, pos2 int) int64 {
	var p1Wins, p2Wins int64
	initState := DiracState{
		P1: Player{Position: pos1 - 1},
		P2: Player{Position: pos2 - 1},
	}
	states := make(map[DiracState]int64)
	states[initState]++
	for len(states) > 0 {
		// player 1 roll
		nextStates := make(map[DiracState]int64)
		for state, count := range states {
			for roll, newCount := range diracX3 {
				newState := DiracState{
					P1: state.P1.Move(roll),
					P2: state.P2,
				}
				if newState.P1.Score >= 21 {
					p1Wins += count * newCount
				} else {
					nextStates[newState] += count * newCount
				}
			}
		}
		states = nextStates

		// player 2 roll
		nextStates = make(map[DiracState]int64)
		for state, count := range states {
			for roll, newCount := range diracX3 {
				newState := DiracState{
					P1: state.P1,
					P2: state.P2.Move(roll),
				}
				if newState.P2.Score >= 21 {
					p2Wins += count * newCount
				} else {
					nextStates[newState] += count * newCount
				}
			}
		}
		states = nextStates
	}
	if p1Wins > p2Wins {
		return p1Wins
	}
	return p2Wins
}

package day23

import (
	"bytes"
	"container/heap"
	"fmt"
	"io"
)

func costFactor(c byte) int {
	switch c {
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	default:
		return 1
	}
}

type State struct {
	Data     [27]byte
	RoomSize int
	Cost     int
}

func NewState(roomSize int, startPositions string) State {
	if len(startPositions) != roomSize*4 {
		panic("bad init")
	}
	s := State{RoomSize: roomSize}
	for i, c := range []byte(startPositions) {
		room := i / roomSize
		slot := i % roomSize
		s.Data[11+(room*4)+slot] = c
	}
	return s
}

func (s State) isAmphipod(i int) bool {
	if i < 0 || i > len(s.Data) {
		return false
	}
	return s.Data[i] >= 'A' && s.Data[i] <= 'D'
}

func (s State) isInBuffer(i int) bool { return i >= 0 && i < 11 }

func (s State) isInRoom(i int) bool {
	if i < 11 || i >= len(s.Data) || (i-11)%4 >= s.RoomSize {
		return false
	}
	return true
}

// inRoom returns the identity in [0..4) of the room containing i.
func (s State) roomID(i int) int {
	if !s.isInRoom(i) {
		panic("not in a room")
	}
	return (i - 11) / 4
}

func (s State) roomWith(i int) int  { return 11 + 4*s.roomID(i) }
func (s State) roomName(i int) byte { return 'A' + byte(s.roomID(i)) }
func (s State) roomSlot(i int) int  { return i - s.roomWith(i) }
func (s State) homeFor(c byte) int  { return 11 + int(c-'A')*4 }

// portalTo returns the space in the buffer zone that is a portal to the room
// that contains i. If i is in the buffer zone, i is returned.
func (s State) portalTo(i int) int {
	if !s.isInRoom(i) {
		return i // no portal needed
	}
	return 2 + (s.roomID(i) * 2)
}

// findHome returns the first available slot in the destination room.
func (s State) findHome(i int) (v State, ok bool) {
	if !s.isAmphipod(i) {
		return
	}
	c := s.Data[i]
	room := s.homeFor(c)

	// find a slot to land in
	for j := 0; j < s.RoomSize; j++ {
		p := room + s.RoomSize - 1 - j
		if s.Data[p] == 0 {
			return s.tryMove(i, p)
		} else if s.Data[p] != c {
			return // foreigner
		}
	}
	return
}

// isHome returns true if the amphipod at i is already home.
func (s State) isHome(i int) bool {
	if !s.isInRoom(i) {
		return false
	}
	c := s.Data[i]
	if c != s.roomName(i) {
		return false // amphipod is in the wrong home
	}
	room := s.roomWith(i)
	for j := 0; j < s.RoomSize; j++ {
		switch s.Data[room+j] {
		case 0, c:
			continue
		default:
			return false
		}
	}
	return true
}

func (s State) pathTo(from, to int) (cost int, ok bool) {
	if s.Data[to] != 0 {
		return
	}

	// try leave the room
	if s.isInRoom(from) {
		room := s.roomWith(from)
		slot := s.roomSlot(from)
		cost += slot + 1
		for i := 0; i < slot; i++ {
			if s.Data[room+i] != 0 {
				return
			}
		}
	}

	// try traverse the buffer portion of the path
	pFrom, pTo := s.portalTo(from), s.portalTo(to)
	if pFrom < pTo {
		cost += pTo - pFrom
		for i := pFrom + 1; i < pTo; i++ {
			if s.Data[i] != 0 {
				return
			}
		}
	} else {
		cost += pFrom - pTo
		for i := pTo; i < pFrom; i++ {
			if s.Data[i] != 0 {
				return
			}
		}
	}

	// try enter the room
	if s.isInRoom(to) {
		room := s.roomWith(to)
		cost += to - room + 1
		for i := room; i < to; i++ {
			if s.Data[i] != 0 {
				return
			}
		}
	}
	ok = true
	return
}

func (s State) isPortal(i int) bool {
	switch i {
	case 2, 4, 6, 8:
		return true
	default:
		return false
	}
}

func (s State) tryMove(from, to int) (v State, ok bool) {
	if from == to {
		return // not moving
	}
	if s.isPortal(to) {
		return // illegal
	}
	if !s.isAmphipod(from) {
		return // nothing to move
	}
	if s.isHome(from) {
		return // already home
	}
	c := s.Data[from]
	if s.isInBuffer(from) {
		if s.isInBuffer(to) {
			return // buffer to buffer not allowed
		}
		if c != s.roomName(to) {
			return // wrong room
		}
	}
	cost, ok := s.pathTo(from, to)
	if !ok {
		return
	}
	v = s
	v.Cost = s.Cost + (cost * costFactor(c))
	v.Data[from], v.Data[to] = 0, c
	// TODO: runtime assertions?
	ok = true
	return
}

// Next produces all possible states that are reachable from this state.
func (s State) Next() []State {
	a := make([]State, 0, 8)
	for from := range s.Data {
		if !s.isAmphipod(from) {
			continue
		}
		// try take amphipod home
		if v, ok := s.findHome(from); ok {
			a = append(a, v)
		}
		if s.isInBuffer(from) {
			// cannot move from buffer to buffer
			continue
		}
		for to := 0; to < 11; to++ {
			// try to move to each buffer space
			if v, ok := s.tryMove(from, to); ok {
				a = append(a, v)
			}
		}
	}
	return a
}

func (s State) Organize(target State) (cost int) {
	q := &StateHeap{s}
	heap.Init(q)
	seen := make(map[[27]byte]State, 4096)
	for q.Len() > 0 {
		s := q.Pop().(State)
		if _, ok := seen[s.Data]; ok {
			continue
		}
		seen[s.Data] = s
		if s.Data == target.Data {
			if cost == 0 || s.Cost < cost {
				cost = s.Cost
			}
			continue
		}
		for _, nextState := range s.Next() {
			seenState, ok := seen[nextState.Data]
			if !ok {
				heap.Push(q, nextState)
				continue
			}
			if seenState.Cost <= nextState.Cost {
				continue
			}
			delete(seen, nextState.Data)
			heap.Push(q, nextState)
		}
	}
	return
}

func (s State) Format(w io.Writer) {
	getChar := func(i int) byte {
		if c := s.Data[i]; c != 0 {
			return c
		}
		return '.'
	}
	fmt.Fprintf(w, "--- State (%d)---\n", s.Cost)
	fmt.Fprint(w, "#############\n#")
	for i := 0; i < 11; i++ {
		fmt.Fprintf(w, "%c", getChar(i))
	}
	fmt.Fprint(w, "#\n")
	fmt.Fprintf(
		w,
		"###%c#%c#%c#%c###\n",
		getChar(11),
		getChar(15),
		getChar(19),
		getChar(23),
	)
	for i := 1; i < s.RoomSize; i++ {
		fmt.Fprintf(
			w,
			"  #%c#%c#%c#%c#\n",
			getChar(11+i),
			getChar(15+i),
			getChar(19+i),
			getChar(23+i),
		)
	}
	fmt.Fprint(w, "  #########\n")
}

func (s State) String() string {
	b := new(bytes.Buffer)
	s.Format(b)
	return b.String()
}

// min-queue for state objects, ordered by cost
type StateHeap []State

func (h StateHeap) Len() int            { return len(h) }
func (h StateHeap) Less(i, j int) bool  { return h[i].Cost < h[j].Cost }
func (h StateHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *StateHeap) Push(x interface{}) { *h = append(*h, x.(State)) }

func (h *StateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

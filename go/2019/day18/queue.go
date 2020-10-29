package main

import "container/heap"

type StateQueue interface {
	Len() int
	Enqueue(v State)
	Dequeue() State
}

type stateQueue struct {
	data []State
}

func NewStateQueue() StateQueue {
	return &stateQueue{data: make([]State, 0, 64)}
}

func (h *stateQueue) Len() int {
	return len(h.data)
}
func (h *stateQueue) Enqueue(v State) {
	h.data = append(h.data, v)
}

func (h *stateQueue) Dequeue() State {
	v := h.data[0]
	h.data = h.data[1:]
	return v
}

type stateHeap struct {
	data []State
}

func NewStateHeap() StateQueue {
	v := &stateHeap{
		data: make([]State, 0, 64),
	}
	heap.Init(v)
	return v
}

func (h *stateHeap) Enqueue(x State) {
	heap.Push(h, x)
}

func (h *stateHeap) Dequeue() State {
	return heap.Pop(h).(State)
}

func (h *stateHeap) Len() int           { return len(h.data) }
func (h *stateHeap) Less(i, j int) bool { return h.data[i].Cost < h.data[j].Cost }
func (h *stateHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *stateHeap) Push(x interface{}) { h.data = append(h.data, x.(State)) }
func (h *stateHeap) Pop() interface{} {
	n := len(h.data)
	v := h.data[n-1]
	h.data = h.data[0 : n-1]
	return v
}

type stateStack struct {
	data []State
}

func NewStateStack() StateQueue {
	return &stateStack{data: make([]State, 0, 64)}
}
func (h *stateStack) Len() int { return len(h.data) }

func (h *stateStack) Enqueue(v State) {
	h.data = append(h.data, v)
}

func (h *stateStack) Dequeue() State {
	n := len(h.data)
	v := h.data[n-1]
	h.data = h.data[0 : n-1]
	return v
}

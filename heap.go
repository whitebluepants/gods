package main

import (
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	p := h.IntSlice[:h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return p
}

func (h *hp) Top() int {
	return h.IntSlice[0]
}

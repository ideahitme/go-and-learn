package binaryheap

//go:generate go run gen.go

import (
	"sort"
)

type (
	// GenericBinaryHeap ...
	GenericBinaryHeap struct {
		data Interface
	}
	// Interface ...
	Interface interface {
		sort.Interface
		Pop() interface{}
		Push(interface{})
	}
)

// NewGeneric returns BinaryHeap implementation
func NewGeneric(data Interface) *GenericBinaryHeap {
	return &GenericBinaryHeap{
		data: data,
	}
}

// Insert assumes it is called after new element is added, it will place newly added element to the right position
func (bh *GenericBinaryHeap) Insert(x interface{}) {
	bh.data.Push(x)
	i := bh.data.Len() - 1
	for i > 0 {
		j := (i - 1) / 2
		if !bh.data.Less(j, i) {
			break
		}
		bh.data.Swap(i, j)
		i = j
	}
}

// Pop removes the maximum value in the binary heap and rebalances it
func (bh *GenericBinaryHeap) Pop() interface{} {
	bh.data.Swap(0, bh.data.Len()-1)
	top := bh.data.Pop()
	i := 0
	size := bh.data.Len()
	for {
		lc := 2*i + 1
		x := i
		if lc >= size {
			break
		}
		if bh.data.Less(x, lc) {
			x = lc
		}
		if lc+1 < size && bh.data.Less(x, lc+1) {
			x = lc + 1
		}
		if x == i {
			break
		}
		bh.data.Swap(x, i)
		i = x
	}
	return top
}

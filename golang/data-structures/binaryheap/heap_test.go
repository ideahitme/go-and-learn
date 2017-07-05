package binaryheap

import (
	"container/heap"
	"math/rand"
	"testing"
	"time"
)

var testdata []int

func Setup() {
	if len(testdata) > 0 {
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1e5; i++ {
		testdata = append(testdata, rand.Intn(1e6))
	}
	return
}

func BenchmarkBinaryHeap(b *testing.B) {
	Setup()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bh := NewBHInt()
		bh.WithPresize(1e5)
		for j := 0; j < 1e5; j++ {
			bh.Insert(testdata[j])
			// bh.Top()
			if j%2 == 1 {
				bh.Pop()
			}
		}
	}
}

type HeapInts []int

func (h HeapInts) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h HeapInts) Len() int {
	return len(h)
}

func (h HeapInts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapInts) Pop() interface{} {
	size := len(*h)
	item := (*h)[size-1]
	(*h) = (*h)[0 : size-1]
	return item
}

func (h *HeapInts) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}

func BenchmarkGenericBinaryHeap(b *testing.B) {
	Setup()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := HeapInts([]int{})
		bh := NewGeneric(&data)
		for j := 0; j < 1e5; j++ {
			bh.Insert(testdata[j])
			if j%2 == 1 {
				bh.Pop()
			}
		}
	}
}

func BenchmarkStdBinaryHeap(b *testing.B) {
	Setup()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := HeapInts([]int{})
		heap.Init(&data)
		for j := 0; j < 1e5; j++ {
			heap.Push(&data, testdata[j])
			if j%2 == 1 {
				heap.Pop(&data)
			}
		}
	}
}

func TestBinaryHeap(t *testing.T) {
	raw := make([]int, 1e3)
	x := []int{}
	for i := 0; i < 1e5; i++ {
		x = append(x, rand.Intn(1e3))
	}
	bh := NewBHInt()

	for i := 0; i < 1e5; i++ {
		add(x[i], raw)
		bh.Insert(x[i])
		if max(raw) != bh.Top() {
			t.Fatalf("maximum elements are not correct %d != %d", max(raw), bh.Top())
		}
		if i%2 == 1 {
			if pop(raw) != bh.Pop() {
				t.Fatal("maximum elements after pop are not correct")
			}
		}
	}
}

func TestGenericBinaryHeap(t *testing.T) {
	raw := make([]int, 1e3)
	x := []int{}
	for i := 0; i < 1e5; i++ {
		x = append(x, rand.Intn(1e3))
	}

	data := HeapInts([]int{})
	bh := NewGeneric(&data)

	for i := 0; i < 1e5; i++ {
		add(x[i], raw)
		bh.Insert(x[i])
		if max(raw) != data[0] {
			t.Fatalf("maximum elements are not correct %d != %d", max(raw), data[0])
		}
		if i%2 == 1 {
			if pop(raw) != bh.Pop() {
				t.Fatal("maximum elements after pop are not correct")
			}
		}
	}
}

// helper methods for rock-solid verifications
func drop(x int, from []int) {
	from[x]--
}

func add(x int, from []int) {
	from[x]++
}

func pop(from []int) int {
	x := max(from)
	drop(x, from)
	return x
}

func max(in []int) int {
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] > 0 {
			return i
		}
	}
	return 0
}

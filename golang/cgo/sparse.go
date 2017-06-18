package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

type Set interface {
	Insert(uint32)     // Insert new element if not present
	Find(uint32) bool  // Check if element is in the set
	Iterate() []uint32 // Provide a way to iterate over inserted elements
	Clear()            // Clear the set
}

var _ Set = MapSet{}
var _ Set = &NaiveSparseSet{}
var _ Set = &SparseSet{}

// SparseSet implementation
// size is the number of elements currently stored
type SparseSet struct {
	size   int32
	dense  []uint32
	sparse unsafe.Pointer
}

// NewSparseSet initializes SparseSet
// maxValue is the maximum integer value possibly present in the set
func NewSparseSet(maxValue uint32) (*SparseSet, error) {
	//allocate a new array via malloc
	sparse := C.malloc(C.size_t(maxValue * 4))
	if sparse == nil {
		return nil, fmt.Errorf("failed to allocate memory")
	}
	return &SparseSet{sparse: unsafe.Pointer(sparse)}, nil
}

func (s *SparseSet) Free() {
	C.free(s.sparse)
}

func (s *SparseSet) Clear() {
	s.size = 0
}

func (s *SparseSet) Insert(x uint32) {
	mem := (*int32)(s.at(x))
	if s.Find(x) {
		return //already in the set
	}
	if int32(len(s.dense)) == s.size { //grow only when necessary
		s.dense = append(s.dense, x)
	} else {
		s.dense[s.size] = x
	}
	*mem = s.size
	s.size++
}

func (s *SparseSet) Find(x uint32) bool {
	mem := (*int32)(s.at(x))
	return *mem < s.size && s.dense[*mem] == x
}

func (s *SparseSet) Iterate() []uint32 {
	return s.dense
}

// at(i) returns pointer to sparse array in position i
func (s *SparseSet) at(shift uint32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(s.sparse) + uintptr(4*shift))
}

type MapSet map[uint32]bool

func (ms MapSet) Insert(x uint32) {
	ms[x] = true
}

func (ms MapSet) Find(x uint32) bool {
	return ms[x]
}

func (ms MapSet) Iterate() []uint32 {
	result := make([]uint32, 0, len(ms))
	for x := range ms {
		result = append(result, x)
	}
	return result
}

func (ms MapSet) Clear() {
	for x := range ms {
		delete(ms, x)
	}
}

type NaiveSparseSet struct {
	dense  []uint32
	sparse []int
}

func NewNaiveSparseSet(maxValue int) *NaiveSparseSet {
	sparse := make([]int, maxValue)
	return &NaiveSparseSet{
		sparse: sparse,
	}
}

func (ns *NaiveSparseSet) Insert(x uint32) {
	if ns.Find(x) {
		return
	}
	ns.sparse[x] = len(ns.dense)
	ns.dense = append(ns.dense, x)
}

func (ns *NaiveSparseSet) Find(x uint32) bool {
	return ns.sparse[x] < len(ns.dense) && ns.dense[ns.sparse[x]] == x
}

func (ns *NaiveSparseSet) Iterate() []uint32 {
	return ns.dense
}

func (ns *NaiveSparseSet) Clear() {
	ns.dense = nil
}

func main() {
	s, err := NewSparseSet(100000)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Free()
	s.Insert(100)
	fmt.Printf("is %d in the set: %t\n", 100, s.Find(100))
	fmt.Printf("is %d in the set: %t\n", 99, s.Find(99))
	s.Insert(99)
	fmt.Printf("is %d in the set: %t\n", 99, s.Find(99))

	s.Clear()
	fmt.Printf("is %d in the set: %t\n", 100, s.Find(100))
	s.Insert(1)
	s.Insert(5)
	s.Insert(4)
	fmt.Println(s.Iterate())

	ms := MapSet{}
	ms.Insert(1)
	ms.Insert(2)
	ms.Insert(3)
	fmt.Println(ms.Iterate())

	ns := NewNaiveSparseSet(1000)
	ns.Insert(1)
}

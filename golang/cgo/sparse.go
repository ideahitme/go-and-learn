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

func main() {
	s, err := New(1000)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Free()

	s.Insert(100)
	fmt.Printf("is %d in the set: %t\n", 100, s.Find(100))
	fmt.Printf("is %d in the set: %t\n", 99, s.Find(99))

	s.Clear()
	fmt.Printf("is %d in the set: %t\n", 100, s.Find(100))
}

// SparseSet implementation
// size is the number of elements currently stored
//
type SparseSet struct {
	size   int32
	dense  []uint32
	sparse unsafe.Pointer
}

// New initializes SparseSet it operates under assumption integers are 4 bytes as most commonly implemented in C compilers
// maxInt is the maximum integer value to be present in the set
func New(maxInt uint32) (*SparseSet, error) {
	//allocate a new array via malloc
	sparse := C.malloc(C.size_t(maxInt)*4 + 1)
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

func (s *SparseSet) Insert(i uint32) {
	s.dense = append(s.dense, i)
	mem := (*int32)(s.at(i))
	*mem = s.size
	s.size++
}

func (s *SparseSet) Find(i uint32) bool {
	mem := (*int32)(s.at(i))
	return *mem < s.size && s.dense[*mem] == i
}

// at(i) returns pointer to sparse array in position i
func (s *SparseSet) at(shift uint32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(s.sparse) + uintptr(4*shift))
}

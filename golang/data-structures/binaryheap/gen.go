// +build ignore

package main

import (
	"os"
	"strings"
	"text/template"
)

var (
	types = []string{"int", "int32", "int64", "uint16", "uint32", "uint64", "float64"}

	header = `package binaryheap
`

	tmplString = `
// BH{{.CapType}} binary heap implementation storing {{.Type}}
type BH{{.CapType}} struct {
	arr []{{.Type}} // x stores all elements in the binary heap
}

// NewBH{{.CapType}} returns BinaryHeap implementation storing {{.Type}}
func NewBH{{.CapType}}() *BH{{.CapType}} {
	return &BH{{.CapType}}{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BH{{.CapType}}) WithPresize(size uint64) *BH{{.CapType}} {
	x := make([]{{.Type}}, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BH{{.CapType}}) Insert(x {{.Type}}) {
	bh.arr = append(bh.arr, x)
	i := len(bh.arr) - 1
	for i > 0 {
		j := (i - 1) / 2
		if bh.arr[i] <= bh.arr[j] {
			break
		}
		bh.arr[i], bh.arr[j] = bh.arr[j], bh.arr[i]
		i = j
	}
}

// Top returns maximum value in the binary heap
func (bh *BH{{.CapType}}) Top() {{.Type}} {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BH{{.CapType}}) Pop() {{.Type}} {
	top := bh.arr[0]
	bh.arr[0] = bh.arr[len(bh.arr)-1]
	bh.arr = bh.arr[0 : len(bh.arr)-1]
	i := 0
	size := len(bh.arr)
	for i < size {
		x := i

		lc := 2*i + 1
		if lc >= size {
			break
		}
		if bh.arr[x] < bh.arr[lc] {
			x = lc
		}

		if lc+1 < size && bh.arr[x] < bh.arr[lc+1] {
			x = lc + 1
		}

		if x == i {
			break
		}

		bh.arr[x], bh.arr[i] = bh.arr[i], bh.arr[x]
		i = x
	}
	return top
}
`
)

func main() {
	f, err := os.Create("./typed.go")
	if err != nil {
		panic(err)
	}
	f.WriteString(header)
	defer f.Close()

	for _, t := range types {
		tmpl, err := template.New("BinaryTreeTemplate").Parse(tmplString)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(f, struct {
			Type    string
			CapType string
		}{
			Type:    t,
			CapType: strings.Title(t),
		})
	}
}

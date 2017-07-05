package binaryheap

// BHInt binary heap implementation storing int
type BHInt struct {
	arr []int // x stores all elements in the binary heap
}

// NewBHInt returns BinaryHeap implementation storing int
func NewBHInt() *BHInt {
	return &BHInt{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHInt) WithPresize(size uint64) *BHInt {
	x := make([]int, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHInt) Insert(x int) {
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
func (bh *BHInt) Top() int {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHInt) Pop() int {
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

// BHInt32 binary heap implementation storing int32
type BHInt32 struct {
	arr []int32 // x stores all elements in the binary heap
}

// NewBHInt32 returns BinaryHeap implementation storing int32
func NewBHInt32() *BHInt32 {
	return &BHInt32{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHInt32) WithPresize(size uint64) *BHInt32 {
	x := make([]int32, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHInt32) Insert(x int32) {
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
func (bh *BHInt32) Top() int32 {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHInt32) Pop() int32 {
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

// BHInt64 binary heap implementation storing int64
type BHInt64 struct {
	arr []int64 // x stores all elements in the binary heap
}

// NewBHInt64 returns BinaryHeap implementation storing int64
func NewBHInt64() *BHInt64 {
	return &BHInt64{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHInt64) WithPresize(size uint64) *BHInt64 {
	x := make([]int64, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHInt64) Insert(x int64) {
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
func (bh *BHInt64) Top() int64 {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHInt64) Pop() int64 {
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

// BHUint16 binary heap implementation storing uint16
type BHUint16 struct {
	arr []uint16 // x stores all elements in the binary heap
}

// NewBHUint16 returns BinaryHeap implementation storing uint16
func NewBHUint16() *BHUint16 {
	return &BHUint16{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHUint16) WithPresize(size uint64) *BHUint16 {
	x := make([]uint16, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHUint16) Insert(x uint16) {
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
func (bh *BHUint16) Top() uint16 {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHUint16) Pop() uint16 {
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

// BHUint32 binary heap implementation storing uint32
type BHUint32 struct {
	arr []uint32 // x stores all elements in the binary heap
}

// NewBHUint32 returns BinaryHeap implementation storing uint32
func NewBHUint32() *BHUint32 {
	return &BHUint32{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHUint32) WithPresize(size uint64) *BHUint32 {
	x := make([]uint32, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHUint32) Insert(x uint32) {
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
func (bh *BHUint32) Top() uint32 {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHUint32) Pop() uint32 {
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

// BHUint64 binary heap implementation storing uint64
type BHUint64 struct {
	arr []uint64 // x stores all elements in the binary heap
}

// NewBHUint64 returns BinaryHeap implementation storing uint64
func NewBHUint64() *BHUint64 {
	return &BHUint64{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHUint64) WithPresize(size uint64) *BHUint64 {
	x := make([]uint64, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHUint64) Insert(x uint64) {
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
func (bh *BHUint64) Top() uint64 {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHUint64) Pop() uint64 {
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

// BHFloat64 binary heap implementation storing float64
type BHFloat64 struct {
	arr []float64 // x stores all elements in the binary heap
}

// NewBHFloat64 returns BinaryHeap implementation storing float64
func NewBHFloat64() *BHFloat64 {
	return &BHFloat64{}
}

// WithPresize allows to increase the size of the internal array for additional performance
func (bh *BHFloat64) WithPresize(size uint64) *BHFloat64 {
	x := make([]float64, len(bh.arr), size)
	copy(x, bh.arr)
	bh.arr = x
	return bh
}

// Insert adds new value to binary heap and places it in the correct position in the heap
func (bh *BHFloat64) Insert(x float64) {
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
func (bh *BHFloat64) Top() float64 {
	return bh.arr[0]
}

// Pop returns the maximum value in the binary heap and pops it up
func (bh *BHFloat64) Pop() float64 {
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

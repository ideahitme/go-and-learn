package main

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

const sizeOfInt = unsafe.Sizeof(0)

// uintptr is an integer type that is large enough to hold the bit pattern of any pointer.

func main() {
	fmt.Printf("size of int: %d\n", sizeOfInt)

	alignOf()
	offsetOf()
	sizeOf()

	// convert type of slice
	b := []byte{1, 1, 0, 0, 5, 0, 0, 0}
	sizeOfInt32 := 4 //bytes
	bheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	bheader.Cap /= sizeOfInt32
	bheader.Len /= sizeOfInt32

	sliceOfInts := *(*[]int32)(unsafe.Pointer(&bheader))
	fmt.Println(sliceOfInts) //[257 5] - depends on endian-ness of the machine

	modifyStringInPlace()
}

// type                               alignment guarantee
// ------                             ------
// bool, byte, uint8, int8            1
// uint16, int16                      2
// uint32, int32, float32, complex64  4
// other types                        size of a native word

/**
Experiment with AlignOf func from unsafe package
*/

func alignOf() {
	var i int
	intBoundary := unsafe.Alignof(i)
	fmt.Println(intBoundary) //on 64 bit machine -> 8 bytes, 32 bit - 4 bytes

	var i64 int
	int64Boundary := unsafe.Alignof(i64)
	fmt.Println(int64Boundary) // 8 bytes

	/**
	For a variable x of struct type: unsafe.Alignof(x) is the largest
	of all the values unsafe.Alignof(x.f) for each field f of x, but at least 1.
	For a variable x of array type: unsafe.Alignof(x) is the same as the alignment
	of a variable of the array's element type..
	*/
	type User struct {
		active bool
		age    int
		money  float64
	}
	u := User{true, 10, 10000.0}
	fmt.Println(unsafe.Alignof(u)) //8

	type Count struct {
		id uint8 //alignment is 1
	}

	fmt.Println(unsafe.Alignof(Count{})) //1

}

func offsetOf() {
	// Offsetof returns the offset within the struct of the field represented by x,
	// which must be of the form structValue.field.
	// In other words, it returns the number of bytes between the start of the struct and the start of the field.
	type person struct {
		name    string
		age     int64
		married bool
	}

	adam := &person{"adam", 32, true}

	println("offset of name:", unsafe.Offsetof(adam.name))       // 0 .. 4 * 4 -1
	println("offset of age:", unsafe.Offsetof(adam.age))         // 4*4..4*4+8-1
	println("offset of married:", unsafe.Offsetof(adam.married)) //24..25

	padam := unsafe.Pointer(adam)
	pmarried := (*bool)(unsafe.Pointer(uintptr(padam) + unsafe.Offsetof(adam.married)))
	*pmarried = false

	fmt.Println(*adam) //divorced O_O
}

func sizeOf() {
	x := "abcd"
	println("size of abcd", unsafe.Sizeof(x))

	s := []int{1, 2, 3, 4}
	println("size of slice", unsafe.Sizeof(s))
	// prints 24 because:
	// slice is uinptr (8 bytes) + Len (8 bytes) + Cap (8 bytes)

	a := [5]int{1, 2, 3, 4, 5}
	println("size of array", unsafe.Sizeof(a))
	// 40 = 5 * 8

	user := struct {
		name string
		age  int64
	}{
		name: "dude",
		age:  55,
	}
	println("size of struct", unsafe.Sizeof(user)) // 4*4+8 = 24
}

/**
use unsafe to modify string in place
*/

func modifyStringInPlace() {
	s := "1234"
	// let's change s to 1235

	stringHeader := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	// stringHeader.Data points to the actual data
	// disable memory protection, because string are read-only
	setMemoryProtect(stringHeader.Data, true)
	defer setMemoryProtect(stringHeader.Data, false) //enable it back

	lastEl := (*byte)(unsafe.Pointer(stringHeader.Data + 1*uintptr(3))) //now lastEl points to the last bytes of string
	*lastEl = '5'

	fmt.Println(s) // 1235
}

func setMemoryProtect(ptr uintptr, w bool) {
	// make sure start is a multiple of page size
	start := ptr & ^(uintptr(syscall.Getpagesize() - 1)) //clears last bits, e.g. on mac pagesize is 4096 = 2^12, hence last 11 bits are cleared
	prot := syscall.PROT_READ
	if w {
		prot |= syscall.PROT_WRITE
	}

	_, _, err := syscall.Syscall(
		syscall.SYS_MPROTECT,
		start, uintptr(syscall.Getpagesize()),
		uintptr(prot),
	) // this would work only if string occupies less than pagesize
	if err != 0 {
		panic(err.Error())
	}
}

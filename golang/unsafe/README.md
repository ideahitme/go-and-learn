### Internals of unsafe package

### unsafe

Exposes three functions: 

1. `func Alignof(x ArbitraryType) uintptr`
2. `func Offsetof(x ArbitraryType) uintptr`
3. `func Sizeof(x ArbitraryType) uintptr`

All three functions are evaluated at **compile** time. 

*see main.go for usage of 1, 2, 3*

Because they are evaluated at compile time this is valid: 

```go
const sizeOfInt = unsafe.Sizeof(1)
```

It also exposes on type: `unsafe.Pointer`. With its help it is possible to convert any pointer type to any pointer type, i.e. it allows to bypass Go type system

```go

  n := 10
  ptrN := &n

  var f *float64
  f = (*float64)(unsafe.Pointer(ptrN))
  fmt.Println(*f)

  *f = 3.1415

  fmt.Println(*f)
  fmt.Println(n) //4614256447914709615

  //convery slice data to another type without copy
  b := []byte{1, 1, 0, 0, 5, 0, 0, 0}
  sizeOfInt32 := 4 //bytes
  bheader := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
  bheader.Cap /= sizeOfInt32
  bheader.Len /= sizeOfInt32

  sliceOfInts := *(*[]int32)(unsafe.Pointer(&bheader))
  fmt.Println(sliceOfInts) //[257 5] - depends on endian-ness of the machine

```
> 1. uintptr is an integer type.
> 2. the memory block pointed by a uintptr value can be garbaged collected even if the uintptr value is still alive.
> 3. unsafe.Pointer is a pointer type, but unsafe.Pointer values can't be dereferenced.
> 4. the memory block pointed by a unsafe.Pointer value will not be garbage collected if the unsafe.Pointer value is still alive.
> 5. *unsafe.Pointer is a general pointer type, just like *int etc. This means values of *unsafe.Pointer can be converted to unsafe.Pointer, and vice versa.

Rules for `uintptr`, `unsafe.Pointer` and other pointer types: 

> - A pointer value of any type can be converted to a Pointer.
> - A Pointer can be converted to a pointer value of any type.
> - A uintptr can be converted to a Pointer.
> - A Pointer can be converted to a uintptr.


#### Pointer arithmetic 

Example of pointer arithmetic and getting slice underlying array: 
```go
package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	p := make([]int, 0, 2)
	p = append(p, 1)
	byValue(p, 2)
	fmt.Println(p)	// [1]
	
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&p)) 
	firstEl := (*int)(unsafe.Pointer(header.Data)) //header.Data is uintptr to array
	println(*firstEl) // 1 - is the first element
	
	secondEl := (*int)(unsafe.Pointer(header.Data + 1 * unsafe.Sizeof(p[0])))
	println(*secondEl) // 2 - it is actually there
}

func byValue(p []int, t int) {
	p = append(p, t)	
}
```


#### Articles 

1. http://bouk.co/blog/monkey-patching-in-go/ - monkey patching example (replacing function in the runtime)
2. http://www.tapirgames.com/blog/golang-unsafe
3. http://learngowith.me/gos-pointer-pointer-type/
4. http://www.tapirgames.com/blog/golang-memory-alignment
5. https://golang.org/pkg/unsafe/

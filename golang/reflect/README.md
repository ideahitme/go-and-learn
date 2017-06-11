## Reflect package

> In Golang reflection allows inspection of struct, interfaces, fields and methods at runtime without knowing the names of the interfaces, fields, methods at compile time. It also allows instantiation of new objects and invocation of methods.

### Laws of reflection

1. Reflection goes from interface value to reflection object.

> A variable of interface type stores a pair: the concrete value assigned to the variable, and that value's type descriptor - https://research.swtch.com/interfaces 

Therefore, reflect allows to inspect value type (`reflect.Type`) and value (`reflect.Value`). Both can be extracted using `TypeOf` and `ValueOf` (see usage in main.go)

Both take a `interface{}` type parameter: 

```go

// TypeOf returns the reflection Type of the value in the interface{}.
func TypeOf(i interface{}) Type {
  ...
}

// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i. ValueOf(nil) returns the zero Value.
func ValueOf(i interface{}) Value { 
  ...
}

// example 

var t int64

fmt.Println(reflect.ValueOf(t).String()) // <int64 Value>
fmt.Println(reflect.TypeOf(t).String())  // int64
```

2. Reflection goes from reflection object to interface value.

`reflect.Value` exposes a method `Interface()` to recover unpacked interface from rule 1. 

```go
var x float64
v := reflect.ValueOf(x)
i := v.Interface()

cx, ok := i.(float64)
println(cx, ok)

```

3. To modify a reflection object, the value must be settable.

**What can be set**:

```go
	var x float64 = 3.14
	fmt.Println(reflect.ValueOf(x).CanSet())  //false
	fmt.Println(reflect.ValueOf(&x).CanSet()) //false
	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.
	fmt.Println(reflect.ValueOf(&x).Elem().CanSet()) //true
```

Let's understand why last statement returns true. 

`reflect.Value(&x)` will unpack the `interface{}`, and extract the value, which is a pointer to the `float64`:

```go
	fmt.Println(reflect.ValueOf(&x).String()) // <*float64 Value>
```

At this stage `reflect.Value` contains a `ptr` field:

>	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.

Now what happens in `Elem()` with dereferencing:

``` go
// how Value is evaluated for kind == reflect.Ptr in Elem()
ptr := v.ptr
if v.flag&flagIndir != 0 { //if indir bit is set
  ptr = *(*unsafe.Pointer)(ptr) //extracts the actual data
}
...
fl := v.flag&flagRO | flagIndir | flagAddr
fl |= flag(typ.Kind())
...
return Value{typ, ptr, fl}
```

Now let's see what `CanSet()` actually checks: 

```go
// CanSet reports whether the value of v can be changed.
// A Value can be changed only if it is addressable and was not
// obtained by the use of unexported struct fields.
// If CanSet returns false, calling Set or any type-specific
// setter (e.g., SetBool, SetInt) will panic.
func (v Value) CanSet() bool {
	return v.flag&(flagAddr|flagRO) == flagAddr
}
```
So `CanSet()` checks that `v.flag` RO (read-only) bit is **not** enabled (i.e. checks that field or method is exported) and `flagAddr` is true, i.e. type is addressable.

See below for flags (Bit Flags section).

**Therefore calling `Elem()` on `Value` holding a pointer will produce a new value with `flag.Addr` set to true. And this is not true for `reflect.ValueOf(&x)`, because only `flag.Indir` will be set for it, therefore it is not settable**. This is obviously correct because we do not want to set the value of the pointer, but rather the data it is pointing to.

See this code with more examples on settables:

```go
	fmt.Println(reflect.ValueOf(&x).Elem().CanSet()) //true

	var t *int64
	fmt.Println(reflect.ValueOf(t).CanSet())  //false
	fmt.Println(reflect.ValueOf(t).CanAddr()) //false

	u := user{Name: "123", id: "321"}
	fmt.Println(reflect.ValueOf(&u).Elem().Field(0).CanSet()) //true
	fmt.Println(reflect.ValueOf(&u).Elem().Field(1).CanSet()) //false
```

**Setting value**:

```go
var t int64 //0
i := reflect.ValueOf(&t).Elem()
i.SetInt64(1)
fmt.Println(t) //1
```

### Bit flags 

```go
// The lowest bits are flag bits:
//	- flagStickyRO: obtained via unexported not embedded field, so read-only
//	- flagEmbedRO: obtained via unexported embedded field, so read-only
//	- flagIndir: val holds a pointer to the data
//	- flagAddr: v.CanAddr is true (implies flagIndir)
//	- flagMethod: v is a method value.
...
//	flagRO = flagStickyRO | flagEmbedRO

```


### Extract methods 

```go

type WithName interface {
	Hello() string
	Param(param string)
}

type user struct {
	Name string
	id   string
}

func (u *user) Hello() string {
	return fmt.Sprintf("Hello, my name is: %s", u.Name)
}

func (u *user) Param(param string) {
	fmt.Printf("User %s got a param %s\n", u.Name, param)
}

...
	//call every single interface method on its typed implementation
	//check if user implements WithName
	var wn *WithName // or can be used as (*WithName)(nil)
	if reflect.TypeOf(u).Implements(reflect.TypeOf(wn).Elem()) {
		fmt.Println("user implements WithName interface")
		//call each method defined on interface
		for i := 0; i < reflect.TypeOf(wn).Elem().NumMethod(); i++ {
			m := reflect.TypeOf(wn).Elem().Method(i)
			params := []reflect.Value{}
			for j := 0; j < m.Type.NumIn(); j++ {
				in := m.Type.In(j)
				// this is not generic
				if in.Kind() == reflect.String {
					params = append(params, reflect.ValueOf("default"))
				}
			}
			fmt.Printf("Calling %s on %s: %s\n", m.Name, reflect.TypeOf(u), reflect.ValueOf(u).MethodByName(m.Name).Call(params))
		}
	}

```

### Channels

Allows to send/receive as per normal chan operators. It also exposes non-blocking APIs, i.e try send/receive. 

Try sending value

```go
	c := make(chan int, 1)
	// TrySend
	// It reports whether the value was sent.
	// As in Go, x's value must be assignable to the channel's element type.

	x := reflect.ValueOf(c).TrySend(reflect.ValueOf(1))
	fmt.Println("sending to channel: ", x) //true
	x = reflect.ValueOf(c).TrySend(reflect.ValueOf(1))
	fmt.Println("sending to channel: ", x) //false

	// nil channel
	var n chan int
	x = reflect.ValueOf(n).TrySend(reflect.ValueOf(1))
	fmt.Println("sending to nil channel: ", x) //false
```
Try receiving value - allows to check if channel is closed/nil

```go
	c := make(chan int, 1)

	x, ok := reflect.ValueOf(c).TryRecv() //no blocking
	if ok {
		//received a value
		fmt.Println("received a value")
		return
	}
	if !ok && x == (reflect.Value{}) {
		//channel not closed by nothing can be received
		fmt.Println("channel is blocking")
		return
	}
	//channel is closed
	fmt.Println("channel is closed")
}
```

### Notes

> The reflect package is great way to make descision at runtime. However, we should be aware that it gives us some performance penalties. I would try to avoid using reflection. It’s not idiomatic, but it’s very powerfull in particular cases. Do not forget to follow the laws of reflection.

### References

1. https://blog.golang.org/laws-of-reflection
2. http://blog.ralch.com/tutorial/golang-reflection/ 
3. https://research.swtch.com/interfaces


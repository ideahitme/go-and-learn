## Profiling and Memory management

*Based on: [Talk from Gophercon SG](https://www.youtube.com/watch?v=2557w0qsDV0&list=PLq2Nv-Sh8EbZEjZdPLaQt1qh_ohZFMDj8&index=3)*

Link: https://www.goinggo.net/2017/06/language-mechanics-on-memory-profiling.html

### Program byte stream transformer using io package:

All source code available in [stream/challenge.go](stream/challenge.go)

#### Correct code

Given a byte stream replace all occurrences of `find` with `replace`

```go
func Replace(data, find, replace []byte, output *bytes.Buffer) {
	stream := bytes.NewReader(data)
	size := len(find)
	i := 0

	var b byte
	var err error

	for {
		b, err = stream.ReadByte()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		if find[i] == b {
			if i == size-1 {
				output.Write(replace)
				i = 0
			} else {
				i++
			}
			continue
		}

		stream.Seek(-int64(i+1), io.SeekCurrent)

		b, _ = stream.ReadByte()
		output.WriteByte(b)

		i = 0
	}

	//write whatever is left
	output.Write(find[:i])
}
```

Run the following: 

```bash
$ cd challenge
$ go build -gcflags=-m # -m    print optimization decisions
./challenge.go:10: inlining call to bytes.NewReader
./challenge.go:18: inlining call to (*bytes.Reader).ReadByte
...
./challenge.go:10: Replace &bytes.Reader literal does not escape
...
```
*Nothing escapes to the heap everything can be inlined*

#### Problematic code
```go
func ReplaceProblematic(data, find, replace []byte, output *bytes.Buffer) {
	stream := bytes.NewReader(data)
	size := len(find)

	var b byte
	var err error
	var n int
	buf := make([]byte, size, size)

	for {
		n, err = io.ReadFull(stream, buf)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				stream.Seek(int64(-n), io.SeekCurrent)
			}
			break
		}

		if bytes.Compare(buf, find) == 0 {
			output.Write(replace)
			continue
		}

		stream.Seek(-int64(size), io.SeekCurrent)
		b, _ = stream.ReadByte()
		output.WriteByte(b)
	}
	//write whatever is left
	stream.WriteTo(output)
}

```
Let's benchmark two:

```bash
$ cd challenge
$ go test -run none -bench . -benchtime 5s -benchmem -memprofile mem.out
BenchmarkReplace-4              	 5000000	      1679 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplaceProblematic-4   	 2000000	      3501 ns/op	      53 B/op	       2 allocs/op
```
Now let's profile problematic algorithm:
```bash
$ go tool pprof -alloc_space challenge.test mem.out
Entering interactive mode (type "help" for commands)
(pprof) list Problematic
Total: 159.01MB
         .          .    143:
ROUTINE ======================== github.com/ideahitme/go-and-learn/golang/profiling/challenge.ReplaceProblematic in /Users/ytussupbekov/gows/src/github.com/ideahitme/go-and-learn/golang/profiling/challenge/challenge.go
  159.01MB   159.01MB (flat, cum)   100% of Total
         .          .     42:	output.Write(find[:i])
         .          .     43:}
         .          .     44:
         .          .     45:// ReplaceProblematic worse version of the Replace
         .          .     46:func ReplaceProblematic(data, find, replace []byte, output *bytes.Buffer) {
  144.51MB   144.51MB     47:	stream := bytes.NewReader(data)
         .          .     48:	size := len(find)
         .          .     49:
         .          .     50:	var b byte
         .          .     51:	var err error
         .          .     52:	var n int
   14.50MB    14.50MB     53:	buf := make([]byte, size, size)
         .          .     54:
         .          .     55:	for {
         .          .     56:		n, err = io.ReadFull(stream, buf)
         .          .     57:		if err != nil {
         .          .     58:			if err == io.ErrUnexpectedEOF {
(pprof)
```
**Problems:**

Check with `go build -gcflags -m` shows that 

1. `./challenge.go:47: &bytes.Reader literal escapes to heap`
2. `./challenge.go:53: make([]byte, size, size) escapes to heap`

Problem 2 is obvious because go compiler cannot make a decision to allocate dynamically sized arrays onto the stack. 

Problem 1 is less obvious. Let's try `go build -gcflags '-m -m'` (double -m) and now we can see:
```bash
./challenge.go:47: &bytes.Reader literal escapes to heap
./challenge.go:47: 	from ~r0 (assign-pair) at ./challenge.go:47
./challenge.go:47: 	from stream (assigned) at ./challenge.go:47
./challenge.go:47: 	from stream (interface-converted) at ./challenge.go:56
./challenge.go:47: 	from stream (passed to call[argument escapes]) at ./challenge.go:56
```

We know that `io.ReadFull` accepts `Reader` interface as its first parameter, therefore it always escapes - because compiler does not know implementation details of the interface implementation, therefore allocates it on the heap to be conservative. 


### CPU Profiling

```bash

$ go test -run none -bench . -benchtime 5s -benchmem -cpuprofile cpu.out
$ go tool pprof challenge.test cpu.out
Entering interactive mode (type "help" for commands)
(pprof) web
(pprof) list ReplaceProblematic
Total: 17.36s
ROUTINE ======================== github.com/ideahitme/go-and-learn/golang/profiling/challenge.BenchmarkReplaceProblematic in /Users/ytussupbekov/gows/src/github.com/ideahitme/go-and-learn/golang/profiling/challenge/challenge_test.go
      30ms      9.56s (flat, cum) 55.07% of Total
         .          .    133:	replace := []byte("Elvis")
         .          .    134:	data := assembleInputStream()
         .          .    135:	output := bytes.Buffer{}
         .          .    136:
         .          .    137:	b.ResetTimer()
      20ms       20ms    138:	for i := 0; i < b.N; i++ {
         .       20ms    139:		output.Reset()
      10ms      9.52s    140:		ReplaceProblematic(data, find, replace, &output)
         .          .    141:	}
         .          .    142:}
         .          .    143:
         .          .    144:func assembleInputStream() []byte {
         .          .    145:	var in []byte
ROUTINE ======================== github.com/ideahitme/go-and-learn/golang/profiling/challenge.ReplaceProblematic in /Users/ytussupbekov/gows/src/github.com/ideahitme/go-and-learn/golang/profiling/challenge/challenge.go
     2.12s      9.51s (flat, cum) 54.78% of Total
         .          .     41:	//write whatever is left
         .          .     42:	output.Write(find[:i])
         .          .     43:}
         .          .     44:
         .          .     45:// ReplaceProblematic worse version of the Replace
     260ms      260ms     46:func ReplaceProblematic(data, find, replace []byte, output *bytes.Buffer) {
     170ms      380ms     47:	stream := bytes.NewReader(data)
         .          .     48:	size := len(find)
         .          .     49:
         .          .     50:	var b byte
         .          .     51:	var err error
         .          .     52:	var n int
      10ms       60ms     53:	buf := make([]byte, size, size)
         .          .     54:
         .          .     55:	for {
     410ms      4.27s     56:		n, err = io.ReadFull(stream, buf)
      20ms       20ms     57:		if err != nil {
         .       20ms     58:			if err == io.ErrUnexpectedEOF {
         .          .     59:				stream.Seek(int64(-n), io.SeekCurrent)
         .          .     60:			}
         .          .     61:			break
         .          .     62:		}
         .          .     63:
     780ms      2.08s     64:		if bytes.Compare(buf, find) == 0 {
      40ms      450ms     65:			output.Write(replace)
         .          .     66:			continue
         .          .     67:		}
         .          .     68:
      70ms       70ms     69:		stream.Seek(-int64(size), io.SeekCurrent)
     310ms      310ms     70:		b, _ = stream.ReadByte()
      40ms      1.57s     71:		output.WriteByte(b)
         .          .     72:	}
         .          .     73:	//write whatever is left
      10ms       20ms     74:	stream.WriteTo(output)
         .          .     75:}


```

### Escape Analysis

1. Returning pointer to an object from a function results in the variable being allocated to a Heap memory: 

```go
package main

import "fmt"

type User struct {
	name string
}

var users []*User

func main() {
	u := createUser() // u points to heap memory address
	fmt.Println(u)

	cu := createCopyUser()
	fmt.Println(cu)

	user1 := &User{}
	doesNotEscape(user1)

	user2 := &User{}
	doesEscape(user2)

	user3 := &User{}
	doesEscapeV2(user3)
}

func createUser() *User {
	user := &User{"yerken"} //this is allocated in the heap,
	// because escape analysis determined variable pointer is returned from the function down below
	return user
}

func createCopyUser() User {
	user := &User{"yerken"} //this is allocated in the stack,
	return *user
}

func doesNotEscape(user *User) {
	user.name = ""
}

func doesEscape(user *User) {
	users = append(users, user)
}

func doesEscapeV2(user *User) *User {
	user.name = "123"
	doesEscape(user)
	return user
}
```

Now if we run main.go with gcflags as: 

```bash
# -m    print optimization decisions
# -l    disable inlining
# `go tool compile -help` to see more options
go run -gcflags '-m -l' main.go 
# command-line-arguments
./main.go:29: &User literal escapes to heap
./main.go:35: createCopyUser &User literal does not escape
./main.go:39: doesNotEscape user does not escape
./main.go:43: leaking param: user
./main.go:47: leaking param: user
./main.go:13: u escapes to heap
./main.go:16: cu escapes to heap
./main.go:21: &User literal escapes to heap
./main.go:24: &User literal escapes to heap
./main.go:13: main ... argument does not escape
./main.go:16: main ... argument does not escape
./main.go:18: main &User literal does not escape
&{yerken}
{yerken}
```
We can see **only** `user` defained in `createUser` function escapes to heap.

More on golang escape analysis [here](http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html)

List of escape false negatives: 
https://groups.google.com/forum/#!topic/golang-nuts/kXEpVo6Bek8
https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/view

Related talk on pragma's [here](https://www.youtube.com/watch?v=nmcPwqjPFbw)

### Flat vs Cum

```text
*Quote*:
Usually, the 'flat' part tells you how much time was spent in a given function/routine/... whereas the 'cum' part refers to 'cumulative' and contains the current function/routine/... plus everything "above"/"below" it.

Suppose we sample the program and whenever we do, we obtain the stack of the program. the bottom of the stack explains where we currently are, and the stack itself tells us where we came from. After a run, we can gather together all the samples and have a rough estimate of where the program is spending its time by looking at the bottom of the stack. This is the flat profile, usually.

However, the stack itself is also interesting, because if you have two call paths a->b->c->d and a->b->c->e, and you are spending 1 second in d and 1 second in e, then you could say you are spending 2 seconds "below" c. Hence the "(ac)cumulative" time spent in c is 2 seconds.

I don't know the exact definition of the pprof output, but usually this is what happens in a sampling profiler, and you need a way to quickly discern how to interpret the price of paths in the call-graph of the program. In the view of the graph, c above would "dominate" d and e (provided c is the only way to get to d and e!)
```

### Using runtime for profiling

See in `cmd/main.go`

### More:

1. https://medium.com/@hackintoshrao/daily-code-optimization-using-benchmarks-and-profiling-in-golang-gophercon-india-2016-talk-874c8b4dc3c5
2. https://www.youtube.com/watch?v=-KDRdz4S81U
3. http://github.com/pkg/profile - utility for runtime profiling
4. https://scene-si.org/2017/06/06/benchmarking-go-programs/?utm_source=golangweekly&utm_medium=email - has a paragraph about using flame graph
5. https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs
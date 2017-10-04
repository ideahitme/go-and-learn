# Go Pipe vs Buffer

## Task #1

You are a given a Reader implementation (assume `*bytes.Reader`) from which u need to produce a Reader with same order of bytes but `'a'` excluded. 

## Solution with Buffer

```go
func BufferedStream(b *bytes.Reader) io.Reader { 
	var buf *bytes.Buffer
	for {
		r, err := b.ReadByte()
		if err == io.EOF {
			break
		}
		if r != byte('a') {
			buf.WriteByte(r)
		}
	}

	return buf
}
```
This is the most straightforward solution, where we create a `bytes.Buffer` (which implements `ReadWriter` interface) and we use the buffer to write elements to it and then return the same buffer, from where we can read filtered bytes. 

> So we use an allocated chunk of memory for buffering data, i.e. exchange zone which will be recreated multiple times if the amount of data to be consumed is too large 

## Solution with Pipe

Alternative solution would be create a Writer and a Reader which could be piped. For which we can use `io.PipeReader` and `io.PipeWriter` from the `io` package. These types represent a pair of synchronized pipeable objects where any bytes written to `io.PipeWriter` become available to `io.PipeReader`, and `EOF` is signaled only when `io.PipeWriter` Close method is called: 

```go
func PipedStream(b *bytes.Reader) io.Reader {
  pr, pw := io.Pipe()

  go func() {
    for {
      r, err := b.ReadByte()
      if err == io.EOF {
        break
      }
      if r != byte('a') {
        pw.Write([]byte{r})
      }
    }
    pw.Close()
  }()
  return pr
}
```

It is important to note here, that `io.PipeWriter` should be run concurrently with the `io.PipeReader` due to : 

> Reading from `io.PipeReader` is only possible when bytes are written into `io.PipeWriter` or writer is closed

> Writing to `io.PipeWriter` is blocked until all data is read from `io.PipeReader` or the reader is closed 

Additionally `io.Pipe` produced `Writer/Reader` pair barely use any additional memory.

## Task #1 Analysis 

Pipe does not use any additional memory and produces Reader almost immediately. However reading from the produced reader will be blocking until data is read from the writer. However for the task purpose `Pipe` solution is definitely a winner. Let's extend the task though, to make the produced reader actually usable.

## Extending task

Now additionally to the previous task, we want to consume all the data from the `io.Reader` first by simply calling `ioutil.ReadAll` and second by making `http.Post` request with a large payload size. We will see the drastic change in performance in two different usages. 

## Benchmark ReadAll

```bash

```

## Conclusion

Careful analysis might signify that since in the second solution we are not using any additional memory it should be more performant. However it is by far not always the case. While `Pipe` solution requires far less memory it is CPU intensive due to usage of `mutex` and goroutine management. Let's do some benchmarking: 

`Buffer`:

`Pipe`:

## Links:
1. https://medium.com/stupid-gopher-tricks/streaming-data-in-go-without-buffering-3285ddd2a1e5
2. http://golang-examples.tumblr.com/post/86169510884/fastest-string-contatenation

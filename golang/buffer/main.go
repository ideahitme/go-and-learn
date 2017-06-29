package main

import (
	"bytes"
	"io"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
}

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

func BufferedStream(b *bytes.Reader) io.Reader {
	buf := bytes.NewBuffer([]byte{})
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

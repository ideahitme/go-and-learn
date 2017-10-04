package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := 100000
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i
		}
		b, _ := json.Marshal(arr)
		w.Write(b)
	}))
	defer server.Close()

	url := server.URL
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
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

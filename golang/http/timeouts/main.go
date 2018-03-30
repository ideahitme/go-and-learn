package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	testWriteTimeout()
}

// write timeout: timeout enforced between the time request is read and response is written
func testWriteTimeout() {
	// 1 second for the server to write the response and 10 seconds to read the request
	go CreateServer(1*time.Second, 10*time.Second, ":3001")
	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:3001")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
}

func CreateServer(writeTimeout, readTimeout time.Duration, addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1*time.Second + writeTimeout)
		w.Write([]byte("OK!")) // this will not be executed because of WriteTimeout
	})
	s := &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		Handler:      mux,
	}
	s.ListenAndServe()
}

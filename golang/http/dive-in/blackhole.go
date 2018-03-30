package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	go launchServer()
	// analyze client part of the http package
	time.Sleep(1 * time.Second)

	postRequest, err := http.NewRequest("POST", "http://localhost:8001", strings.NewReader("hello!"))
	if err != nil {
		panic(err)
	}
	postRequest.Trailer = map[string][]string{
		"X-Trailer-Header": []string{"foo", "bar"},
	}
	// postRequest.TransferEncoding =
	resp, err := http.DefaultClient.Do(postRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	time.Sleep(1 * time.Second)
}

func launchServer() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		fmt.Fprint(w, "hi!")
	}))
	log.Fatal(http.ListenAndServe(":8001", nil))
}

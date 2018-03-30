## HTTP in Golang

1. HTTP2 

Example of using push: 

```go

func HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
  pusher, ok := w.(http.Pusher) //check if client can speak http/2 and supports h2push
  if ok { 
    err := pusher.Push("/assets/app.js") //send resources before actual index.html
    if err != nil {
      log.Println("failed to push")
    }
  }
  fmt.Fprintf(w, indexHTML)
}) 

```

2. HTTP Tracing


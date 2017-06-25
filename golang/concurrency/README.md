## Concurrency

### Channels

**Basic rules:**

Reading from nil channel: block (deadlock)
Writing to nil channel: block (deadlock)

Reading from closed channel: never blocks, receive zero value
Writing to closed channel: panic

### Example 

1. https://ideahitme.github.io/posts/i-like-what-my-friends-like/ - see main.go for the code

### TODOs

Analyse goroutine leaks: https://medium.com/golangspec/goroutine-leak-400063aef468

### Articles

1. https://blog.golang.org/pipelines - by Sameer on different concurrency patterns (fan-in, fan-out, pipeline etc...)
## Concurrency

### Channels

**Basic rules:**

Reading from nil channel: block (deadlock)
Writing to nil channel: block (deadlock)

Reading from closed channel: never blocks, receive zero value
Writing to closed channel: panic


### Articles

1. https://blog.golang.org/pipelines - by Sameer on different concurrency patterns (fan-in, fan-out, pipeline etc...)
## Tricks with bits

1. Number of `1`s in an unsigned integer:

```go
func CountOnes(x int) int { 
    count := 0
    for x != 0 { 
        count++
        x &= (x-1)
    }
}

```

2. Reverse bits in the integer:

...

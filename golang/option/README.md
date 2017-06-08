### Extending Go struct with "unsettable" options

### Run 

```bash
go run cmd/main.go
```

### What it does? 

Package `lib.go` contains a demonstration of how to use options with closures to extend defaults of the structs with a nice API to unset the changes. Code should be self-explanatory. Run `./cmd/main.go` to see how it works.

```go
func main() {
	c := lib.NewCar()
	fmt.Printf("Riding a car with %d wheels\n", c.Wheels)
	changeWheel(c)
	fmt.Printf("Number of wheels after change: %d\n", c.Wheels)
}

func changeWheel(c *lib.Car) {
	fmt.Printf("Changing a wheel\n")
	prev := c.Option(lib.Wheels(3)) // <- generic way to extend the object
	fmt.Printf("Now have %d wheels\n", c.Wheels)
	defer prev(c) // <- unset changes
}
```
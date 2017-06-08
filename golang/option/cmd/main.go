package main

import (
	"fmt"

	lib "github.com/ideahitme/go-and-learn/golang/option"
)

func main() {
	c := lib.NewCar()
	fmt.Printf("Riding a car with %d wheels\n", c.Wheels)
	changeWheel(c)
	fmt.Printf("Number of wheels after change: %d\n", c.Wheels)
}

func changeWheel(c *lib.Car) {
	fmt.Printf("Changing a wheel\n")
	prev := c.Option(lib.Wheels(3))
	fmt.Printf("Now have %d wheels\n", c.Wheels)
	defer prev(c)
}

package main

import (
	"fmt"
	"testing"
)

func TestCircle(t *testing.T) {
	c := NewCircle()
	for i := 0; i < 3; i++ {
		n := &Node{
			address: fmt.Sprintf("org.example.node-%d", i),
		}
		c.AddNode(n)
	}
	fmt.Println(c)
}

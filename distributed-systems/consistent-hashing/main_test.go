package main

import (
	"fmt"
	"testing"
	"crypto/md5"
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

func TestMD5Hashing(t *testing.T) {
	var md5Checksum [16]byte
	md5Checksum = md5.Sum([]byte("1234"))


}


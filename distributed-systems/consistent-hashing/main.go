package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math"
	"sort"
)

func main() {

}

func defaultVirtualReplicationNumFunc(numNodes int) int {
	return int(math.Log(float64(numNodes)))
}

type Node struct {
	address string // just for example
}

func (n *Node) Store(data []byte) error {
	return nil
}

type VirtualNode struct {
	position int
	real     *Node
}

type ByPosition []*VirtualNode

func (b ByPosition) Less(i, j int) bool {
	return b[i].position < b[j].position
}
func (b ByPosition) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByPosition) Len() int {
	return len(b)
}

type CRCHasher struct{}

func (CRCHasher) Int(data []byte) int {
	return int(crc32.ChecksumIEEE(data))
}

type Hasher interface {
	Int(data []byte) int
}

type Circle struct {
	hasher                    Hasher
	virtualReplicationNumFunc func(numNodes int) int
	virtualNodes              []*VirtualNode
	numNodes                  int
}

func NewCircle() *Circle {
	return &Circle{
		hasher: CRCHasher{},
		virtualReplicationNumFunc: defaultVirtualReplicationNumFunc,
		virtualNodes:              []*VirtualNode{},
		numNodes:                  0,
	}
}

func (c *Circle) AddNode(n *Node) {
	c.numNodes++
	numVirtualNodes := c.virtualReplicationNumFunc(c.numNodes)
	for i := 0; i < numVirtualNodes; i++ {
		c.virtualNodes = append(c.virtualNodes, &VirtualNode{
			position: c.hasher.Int([]byte(fmt.Sprintf("%s-%s", n.address, string(i)))),
			real:     n,
		})
	}
}

func (c *Circle) RemoveNode(n *Node) {
	otherNodes := make([]*VirtualNode, 0)
	for _, vn := range c.virtualNodes {
		if vn.real.address != n.address {
			otherNodes = append(otherNodes, vn)
		}
	}
}

func (c *Circle) StoreData(data []byte) error {
	if len(c.virtualNodes) == 0 {
		return errors.New("No nodes available to store the data")
	}
	sort.Sort(ByPosition(c.virtualNodes))
	h := c.hasher.Int(data)
	for _, n := range c.virtualNodes {
		if h > n.position {
			continue
		}
		n.real.Store(data)
		return nil
	}
	c.virtualNodes[0].real.Store(data)
	return nil
}

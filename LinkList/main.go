package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type List struct

func main() {
	var c *Node
	b := &Node{value: 0, next: nil, prev: nil}
	c = b
	c.value = 42
	fmt.Println(c)
}

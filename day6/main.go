package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type OrbitNode struct {
	current string
	next    []*OrbitNode
}

type Orbit map[string]*OrbitNode

func (tree *Orbit) add(prev, curr string) {
	if (*tree)[curr] == nil {
		var newcurrnode = OrbitNode{curr, nil}
		(*tree)[curr] = &newcurrnode
	}
	if (*tree)[prev] == nil {
		var newprevnode = OrbitNode{prev, nil}
		(*tree)[prev] = &newprevnode
	}
	var currnode *OrbitNode = (*tree)[curr]
	(*tree)[prev].next = append((*tree)[prev].next, currnode)
}

func (node *OrbitNode) checksum(depth int) int {
	chcksum := depth
	depth += 1
	for _, child := range node.next {
		chcksum += child.checksum(depth)
	}
	return chcksum
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var tree Orbit = make(Orbit)
	s := bufio.NewScanner(f)
	for s.Scan() {
		s := strings.Split(s.Text(), ")")
		src, dst := s[0], s[1]
		tree.add(src, dst)
	}

	fmt.Println(tree["COM"].checksum(0))
}

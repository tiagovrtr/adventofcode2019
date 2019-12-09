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
	prev    *OrbitNode
	next    []*OrbitNode
}

type Orbit map[string]*OrbitNode

type Set map[string]bool

func (tree *Orbit) add(prev, curr string) {
	var prevnode *OrbitNode = (*tree)[prev]
	var currnode *OrbitNode = (*tree)[curr]
	if currnode == nil {
		var newcurrnode = OrbitNode{curr, nil, nil}
		(*tree)[curr] = &newcurrnode
		currnode = &newcurrnode
	}
	if prevnode == nil {
		var newprevnode = OrbitNode{prev, nil, nil}
		(*tree)[prev] = &newprevnode
		prevnode = &newprevnode
	}
	prevnode.next = append(prevnode.next, currnode)
	currnode.prev = prevnode
}

func (tree *Orbit) shortestPath(src, dst string) int {
	var srcnode *OrbitNode = (*tree)[src]
	var visited = Set{src: true}
	distance := 0
	return srcnode.shortestPath(dst, distance, &visited)

}

func (node *OrbitNode) shortestPath(dst string, distance int, visited *Set) int {
	(*visited)[(*node).current] = true
	if (*node).current == dst {
		return distance
	} else {
		neighb := append((*node).next, (*node).prev)
		for _, next := range neighb {
			if _, found := (*visited)[(*next).current]; !found {
				newdist := next.shortestPath(dst, distance+1, visited)
				if newdist != -1 {
					return newdist
				}
			}
		}
	}
	return -1
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
	fmt.Println(tree.shortestPath("YOU", "SAN") - 2)
}

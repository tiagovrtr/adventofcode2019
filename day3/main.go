package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Position struct {
	x, y int
}

type Trace map[Position]int

func (s1 Trace) intersect(s2 Trace) Trace {
	s_intersection := Trace{}
	for k := range s1 {
		_, found := s2[k]
		if found {
			s_intersection[k] = s1[k] + s2[k]
		}
	}
	return s_intersection
}

type Intersecter interface {
	intersect(s2 Trace) Trace
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func trace_wire(wire []string) Trace {

	var path = Trace{}

	pos := Position{0, 0}
	wire_t := 0

	for _, s := range wire {
		dir := s[0]
		steps, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}

		var vec Position
		switch dir {
		case 'U':
			vec = Position{0, 1}
		case 'D':
			vec = Position{0, -1}
		case 'L':
			vec = Position{-1, 0}
		case 'R':
			vec = Position{1, 0}
		}
		for t := 1; t <= steps; t++ {
			wire_t++
			_, found := path[Position{pos.x + t*vec.x, pos.y + t*vec.y}]
			if !found {
				path[Position{pos.x + t*vec.x, pos.y + t*vec.y}] = wire_t
			}
		}
		pos = Position{pos.x + steps*vec.x, pos.y + steps*vec.y}
	}
	return path
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

	reader := csv.NewReader(f)

	wire1, err := reader.Read()
	if err != nil {
		log.Fatal(err)
		return
	}
	wire2, err := reader.Read()
	if err != nil {
		log.Fatal(err)
		return
	}

	var pos1 Intersecter = trace_wire(wire1)
	var pos2 = trace_wire(wire2)

	intersect := pos1.intersect(pos2)
	var result int = int(^uint(0) >> 1)

	for pos := range intersect {
		result = Min(result, Abs(pos.x)+Abs(pos.y))

	}
	fmt.Println("Part 1: ", result)

	var result2 int = int(^uint(0) >> 1)
	for _, steps := range intersect {
		result2 = Min(result2, steps)

	}
	fmt.Println("Part 2: ", result2)

}

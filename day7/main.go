package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"./quickperm"
)

type Intcoder interface {
	operate(noun, verb int) int
}

type IntCode []int

func (intcode IntCode) extract_opcode(i int, numargs map[int]int) (int, []int) {
	op := intcode[i]
	opcode := op % 100
	if numargs[opcode] == 0 {
		return opcode, nil
	}
	addr := make([]int, numargs[opcode])
	op /= 100
	for j := 0; j < len(addr); j++ {
		param_mode := op % 10
		op /= 10
		if param_mode == 1 {
			addr[j] = i + j + 1
		} else {
			addr[j] = intcode[i+j+1]
		}
	}
	return opcode, addr
}

func (intcode IntCode) operate(input []int) int {
	numargs := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3, 99: 0}
	op, addr := intcode.extract_opcode(0, numargs)
	i := 0
	for op != 99 {
		next := numargs[op] + 1
		switch op {
		case 1:
			log.Printf("intcode[%v] = intcode[%v] + intcode[%v]\n", intcode[i+3], addr[0], addr[1])
			intcode[intcode[i+3]] = intcode[addr[0]] + intcode[addr[1]]
		case 2:
			log.Printf("intcode[%v] = intcode[%v] * intcode[%v]\n", intcode[i+3], addr[0], addr[1])
			intcode[intcode[i+3]] = intcode[addr[0]] * intcode[addr[1]]
		case 3:
			log.Printf("store intcode[%v]\n", intcode[i+1])
			var inputint int
			inputint, input = input[0], input[1:]
			intcode[intcode[i+1]] = inputint
		case 4:
			log.Printf("load intcode[%v]\n", intcode[i+1])
			return intcode[intcode[i+1]]
		case 5:
			log.Printf("jump intcode[%v] if intcode[%v] true \n", addr[0], addr[1])
			if intcode[addr[0]] != 0 {
				i = intcode[addr[1]]
				next = 0
			}
		case 6:
			log.Printf("jump intcode[%v] if intcode[%v] false \n", addr[0], addr[1])
			if intcode[addr[0]] == 0 {
				i = intcode[addr[1]]
				next = 0
			}
		case 7:
			log.Printf("intcode[%v] = intcode[%v] < intcode[%v]\n", intcode[i+3], addr[0], addr[1])
			var bit int
			if intcode[addr[0]] < intcode[addr[1]] {
				bit = 1
			}
			intcode[intcode[i+3]] = bit
		case 8:
			log.Printf("intcode[%v] = intcode[%v] == intcode[%v]\n", intcode[i+3], addr[0], addr[1])
			var bit int
			if intcode[addr[0]] == intcode[addr[1]] {
				bit = 1
			}
			intcode[intcode[i+3]] = bit
		}
		i += next
		op, addr = intcode.extract_opcode(i, numargs)
	}
	return intcode[0]
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

	program, err := reader.Read()
	if err != nil {
		log.Fatal(err)
		return
	}

	var intcode = IntCode{}
	for _, i := range program {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intcode = append(intcode, j)
	}

	maxresult := 0

	var a = []int{0, 1, 2, 3, 4}
	c := make(chan []int)
	go quickperm.Permutate(a, c)

	for perm := range c {
		secondinput := 0
		for _, phase := range perm {
			input := []int{phase, secondinput}
			secondinput = intcode.operate(input)
			if secondinput > maxresult {
				maxresult = secondinput
			}
		}
	}
	fmt.Println(maxresult)
}

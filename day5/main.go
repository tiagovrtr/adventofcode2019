package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (intcode IntCode) operate() int {
	numargs := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 99: 0}
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
			var input string
			fmt.Println("Insert input:")
			fmt.Scanln(&input)
			inputint, err := strconv.Atoi(input)
			if err != nil {
				panic(err)
			}
			intcode[intcode[i+1]] = inputint
		case 4:
			log.Printf("load intcode[%v]\n", intcode[i+1])
			fmt.Println(intcode[intcode[i+1]])
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

	intcode.operate()
}

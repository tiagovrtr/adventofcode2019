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

type Operation struct {
	noun, verb int
}

func (intcode IntCode) operate(noun, verb int) int {
	intcode[1] = noun
	intcode[2] = verb

	op := intcode[0]
	i := 0
	for op != 99 {
		if i > len(intcode)-3 {
			panic(fmt.Sprintf("parameters for instruction in address %v go out of bounds", i))
		}
		switch op {
		case 1:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
		case 2:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
		}
		i += 4
		op = intcode[i]
	}
	return intcode[0]
}

func test_intcode(intcode *IntCode, i, j, target int, ch chan Operation) {
	var intcodecpy = make(IntCode, len(*intcode))
	copy(intcodecpy, *intcode)
	if intcodecpy.operate(i, j) == target {
		ch <- Operation{i, j}
		close(ch)
	}
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

	ch := make(chan Operation)
	for i := 0; i < 100 && i < len(intcode); i++ {
		for j := 0; j < 100 && j < len(intcode); j++ {
			go test_intcode(&intcode, i, j, 19690720, ch)
		}
	}
	result := <-ch
	fmt.Println(100*result.noun + result.verb)
}

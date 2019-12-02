package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	var intcode = []int{}
	for _, i := range program {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intcode = append(intcode, j)
	}

	intcode[1] = 12
	intcode[2] = 2

	op := intcode[0]
	i := 0
	for op != 99 {
		switch op {
		case 1:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
		case 2:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
		}
		i += 4
		op = intcode[i]
	}

	fmt.Println(intcode[0])

}

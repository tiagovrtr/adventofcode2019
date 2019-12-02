package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calc_fuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0) - 2.0)
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

	total := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		mass, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel := calc_fuel(mass)
		extrafuel := calc_fuel(fuel)
		for extrafuel > 0 {
			fuel += extrafuel
			extrafuel = calc_fuel(extrafuel)
		}
		total += fuel
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(int(total))

}

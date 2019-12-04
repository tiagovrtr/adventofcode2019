package main

import (
	"fmt"
)

func is_password(pass int) bool {
	num := pass
	repeat := false
	digit := pass % 10
	num /= 10

	for div := 10; div/10 < num; num /= 10 {
		prevdigit := digit
		digit = num % div
		if digit > prevdigit {
			return false
		}
		repeat = repeat || digit == prevdigit
	}
	if repeat {
		return true
	} else {
		return false
	}
}

func main() {
	count := 0
	for i := 357253; i <= 892942; i++ {
		if is_password(i) {
			count++
		}
	}
	fmt.Println(count)
}

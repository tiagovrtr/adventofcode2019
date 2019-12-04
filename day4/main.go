package main

import (
	"fmt"
)

func is_password(pass int) bool {
	num := pass
	repeat := false
	repeatcount := 0
	digit := num % 10
	num /= 10

	for div := 10; num > 0; num /= 10 {
		prevdigit := digit
		digit = num % div
		if digit > prevdigit {
			return false
		}
		if digit == prevdigit {
			repeatcount++
		} else {
			repeat = repeat || (repeatcount == 1)
			repeatcount = 0
		}
	}
	if repeat || repeatcount == 1 {
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

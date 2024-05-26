package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var res string

	fmt.Scan(&number)

	for number > 0 {
		var digit = number % 10
		number = number / 10
		res += strconv.Itoa(digit * digit)
	}

	fmt.Println(res)
}

package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	for number > 0 {
		fmt.Print(number % 10)
		number = number / 10
	}
}

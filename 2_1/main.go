package main

import "fmt"

func main() {
	var str string
	var res string
	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		res += string(str[i])
		if i < len(str)-1 {
			res += "*"
		}
	}

	fmt.Println(res)
}

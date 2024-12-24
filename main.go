package main

import "fmt"

type kaan string

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	var vc kaan = "kaan"

	fmt.Println(vc, sum(1, 2))
}

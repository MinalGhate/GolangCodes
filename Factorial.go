package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	sol := factorial(num)
	fmt.Printf("Factorial is %d", sol)

}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)

}

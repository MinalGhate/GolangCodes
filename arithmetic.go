package main

import "fmt"

func main() {
	var num1, num2, sum, sub, mul int
	var div float32

	num1 = 2
	num2 = 3

	//add
	sum = num1 + num2

	//subtraction
	sub = num1 - num2

	//multiplication
	mul = num1 * num2

	//division
	div = float32(num1) / float32(num2)

	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
	fmt.Println("sum = ", sum)
	fmt.Println("subtraction = ", sub)
	fmt.Println("multilication = ", mul)
	fmt.Println("division = ", div)

}

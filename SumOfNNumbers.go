package main

import "fmt"

func sum(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + sum(num-1)
	}
}

func main() {
	var number int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &number)
	if number < 0 {
		fmt.Printf("Invalid input")
	} else {
		var result = sum(number)
		fmt.Println("Sum: ", result)
	}
}

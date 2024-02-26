package main

import (
	"fmt"
)

// func main() {
// 	//******************Function
// 	greet()
// }
// func greet() {
// 	fmt.Println("Good morning")
// }

// Function to add 2 numbers
// func add(n1 int, n2 int) int {
// 	// n1 := 5
// 	// n2 := 3

// 	//sum := n1 + n2

// 	//fmt.Printf("Addition of %d and %d is %d", n1, n2, sum)
// 	//return sum
// 	return n1 + n2

// 	// fmt.Println("code after return statment") //return statment must be the last statment unless it is a recurssion function
// 	// return 4 //after this will not get error but this value will never return

// }
// func main() {
// 	var n1 int
// 	var n2 int
// 	fmt.Scanln(&n1, &n2)
// 	//add(n1, n2)
// 	//sol=add(n1,n2)
// 	//fmt.Printf("Addition of %d and %d is %d", n1, n2, sol)
// 	fmt.Printf("Addition of %d and %d is %d", n1, n2, add(n1, n2))

// }

// func main() {
// 	var n1 int
// 	var n2 int

// 	fmt.Scanln(&n1, &n2)

// 	sum, diff := calculate(n1, n2)
// 	fmt.Printf("Sum and Difference of %d and %d is %d and %d respectively", n1, n2, sum, diff)
// }

// func calculate(n1 int, n2 int) (int, int) {
// 	return n1 + n2, n1 - n2
// }

// *****code reusability
// func main() {
// 	getSquare(3)
// 	getSquare(4)
// 	getSquare(5)
// }

// func getSquare(num int) {
// 	fmt.Printf("Square of %d is %d\n", num, num*num)
// }

// ***********go variable scope
// ***local variable
// - define inside function and can be used inside that specidic functions only cannor be used in anoreher function
// ***global variable
// -define above main funtion(outside of all function) and can be used in all function
// compilation is top to bottom and execution from main function

// var sum int = 2

// func main() {
// 	var sum int = 0
// 	addNum()
// 	fmt.Printf("Sum is %d ", sum)
// }
// func addNum() {
// 	//var sum int
// 	sum = 4 + 5
// }

// ***************Recursion
func main() {
	fmt.Println("*****COUNTDOWN STARTS*****")
	countDown(5)

}

func countDown(number int) {
	// if number == 0 {
	// 	fmt.Println("HAPPY NEW YEAR")
	// 	return
	// }
	if number > 0 {
		fmt.Println(number)
		countDown(number - 1)
	} else {
		fmt.Println("Happy New Year")
	}
}

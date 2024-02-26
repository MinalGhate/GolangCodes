//***************Anonymous Function

package main

import "fmt"

// var global = func() {
// 	fmt.Printf("This is anonymous function in global")
// 	//greet() //Error

// }

// func main() {
// 	var greet = func() {
// 		fmt.Println("HEllo, inside anonymous function")
// 		global()
// 	}
// 	var welcome = greet

// 	greet()
// 	global()
// 	fmt.Printf("\n%T", greet)
// 	fmt.Printf("\n%T", welcome)
// 	fmt.Println()
// 	fmt.Println(welcome)
// }

//********Anonymous function with parameter

// func main() {
// 	var sum = func(n1, n2 int) int { //sum datatype func(int,int) int
// 		sum := n1 + n2
// 		return sum
// 	}

// 	result := sum(5, 3) //datatype int
// 	fmt.Println("Sum is: ", result)

// }

//************Anonymous function as argument to other function

// func main() {
// 	var sum = func(num1 int, num2 int) int {
// 		return num1 + num2
// 	}

// 	result := findSquare(sum(6, 4))
// 	fmt.Println("Result is: ", result)
// }

// func findSquare(num int) int {
// 	square := num * num
// 	return square
// }

//add mult sub div using anonymous

func main() {
	var add = func(num1 int, num2 int) int {
		return num1 + num2
	}

	var sub = func(num1 int, num2 int) int {
		return num1 - num2
	}

	var mult = func(num1 int, num2 int) int {
		return num1 * num2
	}

	var div = func(num1 int, num2 int) int {
		return num1 / num2
	}

	var num1, num2 int
	fmt.Println("Enter num1 and num2")
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println("Addition = ", add(num1, num2))
	fmt.Println("Subtraction = ", sub(num1, num2))
	fmt.Println("Multiplication = ", mult(num1, num2))
	fmt.Println("Division = ", div(num1, num2))

}

package main

import (
	"fmt"
)

// var num int
// var i int = 0

// func main() {
// 	var n int
// 	fmt.Print("Enter number:")
// 	fmt.Scanf("%d", &n)
// 	num = n
// 	printFibonacci(n)

// }

// func printFibonacci(n int) {
// 	if i > num {
// 		return
// 	}
// 	fmt.Printf("%d ", fibonacci(i, 0))
// 	i++
// 	printFibonacci(i)

// }

// func fibonacci(n int, f int) int {
// 	if n == 0 || n == 1 {
// 		f += n
// 	} else {
// 		f = fibonacci(n-1, f) + fibonacci(n-2, f)
// 	}
// 	return f
// }

// ************Using anonymous
func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(3))
}

//write a program to print 1 to n using anonymous

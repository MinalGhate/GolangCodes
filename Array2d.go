package main

import "fmt"

func main() {
	var matrix [3][3]int

	fmt.Printf("Enter matrinx elemnt: \n")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Element: matrix[%d][%d]: ", i, j)
			fmt.Scanln(&matrix[i][j])
		}
	}

	fmt.Printf("Matrix: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println("\n")
	}

	fmt.Printf("Left diagonal of matrix: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				fmt.Printf("%d ", matrix[i][j])
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

//Print the sum of left diagnonal elemt
//print right diagonal elemnt
//Print the sum of right diagnonal elemt
//print addition of 2 matrix
//calculate power of given number using go to, power and number both take input
//reverse given number using go to
//check number is pallindrome using goto and for loop
//check armstrong of given number using goto and loop
//chaeck prime number using goto and loop
//check perfect number using go to
//linear search and binary search
//bubble sort ascending and descending
//insertion sort
//selection sort

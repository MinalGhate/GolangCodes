package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	var numbers = make([]int, 3, 5) //slice of type int with length 3 and capacity 5, capacity is how much it can actually accomodate
// 	printSlice(numbers)
// 	numbers[0] = 10
// 	printSlice(numbers)
// 	//numbers[3] = 5 //******Error out of bound (so we have to append it)
// 	printSlice(numbers)

// }

// func printSlice(x []int) {
// 	fmt.Printf("len = %d  cap=%d  slice=%v\n", len(x), cap(x), x)
// }

//*************copy
// func main() {
// 	primeNum := []int{2, 3, 5, 7}
// 	num := []int{1, 2, 3}
// 	fmt.Println("primeNum ", primeNum)
// 	fmt.Println("Numbers: ", num)

// 	copy(num, primeNum) //only 2,3,5 are copied, as the size is 3

// 	fmt.Println("Numbers (after copy): ", num)
// }

//*********len
// func main() {
// 	nums := []int{2, 4, 6, 8, 10}
// 	for i := 0; i < len(nums); i++ {
// 		fmt.Println(nums[i])
// 	}
// }

//*************sort

func main() {
	slice := []string{"honesty", "is", "the", "best", "policy"}

	sort.Strings(slice)

	fmt.Println("Sorted slice: ")
	for _, item := range slice {
		fmt.Printf("%s ", item)
	}
}

//write a function to check if the slice is sorted or not

package main

import "fmt"

func main() {
	var thisMonth int
	fmt.Println("Enter month: ")
	fmt.Scan(&thisMonth)

	switch thisMonth {
	case 1, 2, 3:
		fmt.Println("Quarter 1")
	case 4, 5, 6:
		fmt.Println("Quarter 2")
	case 7, 8, 9:
		fmt.Println("Quarter 3")
	case 10, 11, 12:
		fmt.Println("Quarter 4")
	default:
		fmt.Println("invalid month")
	}

}

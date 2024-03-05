package main

import "fmt"

func main() {
	var year int = 0

	fmt.Printf("Enter year: ")
	fmt.Scanf("%v", &year)
	if year >= 0 {
		if year%4 == 0 {
			fmt.Printf("Year %v is leap year", year)
		} else {
			fmt.Printf("Year %v is not a lear year", year)
		}
	} else {
		fmt.Printf("Enter valid year ")
	}

}

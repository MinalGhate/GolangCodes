//Fallthrough will not chack rest of swwitch caser and execute all case statment

package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter number: ")
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		//fallthrough  //will get out of switch if fallthrough not there
	case 5:
		fmt.Println("5")
		fallthrough
	default:
		fmt.Println("No specified case ")
	}
}

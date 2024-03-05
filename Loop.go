package main

import "fmt"

// func main() {
// 	number := 1

// 	//***************(Do-while)
// 	//loop that runs infinitely
// 	for {
// 		//condition to terminate the loop
// 		fmt.Printf("%d\n", number)
// 		number++
// 		if number > 5 {
// 			break
// 		}
// 	}
// }

//********While loop (To print multiplication table)

// func main() {
// 	multiplier := 1

// 	//run while loop for 10 times
// 	for multiplier <= 10 {
// 		//find the product
// 		product := 5 * multiplier

// 		//print the multiplication table in format 5 * 1 = 5
// 		fmt.Printf("5 * %d = %d\n", multiplier, product)
// 		multiplier++

// 	}
// }

//******continue
//******break

//********Range

func main() {
	colors := []string{"Red", "Yellow", "Green"} //type Slice , Array without specified size = type Slice
	// for index, val := range colors {
	// 	fmt.Println(index, "-", val)
	// }

	for _, val := range colors {
		fmt.Println(val)
	}

}

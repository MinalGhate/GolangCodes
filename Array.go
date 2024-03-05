package main

import "fmt"

func main() {

	var arrayOfInteger = [5]int{1, 5, 8, 0, 3} //Array

	var arrayWithoutSize = []int{1, 5, 8, 0, 3} //Slice

	fmt.Println(arrayOfInteger)

	fmt.Println(arrayOfInteger[:2])
	fmt.Println(arrayOfInteger[2:])
	//fmt.Println(arrayOfInteger[-1:]) //Error***********

	fmt.Println(arrayWithoutSize)
	fmt.Println(arrayWithoutSize[:4])
	fmt.Println(arrayWithoutSize[1:2])
	//fmt.Println(arrayWithoutSize[-1:]) //Error************

	// fmt.Printf("%T\n", arrayOfInteger)
	// fmt.Println(arrayOfInteger)

	// fmt.Printf("%T\n", arrayWithoutSize)
}

// func main() {
// 	//declare array cariable of type string
// 	//undefined size

// 	var arrayOfString = [...]string{"Hello", "Rajesh"}
// 	fmt.Println(arrayOfString)
// 	fmt.Printf("%T", arrayOfString)
// }

//*********len

package main

import "fmt"

// func main() {

// 	var arrayOfInteger = [5]int{1, 5, 8, 0, 3} //Array

// 	var arrayWithoutSize = []int{1, 5, 8, 0, 3} //Slice

// 	fmt.Println(arrayOfInteger)

// 	fmt.Println(arrayOfInteger[:2])
// 	fmt.Println(arrayOfInteger[2:])
// 	//fmt.Println(arrayOfInteger[-1:]) //Error***********

// 	fmt.Println(arrayWithoutSize)
// 	fmt.Println(arrayWithoutSize[:4])
// 	fmt.Println(arrayWithoutSize[1:2])
// 	//fmt.Println(arrayWithoutSize[-1:]) //Error************

// 	// fmt.Printf("%T\n", arrayOfInteger)
// 	// fmt.Println(arrayOfInteger)

// 	// fmt.Printf("%T\n", arrayWithoutSize)
// }

// func main() {
// 	//declare array cariable of type string
// 	//undefined size

// 	var arrayOfString = [...]string{"Hello", "Rajesh"}
// 	fmt.Println(arrayOfString)
// 	fmt.Printf("%T", arrayOfString)
// }

//**********range in for loop
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	intSlice := arr[2:5]

	fmt.Println("Slice element: ")

	for index, ele := range intSlice {
		fmt.Printf("Index = %d, element = %d\n", index, ele)
	}
}

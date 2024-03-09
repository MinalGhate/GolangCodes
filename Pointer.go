package main

import "fmt"

func main() {
	var intData int
	var intPointer *int

	intPointer = &intData

	//&intData and intPointer will store address of intData
	//*intPointer and *(&intData) will store value of intData

	fmt.Println("What intData stores", intData)
	fmt.Println("adress of intData", &intData)
	fmt.Println("what intPointer stores: ", intPointer)

	*intPointer = 30

	fmt.Println("what intData now stores", intData)
}

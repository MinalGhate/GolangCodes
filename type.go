package main

import "fmt"

func main() {
	var a uint8
	//a = 300 //error : cannot use 300 (untyped int constant) as uint8 value in assignment (overflows)
	a = 255
	fmt.Printf("a=%d", a)
}

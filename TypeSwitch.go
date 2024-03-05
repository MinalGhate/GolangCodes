package main

import "fmt"

func main() {
	//var x interface{} = "hello"
	var x interface{} = func(int) float64 {
		return 6.0
	}
	switch i := x.(type) {
	case nil:
		fmt.Println("Type of x: %T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int) float64")
	case bool, string:
		fmt.Printf("x is bool or strinf")
	default:
		fmt.Printf("don't know the type")
	}

}

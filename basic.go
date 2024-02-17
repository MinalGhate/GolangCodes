package main

import "fmt"

func main() {

	const (
		g    int    = 10
		name        = 'g'
		n    string = "it is constant"
		grav        = 9.8
	)

	fmt.Printf("%T ", name) //it will not give char as there is no char datatye in go
	fmt.Println(n)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", grav) //%T to get type of variable
	//printf f means format specifier

}

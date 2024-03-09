package main

import "fmt"

//***********without using pointer

// func Swap(x int, y int) (int, int) {
// 	return y, x
// }
// func main() {
// 	var a int = 10
// 	var b int = 20
// 	fmt.Println(a, b)
// 	a, b = Swap(a, b)

// 	fmt.Println(a, b)
// }

//***********with using pointer

func Swap(x *int, y *int) {
	t := *x
	*x = *y
	*y = t

}

func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	Swap(&a, &b)
	fmt.Println(a, b)
}

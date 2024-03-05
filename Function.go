// ********Assigning a function to a variable
// package main

// import (
// 	"fmt"
// )

// func test(x int) {
// 	fmt.Println("Hello", x)

// }

// func main() {
// 	x := test
// 	x(5)
// }

//****************

//package main

// import "fmt"

// // var test = func(x int) {
// // 	fmt.Println(x)
// // }

// // test:=func(x int){   ///Error******
// // 	fmt.Println(x)
// // }

// func main() {

// 	test:=func(x int){
// 		fmt.Println(x)
// 	}
// 	test(5)
// }

//***********

// package main

// import "fmt"

// //a:=10 //Error: infert varianle cannot be defined globally
// var a = 10

// func main() {
// 	fmt.Println(a, a)
// }

// package main

// import "fmt"

// func main() {
// 	test := func(x int) int {
// 		return x * -1
// 	}(8)

// 	fmt.Println(test)
// }

// ******Multiple value return in function
package main

import "fmt"

func calculate(x int, y int) (int, int) {
	return x + y, x - y
}

func main() {
	sum, diff := calculate(10, 20)
	fmt.Printf("Addition is %d  and subtraction is %d", sum, diff)
}

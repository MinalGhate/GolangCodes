package main

import (
	"fmt"
)

func main() {
	//*****************average of two integer number
	// var num1 int
	// var num2 int
	// //var avg float64
	// num1 = 10
	// num2 = 5
	// //avg = (float64(num1) + float64(num2)) / 2
	// //fmt.Printf("Average of %d and %d is %.2f\n", num1, num2, avg)

	// //fmt.Printf("Average of %d and %d is %.2f\n", num1, num2, (num1+num2)/2) //error
	// fmt.Printf("Average of %d and %d is %d\n", num1, num2, (num1+num2)/2)

	//*********************Decision making/If statement

	// if 20 > 18 {
	// 	fmt.Println("20 is greater than 18")
	// }

	// if true {
	// 	fmt.Println("It's true")
	// }
	// fmt.Println("it's false")

	// if false {
	// 	fmt.Println("It's true")
	// }
	// fmt.Println("it's false")

	// var a = 5
	// var b = 4
	// var c = 9

	// if !((a + b) == c) {
	// 	fmt.Println("a+b is equal to c")
	// }
	// fmt.Println("statment is false")

	// a:=1
	// b:=0

	// if (a&&b) { //error

	// }

	// var a bool = true
	// var b bool = true

	// if a && b {
	// 	fmt.Println("Its true")
	// }

	// x := 8
	// y := 4
	// if y := 10; x < y {  //inside if we can define at most one statment
	// 	fmt.Println("x is less than y")
	// 	fmt.Printf("inside if x = %d and y = %d\n", x, y)
	// }
	// fmt.Printf("outside if x = %d and y = %d", x, y)

	// var name string
	// var age int
	// fmt.Println("Enter name and age: ")
	// // if _, err := fmt.Scan(&name, &age); err != nil {
	// // 	fmt.Println(err)
	// // 	os.Exit(1) //1=>normal termination
	// // }

	// if v, err := fmt.Scan(&name, &age); err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(v)
	// 	os.Exit(1)

	// }

	// fmt.Printf("Your name is %s\n", name)
	// fmt.Printf("Your age is %d\n", age)

	//**********program to bigger number in three integers

	// var num1 int
	// var num2 int
	// var num3 int

	// fmt.Printf("Enter three integers: \n")
	// //fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	// fmt.Scanln(&num1, &num2, &num3)
	// // fmt.Scanln(&num2)
	// // fmt.Scanln(&num3)

	// // if num1 > num2 {
	// // 	if num3 > num1 {
	// // 		fmt.Printf("num3 = %d is bigger", num3)
	// // 		os.Exit(1)
	// // 	}
	// // 	fmt.Printf("num1 = %d is bigger", num1)
	// // 	os.Exit(1)
	// // }

	// if num1 > num2 && num1 > num3 {
	// 	fmt.Printf("num1 = %d is largest \n", num1)
	// }
	// if num2 > num3 && num2 > num1 {
	// 	fmt.Printf("num2 = %d is largest \n", num2)
	// }
	// if num3 > num1 && num3 > num2 {
	// 	fmt.Printf("num3 = %d is largest \n", num3)
	// }

	//******************if else
	// if condition{
	// 	//when condition true
	// } else{
	// 	//when conditon false
	// }

	// x := 3
	// y := 10

	// if x < y {
	// 	fmt.Println("x is less than y")
	// } else {
	// 	fmt.Println("x is greater than y")
	// }

	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	// if num%2 == 0 {
	// 	fmt.Printf("%d is even", num)
	// } else {
	// 	fmt.Printf("%d is odd", num)
	// }

	var abs int
	if num < 0 {
		abs = -num
		fmt.Printf("absolute value of %d is %d", num, abs)
	} else {
		abs = num
		fmt.Printf("absolute value of %d is %d", num, abs)
	}

}

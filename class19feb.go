package main

import (
	"fmt"
)

func main() {

	//*********************************OPERATORS*********************************************************
	//***********1.Arithmetic operator
	//+, -, *, /, %, ++, --
	// a := 9
	// a += 5
	// fmt.Println(a)

	// var num int
	// println("Enter num: ")
	// fmt.Scanf("%d", &num)
	// mod := num % 2
	// fmt.Printf("%d modulo 2 is: %d", num, mod)

	// a := 3
	// b := 5
	// c := a & b //and
	// fmt.Printf("a & b = %d\n", c)
	// d := a | b //or
	// fmt.Printf("a | b = %d\n", d)
	// e := a ^ b //xor
	// fmt.Println("a ^ b = ", e)

	//***************2.Comparision/relational operator
	//a>2 //comparision operator
	//a>>2 //bitwise operatior //right shift //value become small
	//a<<2 //left shift //value become bigger

	// a := 10
	// fmt.Println(a > 5)
	// fmt.Println(a == 10)
	// fmt.Println(a < 5)

	//*******************3.Logical operator
	//&&, ||, !

	// a := true
	// b := false
	// fmt.Println(a && b)
	// fmt.Println(a || b)
	// fmt.Println(!a && b)

	// var num1 int = 5
	// var num2 byte = 2
	// var str string = "Hello"
	// var f float32
	// var f2 float64
	// b := true

	// fmt.Printf("Size of num1: %d\n", unsafe.Sizeof(num1))
	// fmt.Printf("Size of num2: %d\n", unsafe.Sizeof(num2))
	// fmt.Printf("Size of string: %d\n", unsafe.Sizeof(str))
	// fmt.Printf("Size of f: %d\n", unsafe.Sizeof(f))
	// fmt.Printf("Size of f2: %d\n", unsafe.Sizeof(f2)) //why?
	// fmt.Printf("Size of b: %d\n", unsafe.Sizeof(b))

	// fmt.Println()

	// fmt.Println("Type of num1: ", reflect.TypeOf(num1))
	// fmt.Println("Type of num2: ", reflect.TypeOf(num2))
	// fmt.Println("Type of str: ", reflect.TypeOf(str))
	// fmt.Println("Type of f: ", reflect.TypeOf(f))
	// fmt.Println("Type of f2: ", reflect.TypeOf(f2))
	// fmt.Println("Type of b: ", reflect.TypeOf(b))

	// fmt.Println()

	// fmt.Println("Type of num1: ", reflect.ValueOf(num1).Kind())
	// fmt.Println("Type of num2: ", reflect.ValueOf(num2).Kind())
	// fmt.Println("Type of str: ", reflect.ValueOf(str).Kind())
	// fmt.Println("Type of f: ", reflect.ValueOf(f).Kind())
	// fmt.Println("Type of f2: ", reflect.ValueOf(f2).Kind())
	// fmt.Println("Type of b: ", reflect.ValueOf(b).Kind())

	// //**********defer
	// //execution from down to  up
	// defer fmt.Println("Hello")
	// defer fmt.Println("Hii")
	// fmt.Println("Good morning")
	// //defer a := 5 //error

	// var radius float32
	// var area float32
	// var perimeter float32
	// const pi = 3.14

	// fmt.Printf("Enter radius: ")
	// fmt.Scanf("%f", &radius)

	// //area = math.Pi * radius * radius
	// area = pi * radius * radius
	// perimeter = 2 * pi * radius

	// fmt.Printf("Area = %.2f\n", area)
	// fmt.Printf("Perimeter = %.2f\n", perimeter)

	var ftemp float32
	var ctemp float32

	// fmt.Printf("Enter temp in farhenite: ")
	// fmt.Scanf("%f", &ftemp)
	// ctemp = (ftemp - 32) / 1.8
	// fmt.Printf("Temp in celcius: %.2f", ctemp)

	fmt.Printf("Enter temp in celcius: ")
	fmt.Scanf("%f", &ctemp)
	ftemp = (ctemp * 1.8) + 32
	fmt.Printf("Temp in fehranite: %.2f", ftemp)

}

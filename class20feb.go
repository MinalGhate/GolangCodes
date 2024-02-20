package main

import (
	"fmt"
)

func main() {
	//********sprintf, writestring
	// var dd int = 20
	// var mm int = 2
	// var yy int = 2024
	// var str string
	// str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
	// io.WriteString(os.Stdout, str)

	//*************scan
	// var str string
	// var i = 4
	// var f float32
	// var b bool
	// fmt.Printf("Enter string: ")
	// fmt.Scan(&str, &i, &f, &b) //no format specifier
	// fmt.Printf("Result: %s %d %g %t\n", str, i, f, b)

	// var a int = -123.0
	// var b uint = 0
	//b = a //Error: as the number are out of range it cannot assign it, go doesn't use autotypecasting, range doesn't matter type should be same
	//a = b //error
	//math.Abs(a)
	//b = uint(a)
	//a = int(b)
	// a = int(b)
	// fmt.Printf("a = %d, b = %d\n", a, b)

	//**************sqrt
	// var num = 25
	// var r float64
	// r = math.Sqrt(float64(num))
	// fmt.Println(r)
	// println(int(r))

	// var num = 25
	// var r float32
	// r = float32(math.Sqrt(float64(num)))
	// fmt.Println(r)

	//***********Explicit type conversion
	var x uint = 4294967294
	var y float64 = float64(x)
	var z int = int(x)
	fmt.Printf("Value of x is %d and type %T\n", x, x)
	fmt.Printf("Value of y is %.2f and type %T\n", y, y)
	fmt.Printf("Value of z is %d and type %T\n", z, z)

}

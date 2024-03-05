package main

import (
	"fmt"
)

//************Time equals
// func main() {

// 	date1 := time.Date(2024, 3, 10, 5, 30, 8, 0, time.UTC)
// 	date2 := time.Date(2024, 3, 10, 5, 30, 8, 0, time.UTC)
// 	date3 := time.Date(2024, 3, 11, 5, 30, 8, 0, time.UTC)

// 	if date1.Equal(date2) {
// 		fmt.Println("date1 and date2 are equal")
// 	} else {
// 		fmt.Println("date1 and date2 are not equal")
// 	}

// 	if date1.Equal(date3) {
// 		fmt.Println("date1 and date3 are equal")
// 	} else {
// 		fmt.Println("date1 and date3 are not equal")
// 	}

// }

// ***************Do while loop
// func main() {
// 	num := 1

// 	for {
// 		fmt.Printf("%d\n", num)
// 		num++
// 		if num > 5 {
// 			break
// 		}
// 	}
// }

//**********ASCII value of character

// func main() {
// 	str := "Hello World"

// 	fmt.Println("ASCII value of strings: ")

// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("%c %d\n", str[i], str[i])
// 	}
// }

//***********While loop

// func main() {
// 	num := 1
// 	for num <= 5 {
// 		fmt.Println(num)
// 		num++
// 	}
// }

//***********Array

// func main() {
// 	var arrayOfInteger = [5]int{1, 9, 3, 8, 4}
// 	fmt.Println(arrayOfInteger)
// }

// **************for loop
// func main() {
// 	k := 1

// 	for ; k <= 5; k++ {
// 		fmt.Println(k)
// 	}

// 	k = 1
// 	for k <= 5 {
// 		fmt.Println(k)
// 		k++
// 	}

// 	for k := 1; ; k++ {
// 		fmt.Println(k)
// 		if k == 5 {
// 			break
// 		}
// 	}
// }

//*********Multiplication table

// func main() {
// 	var multiple int
// 	fmt.Print("Enter multiplier: ")
// 	fmt.Scanln(&multiple)

// 	multiplier := 1

// 	for multiplier <= 10 {
// 		product := multiple * multiplier
// 		fmt.Printf("%d x %d = %d\n", multiple, multiplier, product)
// 		multiplier++
// 	}

// }

//*******Other way of declaring array

// func main() {
// 	var arrayOfString = [...]string{"Hello", "World", "Wassup"}

// 	fmt.Println(arrayOfString)
// }

//*********Changing array elemnt

// func main() {
// 	names := [3]string{"Minal", "Prachi", "John"}
// 	fmt.Println(names)
// 	names[2] = "Selena"
// 	fmt.Println(names)
// }

//*******Initializong specific elements of an array

// func main() {
// 	arrayOfInteger := [5]int{0: 7, 3: 9}
// 	fmt.Println(arrayOfInteger)
// }

// ***********initializing an array in go
func main() {
	var arrayOfInteger [3]int
	arrayOfInteger[0] = 8
	arrayOfInteger[2] = 1

	fmt.Println(arrayOfInteger)
}

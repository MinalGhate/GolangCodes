package main

import "fmt"

//*********Print odd number using golang closure

// func calculate() func() int {
// 	num := 1

// 	return func() int {
// 		num += 2
// 		return num
// 	}
// }

// func main() {
// 	odd := calculate()

// 	fmt.Println(odd())
// 	fmt.Println(odd())
// 	fmt.Println(odd())
// 	fmt.Println(odd())

// 	fmt.Println()

// 	odd2 := calculate()
// 	fmt.Println(odd2())
// 	fmt.Println(odd2())
// }

//*********
// func greet() func(string) string {
// 	name := "Minal"

// 	return func(str string) string {
// 		name = "Hi " + name + str
// 		return name
// 	}
// }

// func main() {
// 	message := greet()

// 	fmt.Println(message("Ghate"))
// }

//*********Closure helps in data isolation

func displayNumber() func() int {
	fmt.Println("Numbers")
	num := 0

	return func() int {
		fmt.Println("Next")
		num++
		return num
	}
}

func main() {
	num1 := displayNumber()

	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())

	fmt.Println()

	num2 := displayNumber()
	fmt.Println(num2())
	fmt.Println(num2())

	fmt.Println()

	fmt.Println(num1())
	fmt.Println()

	fmt.Println(displayNumber())
	fmt.Println()

	fmt.Println(displayNumber()())
	fmt.Println(displayNumber()())

}

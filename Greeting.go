//time less than 10 =>good morn
//<20 =>good day
//exceeding that => good evening

package main

import "fmt"

func main() {
	var time int

	fmt.Println("Enter time: ")
	fmt.Scan(&time)

	if time <= 24 && time >= 0 {
		if time <= 10 {
			fmt.Println("Good Morning")
		} else if time <= 20 {
			fmt.Println("Good day")
		} else {
			fmt.Println("Good Evening")
		}
	} else {
		fmt.Println("Enter valid time")
	}
}

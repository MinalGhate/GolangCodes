package main

import (
	"fmt"
	"time"
)

// func main() {
// 	today := time.Now()
// 	fmt.Println(today)
// 	fmt.Println(today.Day())
// 	switch today.Day() {
// 	case 5:
// 		fmt.Println("Today is 5th, clean home")
// 	case 10:
// 		fmt.Println("Today is 10th, buy wine")
// 	case 25:
// 		fmt.Println("Today is 25th, buy food")
// 	case 31:
// 		fmt.Println("Party")
// 	default:
// 		fmt.Println("No info available")
// 	}
// }

//*****************

// func main() {
// 	year, month, date := time.Now().Date()
// 	fmt.Printf("%d/%d/%d \n", date, month, year)
// 	fmt.Printf("%d/%s/%d \n", date, month, year)
// 	// fmt.Printf("%s/%s/%s \n", date, month, year) //will ask for %d
// 	fmt.Println(date, month, year)

// }

func main() {
	currentDateTime := time.Now()

	day := currentDateTime.Day()
	month := currentDateTime.Month()
	year := currentDateTime.Year()

	hour := currentDateTime.Hour()
	min := currentDateTime.Minute()
	sec := currentDateTime.Second()

	fmt.Printf("Report printed on %d/%d/%d at %d:%d:%d", day, month, year, hour, min, sec)
}

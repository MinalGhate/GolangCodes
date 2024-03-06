package main

import "fmt"

// func main() {
// 	subjectMatks := map[string]float32{"Golang": 85, "java": 80, "python": 81}

// 	fmt.Println(subjectMatks)
// }

// func main() {
// 	flowerColor := map[string]string{"Sunflower": "yellow", "Rose": "Red", "jasmin": "white"}

// 	fmt.Println(flowerColor["Sunflower"])

// 	fmt.Println(flowerColor["red"])  //will leave blank space
// 	fmt.Println(flowerColor["rose"]) //will leave black space

// 	fmt.Println(flowerColor["Rose"])
// }

//*********chainging value of particular key in map

// func main() {
// 	capital := map[string]string{"Nepal": "Katmandu", "US": "NewYork"}

// 	fmt.Println("Initial Map: ", capital)

// 	capital["US"] = "Washington DC"

// 	fmt.Println("Updated map: ", capital)
// }

//************Add delete elemetn to map
// func main() {
// 	students := map[int]string{2: "John", 1: "Lily"}

// 	fmt.Println("Initial map: ", students)

// 	students[5] = "Robin"

// 	students[3] = "Julie" //it add on the sorting basis of key

// 	fmt.Println("Updated map : ", students)

// 	delete(students, 1)
// 	fmt.Println("Map after delete: ", students)

// }

//**********looping in map

func main() {
	num := map[int]int{3: 9, 2: 4, 4: 16, 5: 25, 6: 26}
	fmt.Println(num)

	for key, value := range num {
		fmt.Printf("Key: %d and value: %d\n", key, value)
	}
	fmt.Println()
	for _, value := range num {
		fmt.Printf("value: %d\n", value)
	}
	fmt.Println()
	for key, _ := range num {
		fmt.Printf("key: %d\n", key)
	}
}

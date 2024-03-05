package main

// import (
// 	"fmt"
// )

// func main() {
// 	arrayOfInteger := [...]int{4, 2, 7, 1}
// 	temp := 0
// 	for i := 0; i < len(arrayOfInteger); i++ {
// 		for j := i; j < len(arrayOfInteger); j++ {
// 			if arrayOfInteger[j] < arrayOfInteger[i] {
// 				temp = arrayOfInteger[i]
// 				arrayOfInteger[i] = arrayOfInteger[j]
// 				arrayOfInteger[j] = temp
// 			}
// 		}
// 	}

// 	fmt.Println(arrayOfInteger)
// }

//**********************************************************

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{4, 2, 7, 1, 9, 3}
	sort.Ints(num) //this function only accept argument slice not array
	fmt.Println("Sorted: ", num)
	nums := []float64{4.4, 2.0, 7.0, 7, 9, 3}
	sort.Float64s(nums)
	fmt.Println(nums)
	names := []string{"priyanshi", "minal", "prachi"}
	sort.Strings(names)
	fmt.Println(names)

}

//h.w: by default sorting function sort in ascending, do it in descending
//h.w: sort to sort with argument array

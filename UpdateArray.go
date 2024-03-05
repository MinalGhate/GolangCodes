package main

import "fmt"

//**************Inserting element by shifting other elements

func main() {
	//Orignal array
	orignalArray := []int{1, 2, 3, 4, 5}
	//item to insert
	newItem := 6
	//index where the new item should be inserted
	index := 2
	//insert the item into the array
	newArray := insertIntoArray(orignalArray, newItem, index)
	//print the new array
	fmt.Println("Orignal array:", orignalArray)
	fmt.Println("item to insert: ", index)
	fmt.Println("New array: ", newArray)

}

//Funtion to insert an item into an specified index
func insertIntoArray(arr []int, item int, index int) []int {
	//create a new array with increased length
	newArr := make([]int, len(arr)+1)

	//copy element before the insertion point
	copy(newArr[:index], arr[:index])

	//insert the new item
	newArr[index] = item

	//copy elemetn after the insertion point
	copy(newArr[index+1:], arr[index:])
	return newArr
}

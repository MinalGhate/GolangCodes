package main

//strings => package
//string=> datatype
import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello World"
	var subStr string = "World"

	//***Contains
	// if strings.Contains(str, subStr) {
	// 	fmt.Printf("String (%s) contains sub-string (%s)", str, subStr)
	// } else {
	// 	fmt.Printf("String (%s) does not contains sub-string (%s)", str, subStr)
	// }

	//******upper and lower case
	// fmt.Println("converting str to uppercase: ", strings.ToUpper(str))
	// fmt.Println("converting str to lowercase: ", strings.ToLower(str))

	//******index

	fmt.Println("Index is: ", strings.Index(str, "W"))
	fmt.Println("Index is: ", strings.Index(str, "orld"))
	fmt.Println("Index is: ", strings.Index(str, "w"))
	ind := strings.Index(str, subStr)
	fmt.Println("Index is: ", ind)

}

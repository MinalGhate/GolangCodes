package main
import "fmt"

type Rectangle func(int, int) int  // Rectangle is a type of anonymous function which take 2 int argument and int as return statment
//ex: type integer int
//var a int     var a integer //this both statment are similar


type rectanglePara struct { //rectanglePara is structure type
	length int
	bredth int
	color string //length, bredth, color all are member of rectanglePara
	rect Rectangle //rect is member of rectanglePara of type Rectangle function whose argument is anonymous function and return statement is  int
}

func main(){
	//assign values to struct variables
	result := rectanglePara{ //result is a structure variable of type rectanglePara
		length: 10,
		bredth : 20,
		color: "Red"
		rect: 
	}
}


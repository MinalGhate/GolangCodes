package main
import "fmt"

func main(){
	var arr[6] int
	var item int = 0 
	var index int = 0
	fmt.Printf("Enter array elements: \n")

	for i:=0; i<=5; i++ {
		fmt.Printf("Elemts: arr[%d]: ",i)
		fmt.Scanf("%d",&arr[i])
	}

	index = -1

	for i:=0; i<5; i++{
		for j:=i+1; j<=5; j++{
			if arr[i]==arr[j]
		}
	}
}
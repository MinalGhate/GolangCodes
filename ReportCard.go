package main

import (
	"fmt"
)

func main() {
	var name string
	var roll int
	var sub1 string
	var sub2 string
	var sub3 string
	var m1 float32
	var m2 float32
	var m3 float32
	//var total float32

	fmt.Printf("Enter name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter roll number: ")
	fmt.Scanln(&roll)
	fmt.Printf("Enter subject1 and marks: ")
	fmt.Scanln(&sub1, &m1)
	fmt.Printf("Enter subject2 and marks: ")
	fmt.Scanln(&sub2, &m2)
	fmt.Printf("Enter subject3 and marks: ")
	fmt.Scanln(&sub3, &m3)

	fmt.Printf("Name\t")
	fmt.Printf("Roll No.\t")
	fmt.Printf("Subjects\t")
	fmt.Printf("Marks\t")
	fmt.Printf("Avg Marks\t")
	fmt.Println()
	fmt.Printf("%s", name)
	fmt.Printf("%4d", roll)
	fmt.Printf("%22s", sub1)
	fmt.Printf("%10g", m1)
	fmt.Println()
	fmt.Printf("%30s", sub2)
	fmt.Printf("%10g", m2)
	fmt.Println()
	fmt.Printf("%30s", sub3)
	fmt.Printf("%10g", m3)
	fmt.Println()

}

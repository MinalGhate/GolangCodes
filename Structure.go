package main

import "fmt"

type Books struct {
	title   string
	author  string
	sub     string
	book_id int
}

func main() {
	var Book1, Book2 Books

	Book1 = Books{"Go Prog", "Mahesh Kumar", "Go prog tutorial", 12345}

	// Book1.title = "Go Prog"
	// Book1.author = "Mahesh Kumar"
	// Book1.sub = "Go prog tutorial"
	// Book1.book_id = 12345

	Book2.title = "Telecom billing"
	Book2.author = "Zara ali"
	Book2.sub = "Telecom billing tutorial"
	Book2.book_id = 5678

	//fmt.Printf("%T", Book1)

	// fmt.Printf("Book1 title: %s\n", Book1.title)
	// fmt.Printf("Book1 author: %s\n", Book1.author)
	// fmt.Printf("Book1 subject: %s\n", Book1.sub)
	// fmt.Printf("Book1 book id: %d\n", Book1.book_id)

	printBook(Book1)

	// fmt.Printf("Book2 title: %s\n", Book2.title)
	// fmt.Printf("Book2 author: %s\n", Book2.author)
	// fmt.Printf("Book2 subject: %s\n", Book2.sub)
	// fmt.Printf("Book2 book id: %d\n", Book2.book_id)

	printBook(Book2)

}

func printBook(book Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.sub)
	fmt.Printf("Book book id: %d\n", book.book_id)
	fmt.Println()
}

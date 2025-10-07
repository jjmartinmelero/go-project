package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	// cannot access unexported fields *****************
	// var myBook = book.Book {
	// 	Title: "Go programming",
	// 	Author: "John Doe",
	// 	Pages: 300,
	// }

	myBook := book.NewBook("Go programming", "John Doe", 300)
	// myBook.PrintInfo()

	myBook.SetTitle("Advanced Go programming")
	// myBook.PrintInfo()


	myEBook := book.NewEBook("Learning Go", "Jane Smith", 250, 5, "PDF")
	// myEBook.PrintInfo()


	book.Print(myBook)
	book.Print(myEBook)

	fmt.Println(" ------------------------------------------------------------")
	myDog := animal.Dog{Name: "Buddy"}
	myCat := animal.Cat{Name: "Whiskers"}

	animal.MakeSound(myCat)
	animal.MakeSound(myDog)

	animals := []animal.Animal{
		myDog,
		myCat,
		animal.Dog{Name: "Rex"},
	}

	for _, animal := range animals {
		animal.Sound()
	}

}	
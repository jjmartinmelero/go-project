package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable){
	p.PrintInfo()
}

type Book struct {
	// Title  string public
	// Author string public
	// Pages  int public
	title  string
	author string
	pages  int
}

func NewBook(title string, author string, pages int) *Book{
	return &Book{
		title: title,
		author: author,
		pages: pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.title, b.author, b.pages)
}

type EBook struct {
	Book
	fileSize int
	format string
}

func NewEBook(title string, author string, pages int, fileSize int, format string) *EBook {
	return &EBook{
		Book: Book {title, author, pages},
		fileSize: fileSize,
		format: format,
	}
}

func (b *EBook) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d, FileSize: %d, Format: %s\n", b.title, b.author, b.pages, b.fileSize, b.format)
}
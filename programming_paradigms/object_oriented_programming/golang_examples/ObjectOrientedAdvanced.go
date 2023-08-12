package golangexamples

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Library struct {
	Name  string
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l Library) DisplayBooks() {
	fmt.Printf("Books in %s library:\n", l.Name)
	for _, book := range l.Books {
		fmt.Printf("%s by %s\n", book.Title, book.Author)
	}
}

func main() {
	library := Library{Name: "City Library"}

	book1 := Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}
	book2 := Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"}

	library.AddBook(book1)
	library.AddBook(book2)

	library.DisplayBooks()
}

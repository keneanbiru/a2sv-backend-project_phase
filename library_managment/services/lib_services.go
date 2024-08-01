package services

import (
	"errors"
	"fmt"
	"library_management/models"
	// "library_management/controllers"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}
type lib struct {
	books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *lib {
	return &lib{
		books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *lib) AddBook(book models.Book) {

	l.books[book.Id] = book

}
func (l *lib) RemoveBook(bookId int) error {
	// Check if the bookId exists
	if _, exists := l.books[bookId]; !exists {
		fmt.Println("book not found in the library")
		return errors.New("book not found in the library")
	}

	// Remove the book
	delete(l.books, bookId)
	fmt.Println("Book removed successfully.")
	return nil
}

func (l *lib) BorrowBook(bookID int, memberID int) error {
	book, ex := l.books[bookID]
	//if not available
	if !ex || book.Status == "borrowed" {
		return errors.New("Book not available")
	} else {
		member, ex := l.Members[memberID]
		//if member not found
		if !ex {
			return errors.New("Member not found")
		} else {
			book.Status = "borrowed"
			member.Br_books = append(member.Br_books, book)
		}
	}
	return nil
}

func (l *lib) ReturnBook(bookID int, memberID int) error {
	book, ex := l.books[bookID]
	//if not available
	if !ex || book.Status == "available" {
		return errors.New("Book not borrowed")
	} else {

		_, ex := l.Members[memberID]
		if !ex {
			return errors.New("Member not found")
		} else {
			book.Status = "available"
			l.books[bookID] = book
		}
	}
	member, _ := l.Members[memberID]
	// Remove the book from the borrowed list
	for i, b := range member.Br_books {
		if b.Id == bookID {
			member.Br_books = append(member.Br_books[:i], member.Br_books[i+1:]...)
			break
		}
	}
	l.Members[memberID] = member
	return nil
}

func (l *lib) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *lib) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return nil
	}
	return member.Br_books
}

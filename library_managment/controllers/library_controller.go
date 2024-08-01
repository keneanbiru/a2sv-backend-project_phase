package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func RunLibrary() {
	library := services.NewLibrary()

	// Sample data
	library.AddBook(models.Book{Id: 1, Title: "Go Programming", Author: "Author A", Status: "Available"})
	library.AddBook(models.Book{Id: 2, Title: "Python Programming", Author: "Author B", Status: "Available"})
	library.Members[1] = models.Member{ID: 1, Name: "John Doe"}

	var choice int
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter book author: ")
			fmt.Scan(&author)
			library.AddBook(models.Book{Id: id, Title: title, Author: author, Status: "Available"})
			fmt.Println("Book added successfully.")

		case 2:
			var bookID int
			fmt.Print("Enter book ID to remove: ")
			fmt.Scan(&bookID)
			library.RemoveBook(bookID)
			

		case 3:
			var bookID, memberID int
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Enter book ID to return: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}

		case 5:
			fmt.Println("Available Books:")
			for _, book := range library.ListAvailableBooks() {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
			}

		case 6:
			var memberID int
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			books := library.ListBorrowedBooks(memberID)
			if books == nil {
				fmt.Println("no book found.")
			} else {
				fmt.Println("Borrowed Books:")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
				}
			}

		case 0:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

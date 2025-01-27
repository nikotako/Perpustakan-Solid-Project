package main

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/handler"
	"Project_Niko/internal/repository"
	"Project_Niko/internal/usecase"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bookRepo := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	personRepo := repository.NewPersonRepository()
	personUsecase := usecase.NewPersonUsecase(personRepo)
	personHandler := handler.NewPersonHandler(personUsecase)

	loanRepo := repository.NewLoanRepository()
	loanUsecase := usecase.NewLoanUsecase(loanRepo, bookRepo, personRepo)
	loanHandler := handler.NewLoanHandler(loanUsecase)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n========== Library Management System ==========")
		fmt.Println("1. Manage Book")
		fmt.Println("2. Manage Person")
		fmt.Println("3. Manage Loans")
		fmt.Println("4. Exit")
		fmt.Println("===============================================")
		fmt.Println("Chose the menu: ")
		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			manageBooks(reader, bookHandler)
		case 2:
			managePersons(reader, personHandler)
		case 3:
			manageLoans(reader, loanHandler)
		case 4:
			fmt.Println("Exit, Thankyou")
			return
		default:
			fmt.Println("Invalid option, please try again brow")
		}
	}
}

func manageBooks(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	for {
		fmt.Println("\n========== Manage Books ==========")
		fmt.Println("1. Add Book")
		fmt.Println("2. Update Book")
		fmt.Println("3. Delete Book")
		fmt.Println("4. List All Books")
		fmt.Println("5. Back to Main Menu")
		fmt.Println("===================================")
		fmt.Print("Choose an option: ")
		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			addBook(reader, bookHandler)
		case 2:
			updateBook(reader, bookHandler)
		case 3:
			deleteBook(reader, bookHandler)
		case 4:
			listBooks(bookHandler)
		case 5:
			return
		default:
			fmt.Println("Invalid option. Please try again brow.")
		}
	}
}

func addBook(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	fmt.Print("Enter book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
		return
	}

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')

	book := domain.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
	}

	err = bookHandler.StoreNewBook(book)
	if err != nil {
		fmt.Println("Error storing book:", err)
	} else {
		fmt.Println("Book added successfully!")
	}
}

func updateBook(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	fmt.Print("Enter book ID to update: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
	}

	fmt.Print("Enter new book title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter new book author: ")
	author, _ := reader.ReadString('\n')

	book := domain.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
	}

	err = bookHandler.UpdateBook(book)
	if err != nil {
		fmt.Println("Error updating book:", err)
	} else {
		fmt.Println("Book updated successfully!")
	}
}

func deleteBook(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	fmt.Print("Enter book ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
	}

	book := domain.Book{
		ID: id,
	}

	err = bookHandler.DeleteBook(book)
	if err != nil {
		fmt.Println("Error deleting book:", err)
	} else {
		fmt.Println("Book deleted successfully!")
	}
}

func listBooks(bookHandler handler.BookHendlerInterface) {
	books := bookHandler.ListBooks()
	fmt.Println("\nBooks in Library")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func managePersons(reader *bufio.Reader, personHandler handler.PersonHandlerInterface) {
	for {
		fmt.Println("\n========== Manage Persons ==========")
		fmt.Println("1. Add Person")
		fmt.Println("2. Update Person")
		fmt.Println("3. Delete Person")
		fmt.Println("4. List All Person")
		fmt.Println("5. Back to Main Menu")
		fmt.Println("====================================")
		fmt.Print("Choose an option: ")
		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			addPerson(reader, personHandler)
		case 2:
			updatePerson(reader, personHandler)
		case 3:
			deletePerson(reader, personHandler)
		case 4:
			listPersons(personHandler)
		case 5:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addPerson(reader *bufio.Reader, personHandler handler.PersonHandlerInterface) {
	fmt.Print("Enter person ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid person ID: must be a positive integer")
		return
	}

	fmt.Print("Enter person name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter city: ")
	city, _ := reader.ReadString('\n')

	person := domain.Person{
		ID:   id,
		Name: strings.TrimSpace(name),
		Address: domain.Address{
			City: strings.TrimSpace(city),
		},
	}

	err = personHandler.StoreNewPerson(person)
	if err != nil {
		fmt.Println("Error storing person:", err)
	} else {
		fmt.Println("Person added successfully!")
	}
}

func updatePerson(reader *bufio.Reader, personHandler handler.PersonHandlerInterface) {
	fmt.Print("Enter person ID to update: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid person ID: must be a positive integer")
		return
	}

	fmt.Print("Enter new person name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter new city: ")
	city, _ := reader.ReadString('\n')

	person := domain.Person{
		ID:   id,
		Name: strings.TrimSpace(name),
		Address: domain.Address{
			City: strings.TrimSpace(city),
		},
	}

	err = personHandler.UpdatePerson(person)
	if err != nil {
		fmt.Println("Error updating person:", err)
	} else {
		fmt.Println("Person updated successfully!")
	}
}

func deletePerson(reader *bufio.Reader, personHandler handler.PersonHandlerInterface) {
	fmt.Print("Enter person ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
		return
	}

	person := domain.Person{
		ID: id,
	}

	err = personHandler.DeletePerson(person)
	if err != nil {
		fmt.Println("Error deleting person:", err)
	} else {
		fmt.Println("Person deleted successfully!")
	}
}

func listPersons(personHandler handler.PersonHandlerInterface) {
	persons := personHandler.ListPersons()
	fmt.Println("\nPersons in library:")
	for _, person := range persons {
		fmt.Printf("ID: %d, Name: %s, City: %s\n", person.ID, person.Name, person.Address.City)
	}
}

func manageLoans(reader *bufio.Reader, loanHandler handler.LaonHandlerInterface) {
	for {
		fmt.Println("\n========== Manage Loans ==========")
		fmt.Println("1. Add Loan")
		fmt.Println("2. Return Loan")
		fmt.Println("3. List All Loans")
		fmt.Println("4. Back to Main Menu")
		fmt.Println("==================================")
		fmt.Print("Choose an option: ")
		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			addLoan(reader, loanHandler)
		case 2:
			returnLoan(reader, loanHandler)
		case 3:
			listLoans(loanHandler)
		case 4:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addLoan(reader *bufio.Reader, loanHandler handler.LaonHandlerInterface) {
	fmt.Print("Enter loan ID: ")
	loanIDStr, _ := reader.ReadString('\n')
	loanID, err := strconv.Atoi(strings.TrimSpace(loanIDStr))
	if err != nil {
		fmt.Println("Invalid loan ID: must be a positive integer")
		return
	}

	fmt.Print("Enter book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(bookIDStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
		return
	}

	fmt.Print("Enter person ID: ")
	personIDStr, _ := reader.ReadString('\n')
	personID, err := strconv.Atoi(strings.TrimSpace(personIDStr))
	if err != nil {
		fmt.Println("Invalid person ID: must be a positive integer")
		return
	}

	if loanHandler.IsBookLoaned(bookID) {
		fmt.Println("Book is already loan by other.")
		return
	}

	if loanHandler.IsPersonLoaning(personID) {
		fmt.Println("Person already loan.")
		return
	}


	fmt.Print("Enter due date: ")
	dueDate, _ := reader.ReadString('\n')

	loan := domain.Loan{
		ID: loanID,
		Book: domain.Book{
			ID: bookID,
		},
		Person: domain.Person{
			ID: personID,
		},
		DueDate:  strings.TrimSpace(dueDate),
		Returned: false,
	}

	err = loanHandler.StoreNewLoan(loan)
	if err != nil {
		fmt.Println("Error adding loan:", err)
	} else {
		fmt.Println("Loan added successfully!")
	}
}

func returnLoan(reader *bufio.Reader, loanHandler handler.LaonHandlerInterface) {
	fmt.Print("Enter loan ID to return: ")
	loanIDStr, _ := reader.ReadString('\n')
	loanID, err := strconv.Atoi(strings.TrimSpace(loanIDStr))
	if err != nil {
		fmt.Println("Invalid loan ID: must be a positive integer")
		return
	}

	// kita cari loan yang saya mau update
	loans := loanHandler.ListLoans()
    var loanToUpdate domain.Loan
    found := false
    for _, loan := range loans {
        if loan.ID == loanID {
            loanToUpdate = loan
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Loan ID not found")
        return
    }

	loanToUpdate.Returned = true
	err = loanHandler.UpdateLoan(loanToUpdate)
	if err != nil {
		fmt.Println("Error returning loan:", err)
	} else {
		fmt.Println("Loan returned successfully!")
	}
}

func listLoans(loanHandler handler.LaonHandlerInterface) {
	loans := loanHandler.ListLoans()
	fmt.Println("\nLoans in library:")
	for _, loan := range loans {
		fmt.Printf("Loan ID: %d, Book ID: %d, Person ID: %d, Due Date: %s, Returned: %t\n",
			loan.ID, loan.Book.ID, loan.Person.ID, loan.DueDate, loan.Returned)
	}
}

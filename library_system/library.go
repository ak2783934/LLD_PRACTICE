package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Library struct {
	Users   map[string]*User
	Books   map[string]*Book
	Tickets map[string]*Ticket
}

/*
AddMember
AddBook
SearchBook
BorrowBook
ReturnBook
*/

func (l *Library) AddMember(name, email string) (*User, error) {
	// check if the user already exists.
	_, ok := l.Users[email]
	if ok {
		return nil, errors.New("user already exists")
	}

	user := &User{
		Id:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	l.Users[email] = user
	return user, nil
}

func (l *Library) AddBook(title, genre string, price int, author Author) *Book {
	book := &Book{
		Id:                uuid.New().String(),
		Title:             title,
		Genre:             genre,
		OverDueRatePerDay: price,
		Author:            author,
		IsBorrowed:        false,
	}
	l.Books[book.Id] = book

	return book
}

func (l *Library) SearchBook(id string, name string, genre string) []*Book {
	// logic to fetch using id first.
	// using nsame
	// use genre.

}

func (l *Library) BorrowBook(bookId, userID string) (*Ticket, error) {
	book, ok := l.Books[bookId]
	if !ok || book.IsBorrowed {
		return nil, errors.New("book does not exist")
	}

	borrowTicket := CreateTicket(book, userID)
	l.Tickets[borrowTicket.Id] = borrowTicket
	// Make the book not available?
	book.IsBorrowed = true
	return borrowTicket, nil
}

func (l *Library) ReturnBook(ticketID string) (*Ticket, error) {
	ticket, ok := l.Tickets[ticketID]
	if !ok {
		return nil, errors.New("ticket doesn't exist")
	}

	bookId := ticket.BookID
	book := l.Books[bookId]
	bookRate := book.OverDueRatePerDay
	cost := 0
	if ticket.DueDate.Compare(time.Now()) == 1 {
		numberOfDaysOverDue := time.Since(ticket.DueDate).Hours() / 24
		if numberOfDaysOverDue > 0 {
			cost = bookRate * int(numberOfDaysOverDue)
		}
	}

	// complete the ticket
	ticket.ReturnDate = time.Now()
	ticket.Price = cost
	ticket.Status = PAID

	// empty the book
	book.IsBorrowed = false

	return ticket, nil
}

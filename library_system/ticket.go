package main

import (
	"time"

	"github.com/google/uuid"
)

type TicketStatus int

const (
	PAID TicketStatus = iota
	PENDING
)

type Ticket struct {
	Id         string
	BookID     string
	UserID     string
	StartDate  time.Time
	ReturnDate time.Time
	DueDate    time.Time
	Status     TicketStatus
	Price      int
}

// Default due date of 15 days.

func CreateTicket(book *Book, borrowerID string) *Ticket {
	return &Ticket{
		Id:        uuid.New().String(),
		BookID:    book.Id,
		UserID:    borrowerID,
		StartDate: time.Now(),
		DueDate:   time.Now().AddDate(0, 0, 15),
		Status:    PENDING,
	}
}

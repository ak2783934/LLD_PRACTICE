package main

import (
	"github.com/google/uuid"
)

type OrderStatus int

const (
	PLACED OrderStatus = iota
	CANCELLED
	DELIVERED
)

type Order struct {
	OrderID string
	UserID  string
	Items   []OrderItem
	Status  OrderStatus
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     int
}

func PlaceOrder(userID string, Items []OrderItem) *Order {
	return &Order{
		OrderID: uuid.NewString(),
		UserID:  userID,
		Items:   Items,
		Status:  PLACED,
	}
}

func (O *Order) UpdateStatus(status OrderStatus) {
	O.Status = status
}

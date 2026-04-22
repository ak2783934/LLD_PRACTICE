package main

type DiscountStrategy interface {
	getDiscount(Order, User) int
}

package main

type Item struct {
	id       string
	price    int
	category string
}

type Order struct {
	items     []Item
	totalCost int
}

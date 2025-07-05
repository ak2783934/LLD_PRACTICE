package main

type Book struct {
	Id                string
	Title             string
	Author            Author
	Genre             string
	OverDueRatePerDay int
	IsBorrowed        bool
}

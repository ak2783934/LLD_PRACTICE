package main

type CarType int

const (
	SEDAN CarType = iota
	SUV
	HATCHBACK
)

type Car struct {
	ID         int
	RegNo      string
	Brand      string
	Model      string
	CarType    CarType
	MakingYear string
	HourlyRate int
	BranchID   int
}

package main

type State interface {
	InsertCoin(money int) error
	SelectProduct(code string) error
	Dispense() error
	Refund() (int, error)
	String() string
}

// what are all the states that we should have?

// NoMoneyState?
// HasMoneyState?
// DispensingState
// OutOfStockState



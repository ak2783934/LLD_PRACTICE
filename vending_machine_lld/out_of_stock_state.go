package main

import "errors"

type OutOfStockState struct {
	vm *VendingMachine
}

func (n *OutOfStockState) InsertCoin(money int) error {
	return errors.New("no items available")
}
func (n *OutOfStockState) SelectProduct(code string) error {
	return errors.New("no items available")
}
func (n *OutOfStockState) Dispense() error {
	return errors.New("no items available")
}
func (n *OutOfStockState) Refund() (int, error) {
	return 0, errors.New("no items available")
}
func (n *OutOfStockState) String() string {
	return "No items available"
}

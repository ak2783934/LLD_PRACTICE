package main

import "errors"

type NoMoneyState struct {
	vm *VendingMachine
}

func (n *NoMoneyState) InsertCoin(money int) error {
	if n.vm.totalStock() == 0 {
		n.vm.setState(&OutOfStockState{n.vm})
		return errors.New("no items available")
	}
	if money <= 0 {
		return errors.New("invalid amount")
	}

	n.vm.balance += money
	n.vm.setState(&HasMoneyState{n.vm})
	return nil
}
func (n *NoMoneyState) SelectProduct(code string) error {
	return errors.New("no money, please insert money first")
}
func (n *NoMoneyState) Dispense() error {
	return errors.New("no money, nothing to dispense")
}
func (n *NoMoneyState) Refund() (int, error) {
	return 0, errors.New("no money, no balance to refund")
}
func (n *NoMoneyState) String() string {
	return "No money"
}

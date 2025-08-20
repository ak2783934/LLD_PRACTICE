package main

import "errors"

type NoMoneyState struct {
	vm *VendingMachine
}

func (n *NoMoneyState) InsertCoin(money int) error {
	n.vm.balance += money
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

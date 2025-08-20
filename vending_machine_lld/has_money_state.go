package main

import "errors"

type HasMoneyState struct {
	vm *VendingMachine
}

func (n *HasMoneyState) InsertCoin(money int) error {
	n.vm.balance += money
	return nil
}
func (n *HasMoneyState) SelectProduct(code string) error {
	product, ok := n.vm.products[code]
	if !ok {
		return errors.New("requested item doesn't exist")
	}

	count := n.vm.stock[code]
	if count == 0 {
		return errors.New("requested item is out of stock")
	}

	if product.Price > n.vm.balance {
		return errors.New("insufficient balance")
	}

	n.vm.choice = code
	n.vm.setState(&DispensingState{n.vm})
	return nil
}
func (n *HasMoneyState) Dispense() error {
	return errors.New("select item first")
}
func (n *HasMoneyState) Refund() (int, error) {
	refund := n.vm.balance
	n.vm.choice = ""
	n.vm.balance = 0
	n.vm.setState(&NoMoneyState{n.vm})
	return refund, nil
}
func (n *HasMoneyState) String() string {
	return "Has Money"
}

package main

import (
	"errors"
	"fmt"
)

type DispensingState struct {
	vm *VendingMachine
}

func (n *DispensingState) InsertCoin(money int) error {
	return errors.New("dispensing item, can't insert moneu")
}
func (n *DispensingState) SelectProduct(code string) error {
	return errors.New("already selected; dispensing item")
}
func (n *DispensingState) Dispense() error {
	itemCode := n.vm.choice
	product := n.vm.products[itemCode]
	price := product.Price
	n.vm.stock[itemCode] = n.vm.stock[itemCode] - 1
	if n.vm.balance > price {
		fmt.Println("refunding amount: ", price)
		n.vm.balance = 0
	}

	n.vm.choice = ""

	if n.vm.totalStock() == 0 {
		n.vm.setState(&OutOfStockState{n.vm})
	} else {
		n.vm.setState(&NoMoneyState{n.vm})
	}

	return nil
}
func (n *DispensingState) Refund() (int, error) {
	return 0, errors.New("can't refund, dispensing")
}
func (n *DispensingState) String() string {
	return "dispensing"
}

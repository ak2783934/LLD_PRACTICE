package main

type VendingMachine struct {
	state    State
	products map[string]Product // code -> Product
	stock    map[string]int     // code -> QTY

	balance int
	choice  string
}

func (v *VendingMachine) setState(state State) {
	v.state = state
}

func (v *VendingMachine) totalStock() int {
	total := 0
	for _, q := range v.stock {
		total += q
	}
	return total
}

func NewVendingMachine(products map[string]Product, stock map[string]int) *VendingMachine {
	vendingMachine := &VendingMachine{
		products: products,
		stock:    stock,
		balance:  0,
	}

	vendingMachine.setState(&NoMoneyState{})
	if vendingMachine.totalStock() == 0 {
		vendingMachine.setState(&OutOfStockState{})
	}
	return vendingMachine
}

func (v *VendingMachine) InsertCoin(money int) error {
	return v.state.InsertCoin(money)
}

func (v *VendingMachine) SelectProduct(code string) error {
	return v.state.SelectProduct(code)
}

func (v *VendingMachine) Dispense() error {
	return v.state.Dispense()
}

func (v *VendingMachine) Refund() (int, error) {
	return v.state.Refund()
}

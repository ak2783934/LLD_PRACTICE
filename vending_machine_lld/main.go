package main

import "fmt"

func main() {
	products := map[string]Product{
		"A1": {Name: "Chips", Price: 150},
		"B2": {Name: "Soda", Price: 120},
		"C3": {Name: "Chocolate", Price: 200},
	}
	stock := map[string]int{"A1": 1, "B2": 2, "C3": 1}

	vm := NewVendingMachine(products, stock)

	fmt.Println("State:", vm.state.String())
	_ = vm.InsertCoin(100)        // insufficient yet
	err := vm.SelectProduct("C3") // won't work, not enough funds
	fmt.Println("Select C3:", err)
	_ = vm.InsertCoin(100)     // now enough
	_ = vm.SelectProduct("C3") // OK, transitions to Dispensing
	fmt.Println("State:", vm.state.String())
	_ = vm.Dispense() // Dispense chocolate + change

	fmt.Println("State:", vm.state.String())
	_ = vm.InsertCoin(200)
	_ = vm.SelectProduct("A1") // vend chips
	_ = vm.Dispense()

	// Refund example
	_ = vm.InsertCoin(120)
	ref, _ := vm.Refund()
	fmt.Println("Refunded:", ref)

	// Try an out-of-stock path quickly
	_ = vm.InsertCoin(120)
	_ = vm.SelectProduct("B2")
	_ = vm.Dispense()
	_ = vm.InsertCoin(120)
	_ = vm.SelectProduct("B2")
	_ = vm.Dispense()

	fmt.Println("Final State:", vm.state.String())
}

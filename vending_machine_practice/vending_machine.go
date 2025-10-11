package main

import (
	"errors"
	"fmt"
)

type Item struct {
	count int
	price int
}

type VendingMachine struct {
	Inventory    map[string]*Item
	state        VendingMachineState
	money        int
	selectedItem string
}

func NewVendingMachine(inventory map[string]*Item) *VendingMachine {
	return &VendingMachine{
		Inventory:    inventory,
		state:        &IdleState{},
		money:        0,
		selectedItem: "",
	}
}

func (v *VendingMachine) InsertMoney(amount int) {
	fmt.Println("inserting money ", amount)
	v.state.InsertMoney(v, amount)
}

func (v *VendingMachine) SelectItem(itemCode string) error {
	fmt.Println("selecting item ", itemCode)
	return v.state.SelectItem(v, itemCode)
}

func (v *VendingMachine) DispenseItem() {
	v.state.DispenseItem(v)
}

func (v *VendingMachine) CancelRequest() {
	v.state.CancelRequest(v)
}

func (v *VendingMachine) ReturnChange() {
	v.state.ReturnChange(v)
}

func (v *VendingMachine) GetItemStock(itemCode string) (int, error) {
	item, ok := v.Inventory[itemCode]
	if !ok {
		return 0, errors.New("item not available")
	}
	return item.count, nil
}

func (v *VendingMachine) UpdateItemStock(itemCode string, count int) error {
	item, ok := v.Inventory[itemCode]
	if !ok {
		return errors.New("item not available")
	}
	item.count += count
	return nil
}

func (v *VendingMachine) GetItemPrice(itemCode string) (int, error) {
	item, ok := v.Inventory[itemCode]
	if !ok {
		return 0, errors.New("item not available")
	}
	return item.price, nil
}

func (v *VendingMachine) IsItemAvailable(itemCode string) bool {
	item, ok := v.Inventory[itemCode]
	if !ok {
		return false
	}
	if item.count > 0 {
		return true
	}
	return false
}

func (v *VendingMachine) SetState(state VendingMachineState) {
	v.state = state
}

func (v *VendingMachine) PrintInventory() {
	fmt.Println("############### Printing inventory ##################")
	for idx, item := range v.Inventory {
		fmt.Println("Item Name : ", idx, "item count : ", item.count)
	}
	fmt.Println("############### Printing inventory ##################")
}

type VendingMachineState interface {
	InsertMoney(v *VendingMachine, amount int)
	SelectItem(v *VendingMachine, itemCode string) error
	DispenseItem(v *VendingMachine)
	CancelRequest(v *VendingMachine) error
	ReturnChange(v *VendingMachine)
}

type IdleState struct{}

func (i *IdleState) InsertMoney(v *VendingMachine, amount int) {
	fmt.Println("inserting money into machine")
	v.money = amount
	v.SetState(&HasMoneyState{})
}
func (i *IdleState) SelectItem(v *VendingMachine, itemCode string) error {
	isAvailable := v.IsItemAvailable(itemCode)
	if !isAvailable {
		fmt.Println("Selected item not available")
		return errors.New("Item not available")
	}
	v.selectedItem = itemCode
	return nil
}
func (i *IdleState) DispenseItem(v *VendingMachine) {
	fmt.Println("No money inserted or item selected")
}
func (i *IdleState) CancelRequest(v *VendingMachine) error {
	fmt.Println("No request processing to cancel")
	return nil
}
func (i *IdleState) ReturnChange(v *VendingMachine) {
	fmt.Println("No amount inserted")
}

type HasMoneyState struct{}

func (h *HasMoneyState) InsertMoney(v *VendingMachine, amount int) {
	v.money += amount
	fmt.Println("added more money")
}
func (h *HasMoneyState) SelectItem(v *VendingMachine, itemCode string) error {
	if v.selectedItem != "" {
		fmt.Println("item already selected")
	} else {
		itemPrice, err := v.GetItemPrice(itemCode)
		if err != nil {
			fmt.Println("item not available")
			return errors.New("item not available")
		}
		if itemPrice > v.money {
			v.selectedItem = ""
			v.ReturnChange()
			v.SetState(&IdleState{})
			fmt.Println("Not enough money, try again")
			return errors.New("Not enough money, try again")
		}
		v.selectedItem = itemCode
	}
	v.SetState(&DispensingState{})
	return nil
}
func (h *HasMoneyState) DispenseItem(v *VendingMachine) {
	fmt.Println("Please add money first or select item")
}
func (h *HasMoneyState) CancelRequest(v *VendingMachine) error {
	v.selectedItem = ""
	v.ReturnChange()
	v.SetState(&IdleState{})
	return nil
}
func (h *HasMoneyState) ReturnChange(v *VendingMachine) {
	fmt.Println("returned money change ", v.money)
	v.money = 0
}

type DispensingState struct{}

func (d *DispensingState) InsertMoney(v *VendingMachine, amount int) {
	fmt.Println("In dispensing state, can't add money")
}
func (d *DispensingState) SelectItem(v *VendingMachine, itemCode string) error {
	fmt.Println("item already dispensing, can't select item now")
	return errors.New("can't select item during dispensing state")
}
func (d *DispensingState) DispenseItem(v *VendingMachine) {
	fmt.Println("dispensing the item ", v.selectedItem)
	// update the inventory.
	itemPrice, _ := v.GetItemPrice(v.selectedItem)

	// reducing the count by one
	v.UpdateItemStock(v.selectedItem, -1)

	// reduce the amount.
	v.money -= itemPrice

	// if amount left, call return change.
	if v.money > 0 {
		v.ReturnChange()
	}
	// make machine idle
	v.SetState(&IdleState{})
	v.selectedItem = ""
}
func (d *DispensingState) CancelRequest(v *VendingMachine) error {
	fmt.Println("can't cancel request during dispensing state")
	return errors.New("can't cancel request during dispensing state")
}
func (d *DispensingState) ReturnChange(v *VendingMachine) {
	fmt.Println("Returning the amount ", v.money)
	v.money = 0
}

func main() {
	inventory := map[string]*Item{
		"coke": &Item{
			price: 40,
			count: 10,
		},
		"cookies": &Item{
			price: 10,
			count: 10,
		},
		"chips": &Item{
			price: 40,
			count: 10,
		},
		"water": &Item{
			price: 20,
			count: 10,
		},
	}

	vendingMachine := NewVendingMachine(inventory)

	vendingMachine.PrintInventory()

	// success cases
	vendingMachine.InsertMoney(15)
	err := vendingMachine.SelectItem("cookies")
	if err == nil {
		vendingMachine.DispenseItem()
	}

	vendingMachine.PrintInventory()

	// succes cases
	vendingMachine.InsertMoney(10)
	err1 := vendingMachine.SelectItem("water")
	if err1 == nil {
		vendingMachine.DispenseItem()
	}

	vendingMachine.PrintInventory()

	// failure case cases
	vendingMachine.InsertMoney(50)
	err2 := vendingMachine.SelectItem("chips")
	if err2 == nil {
		vendingMachine.DispenseItem()
	}
	vendingMachine.PrintInventory()
}

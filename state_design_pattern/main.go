package main



type State interface {
    RequestItem()
    AddItem(count int)
    InsertMoney(money int)
    DispenseItem()
}

type HasItemState struct {
    vendingMachine *VendingMachine
}

func (s *HasItemState) RequestItem() {
    fmt.Println("Item requested.")
    s.vendingMachine.setState(s.vendingMachine.itemRequestedState)
}
func (s *HasItemState) AddItem(count int) { /* ... */ }
func (s *HasItemState) InsertMoney(money int) {
    fmt.Println("First request item, then insert money.")
}
func (s *HasItemState) DispenseItem() {
    fmt.Println("First insert money.")
}

type NoItemState struct {
    vendingMachine *VendingMachine
}

func (s *NoItemState) RequestItem()    { fmt.Println("No item available.") }
func (s *NoItemState) AddItem(count int) {/* ... */ }
func (s *NoItemState) InsertMoney(money int) { fmt.Println("No item available.") }
func (s *NoItemState) DispenseItem()    { fmt.Println("No item available.") }


type VendingMachine struct {
    hasItemState        State
    noItemState         State
    itemRequestedState  State
    hasMoneyState       State
    currentState        State
    itemCount           int
}

func NewVendingMachine(count int) *VendingMachine {
    m := &VendingMachine{itemCount: count}
    m.hasItemState = &HasItemState{vendingMachine: m}
    m.noItemState = &NoItemState{vendingMachine: m}
    // ... (Initialize other states)
    if count > 0 {
        m.setState(m.hasItemState)
    } else {
        m.setState(m.noItemState)
    }
    return m
}

func (v *VendingMachine) setState(s State) {
    v.currentState = s
}

// Delegate methods:
func (v *VendingMachine) RequestItem() {
    v.currentState.RequestItem()
}
func (v *VendingMachine) AddItem(count int) {
    v.currentState.AddItem(count)
}
func (v *VendingMachine) InsertMoney(money int) {
    v.currentState.InsertMoney(money)
}
func (v *VendingMachine) DispenseItem() {
    v.currentState.DispenseItem()
}


func main() {
    vm := NewVendingMachine(2)

    vm.RequestItem()      // "Item requested."
    vm.InsertMoney(10)    // "Please insert money."
    vm.DispenseItem()     // "First insert money."
    // Continue with state-changing actions...
}


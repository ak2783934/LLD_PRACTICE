You need to design a Vending Machine system.

The machine dispenses items like chips, chocolates, and drinks after accepting money. It should support multiple states like when it’s idle, has received money, out of stock, etc.

Requirements:

The vending machine should allow a user to:

    Insert coins or currency notes.
    Select a product.
    Dispense the selected item if sufficient money is inserted.
    Return balance (if any).
    Cancel the transaction and return inserted money.

The machine can go into different states like:
    Idle (waiting for user)
    Has Money (user inserted coins)
    Dispense (when product is being released)
    Sold Out / Maintenance mode (no stock)

The design should be extensible — new states or actions should be easily added.

You need to handle edge cases like:
    Selecting product before inserting money.
    Insufficient money for selected product.    
    Out of stock.
    Cancelling a transaction midway.

The system should log meaningful messages for each action.



Assumption: 
Each item can be in any number, that depends on the stock. 
Each item has different pricing. 
only one request at a time. 
Cancel of request is allowed. 
Not taking care of the part where we have to restock any system. While initiation only we will stock it. 


Flow -> 
Idle -> user insert money -> selects an item -> item given -> change returned. 
can be cancelled at any state, 


Things to identify 

States
Functions of the initails machine
Functions in a state
VendingMachine object

Item{
    count int
    price int
}

VendingMachine{
    Inventory map[string]Item
    state VendingMachineState
    money int

    InsertMoney()
    SelectItem()
    DispenseItem()
    CancelRequest()
    ReturnChange()
}
NewVendingMachine()

VendingMachineState interface 

IdleState
HasMoneyState
DispensingItemState
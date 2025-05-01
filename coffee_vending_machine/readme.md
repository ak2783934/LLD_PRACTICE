
# Designing a Coffee Vending Machine
## Requirements

- The coffee vending machine should support different types of coffee, such as espresso, cappuccino, and latte.
- Each type of coffee should have a specific price and recipe (ingredients and their quantities).
- The machine should have a menu to display the available coffee options and their prices.
- Users should be able to select a coffee type and make a payment.
- The machine should dispense the selected coffee and provide change if necessary.
- The machine should track the inventory of ingredients and notify when they are running low.
- The machine should handle multiple user requests concurrently and ensure thread safety.


## Thought process
Coffee

Recipe 

Ingredient 

Price 

Payment should work and return left money. 

Ingredient low alert? Can trigger after each orders? 

concurrent 


User orders a coffee. 
we show the price and ask for payments.
simultenously we also check if all the ingredients are available based on recipe. Recipe can be like 5 suger, 2 coffee, 2 milk, 1 water, 1 ice.. 
So there should be a inventory, where we can ask these informations. 

If things are present, prepare the coffee and dispense. 
And return the change. 

Money management is not the task of this. 

## classes

```
CoffeeMachine
    - list of coffee types 
    - order coffee
    - pay


CoffeeType
    - id
    - name 
    - prices
    - Recipe


Recipe
    - id 
    - array[{ingredient, quantity}]

Inventory(singleton)
    - list of {ingredient, quantity}


Ingredient 
    - id
    - name
    



## apis
What are the apis to be designed?

## schema

```
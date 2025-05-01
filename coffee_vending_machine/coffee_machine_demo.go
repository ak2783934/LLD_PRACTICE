package coffeevendingmachine

import "fmt"

func Run() {
	CapaccinoRecipe := NewRecipe("1", map[Ingredient]int{
		CoffeePowder: 2,
		Milk:         1,
		Sugar:        1,
		Water:        1,
	})

	fmt.Println("Capaccino Recipe:", CapaccinoRecipe)

	// LatteRecipe := NewRecipe("2", map[Ingredient]int{
	// 	CoffeePowder: 1,
	// 	Milk:         2,
	// 	Sugar:        1,
	// 	Water:        1,
	// })

	// ColdBrewRecipe := NewRecipe("3", map[Ingredient]int{
	// 	CoffeePowder: 3,
	// 	Water:        2,
	// })

	// HotChocolateRecipe := NewRecipe("4", map[Ingredient]int{
	// 	Chocolate: 2,
	// 	Milk:      1,
	// 	Sugar:     1,
	// 	Water:     1,
	// })
	// BlackCoffeeRecipe := NewRecipe("5", map[Ingredient]int{
	// 	CoffeePowder: 2,
	// 	Water:        2,
	// })

	// Capaccino := CoffeeType{
	// 	Id:     "1",
	// 	Name:   "Capaccino",
	// 	Price:  50,
	// 	Recipe: CapaccinoRecipe,
	// }

	// Latte := CoffeeType{
	// 	Id:     "2",
	// 	Name:   "Latte",
	// 	Price:  40,
	// 	Recipe: LatteRecipe,
	// }

	// ColdBrew := CoffeeType{
	// 	Id:     "3",
	// 	Name:   "Cold Brew",
	// 	Price:  60,
	// 	Recipe: ColdBrewRecipe,
	// }

	// HotChocolate := CoffeeType{
	// 	Id:     "4",
	// 	Name:   "Hot Chocolate",
	// 	Price:  30,
	// 	Recipe: HotChocolateRecipe,
	// }

	// BlackCoffee := CoffeeType{
	// 	Id:     "5",
	// 	Name:   "Black Coffee",
	// 	Price:  20,
	// 	Recipe: BlackCoffeeRecipe,
	// }

	// CoffeeMachine := NewCoffeeMachine()
	// // add ingredients to inventory
	// CoffeeMachine.AddIngredient(CoffeePowder, 10)
	// CoffeeMachine.AddIngredient(Milk, 5)
	// CoffeeMachine.AddIngredient(Sugar, 20)
	// CoffeeMachine.AddIngredient(Water, 15)
	// CoffeeMachine.AddIngredient(Chocolate, 8)
	// CoffeeMachine.AddIngredient(Cream, 10)

	// // add coffee types to the machine
	// CoffeeMachine.AddCoffeeType(Capaccino)
	// CoffeeMachine.AddCoffeeType(Latte)
	// CoffeeMachine.AddCoffeeType(ColdBrew)
	// CoffeeMachine.AddCoffeeType(HotChocolate)
	// CoffeeMachine.AddCoffeeType(BlackCoffee)
	// CoffeeMachine.ShowAllCoffeeAndPrices()
}

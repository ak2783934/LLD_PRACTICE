package coffeevendingmachine

type CoffeeMachine struct {
	Coffees   map[string]CoffeeType
	Inventory Inventory
}

func NewCoffeeMachine() *CoffeeMachine {
	return &CoffeeMachine{
		Coffees:   map[string]CoffeeType{},
		Inventory: *NewInventory(),
	}
}

func (c *CoffeeMachine) AddCoffeeType(coffee CoffeeType) {
	c.Coffees[coffee.Id] = coffee
}

func (c *CoffeeMachine) RemoveCoffeeType(coffeeId string) {
	delete(c.Coffees, coffeeId)
}

func (c *CoffeeMachine) ShowAllCoffeeAndPrices() {
	for _, coffee := range c.Coffees {
		println("Coffee ID:", coffee.Id, "Name:", coffee.Name, "Price:", coffee.Price)
	}
}

func (c *CoffeeMachine) AddIngredient(ingredient Ingredient, quantity int) {
	c.Inventory.AddIngredient(ingredient, quantity)
}

func (c *CoffeeMachine) RemoveIngredient(ingredient Ingredient, quantity int) bool {
	return c.Inventory.RemoveIngredient(ingredient, quantity)
}

func (c *CoffeeMachine) GetIngredientQuantity(ingredient Ingredient) int {
	return c.Inventory.GetQuantity(ingredient)
}

func (c *CoffeeMachine) HasSufficientIngredients(coffee CoffeeType) bool {
	return c.Inventory.HasSufficientIngredients(coffee.Recipe.Ingredients)
}

func (c *CoffeeMachine) GetCoffeePrice(coffeeId string) int {
	for _, coffee := range c.Coffees {
		if coffee.Id == coffeeId {
			return coffee.Price
		}
	}
	return 0
}

func (c *CoffeeMachine) OrderCoffee(coffeeId string, paidAmount int) (bool, int) {
	// check if the coffee exists
	coffee, exists := c.Coffees[coffeeId]
	if !exists {
		return false, paidAmount // coffee not found
	}

	// check if the user has paid enough
	coffeePrice := coffee.Price
	if paidAmount < coffeePrice {
		return false, paidAmount // not enough money
	}

	// check if the machine has enough ingredients
	if c.HasSufficientIngredients(coffee) == false {
		return false, paidAmount // not enough ingredients
	}

	// deduct the ingredients from the inventory
	c.MakeCoffee(coffee)

	return true, paidAmount - coffeePrice // return the change amount
}

func (c *CoffeeMachine) MakeCoffee(coffee CoffeeType) {
	requiredIngredients := coffee.Recipe.Ingredients

	for Ingredient, quantity := range requiredIngredients {
		c.Inventory.RemoveIngredient(Ingredient, quantity)
	}
	println("Coffee", coffee.Name, "is being prepared with the following ingredients:")
}

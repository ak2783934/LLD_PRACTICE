package coffeevendingmachine

type Inventory struct {
	Ingredients map[Ingredient]int // Ingredient ID to quantity
}

func NewInventory() *Inventory {
	return &Inventory{
		Ingredients: make(map[Ingredient]int),
	}
}

func (i *Inventory) AddIngredient(ingredient Ingredient, quantity int) {
	i.Ingredients[ingredient] += quantity
}

func (i *Inventory) RemoveIngredient(ingredient Ingredient, quantity int) bool {
	if _, exists := i.Ingredients[ingredient]; !exists {
		return false
	}
	if i.Ingredients[ingredient] < quantity {
		return false
	}
	i.Ingredients[ingredient] -= quantity
	return true
}

func (i *Inventory) GetQuantity(ingredient Ingredient) int {
	if quantity, exists := i.Ingredients[ingredient]; exists {
		return quantity
	}
	return 0
}

func (i *Inventory) HasSufficientIngredients(ingredients map[Ingredient]int) bool {
	for id, quantity := range ingredients {
		if i.GetQuantity(id) < quantity {
			return false
		}
	}
	return true
}

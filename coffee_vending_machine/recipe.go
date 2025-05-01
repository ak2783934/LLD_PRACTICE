package coffeevendingmachine

type Recipe struct {
	Id          string
	Ingredients map[Ingredient]int // Ingredient to quantity
}

func NewRecipe(id string, ingredients map[Ingredient]int) *Recipe {
	return &Recipe{
		Id:          id,
		Ingredients: ingredients,
	}
}

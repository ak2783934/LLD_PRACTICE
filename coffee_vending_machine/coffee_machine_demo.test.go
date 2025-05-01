package coffeevendingmachine

import (
	"fmt"
	"testing"
)

func TestCoffeeMachineDemo(t *testing.T) {
	// Initialize coffee machine
	machine := NewCoffeeMachine()
	t.Log("Coffee Machine initialized")
	fmt.Println(machine.Coffees)

	// // Test different types of coffee
	// coffeeTypes := []string{"Espresso", "Latte", "Cappuccino"}

	// for _, coffeeType := range coffeeTypes {
	// 	t.Logf("\nMaking %s:", coffeeType)
	// 	machine.MakeCoffee(coffeeType)

	// 	t.Logf("Coffee Details:")
	// 	t.Logf("- Type: %s", coffee.GetType())
	// 	t.Logf("- Water: %d ml", coffee.GetWater())
	// 	t.Logf("- Coffee: %d g", coffee.GetCoffee())
	// 	t.Logf("- Milk: %d ml", coffee.GetMilk())
	// 	t.Logf("- Price: $%.2f", coffee.GetPrice())
	// }

	t.Log("\nCoffee Machine Demo Completed")
}

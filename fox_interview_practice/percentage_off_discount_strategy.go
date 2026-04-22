package main

type PercentageOffDiscountStrategy struct {
	minPurchaseAmount  int
	discountPercentage int
}

func (p *PercentageOffDiscountStrategy) getDiscount(order Order, user User) int {
	if order.totalCost < p.minPurchaseAmount {
		return 0
	}
	return (order.totalCost * p.discountPercentage) / 100
}

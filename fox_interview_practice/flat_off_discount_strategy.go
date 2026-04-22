package main

type FlatOffDiscountStrategy struct {
	discountAmount    int
	minPurchaseAmount int
}

func (f *FlatOffDiscountStrategy) getDiscount(order Order, user User) int {
	if order.totalCost < f.minPurchaseAmount {
		return 0
	}
	return f.discountAmount
}

package main

type BuyXGetYDiscountStrategy struct {
	minBuy  int
	getFree int
}

func (b *BuyXGetYDiscountStrategy) getDiscount(order Order, user User) int {
	itemCount := len(order.items)
	if itemCount >= b.minBuy {
		return b.getFree * 100
	}
	return 0
}

package main

type PricingStrategy interface {
	Calculate(*Ticket) int
}

type TimeBasedPricingStrategy struct {
}

func (t *TimeBasedPricingStrategy) Calculate(ticket *Ticket) int {
	return 0
}

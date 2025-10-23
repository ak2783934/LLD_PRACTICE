package main

type PercentageExecutor struct{}

func (p *PercentageExecutor) Execute(paidBy int, participants []int, totalAmount int, metadata interface{}, manager *SplitwiseManager) {
	percentages := metadata.([]int)
	for idx, participant := range participants {
		if participant != paidBy {
			amount := (totalAmount * percentages[idx]) / 100
			manager.updateBalance(participant, paidBy, amount)
		}
	}
}

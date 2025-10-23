package main

type ExactExecutor struct{}

func (e *ExactExecutor) Execute(paidBy int, participants []int, totalAmount int, metadata interface{}, manager *SplitwiseManager) {
	shares := metadata.([]int)
	for idx, participant := range participants {
		if participant != paidBy {
			manager.updateBalance(participant, paidBy, shares[idx])
		}
	}
}

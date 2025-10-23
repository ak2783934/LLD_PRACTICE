package main

type EqualExecutor struct{}

func (e *EqualExecutor) Execute(paidBy int, participants []int, totalAmount int, metadata interface{}, manager *SplitwiseManager) {
	shareAmount := totalAmount / len(participants)
	for _, participant := range participants {
		if participant != paidBy {
			manager.updateBalance(participant, paidBy, shareAmount)
		}
	}
}

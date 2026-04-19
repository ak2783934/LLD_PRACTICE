package main

type ExpenseExecutor interface {
	ExecuteExpense(paidBy int, participants []int, totalAmount int, metadata map[string]string, splitwiseMangager *SplitwiseManager) error
}

type EqualExecutor struct {
}

func (e *EqualExecutor) ExecuteExpense(paidBy int, participants []int, totalAmount int, metadata map[string]string, splitwiseMangager *SplitwiseManager) {
	// split the balance amoung each other equally.
}

type PercentageExecutor struct {
}

func (p *PercentageExecutor) ExecuteExpense(paidBy int, participants []int, totalAmount int, metadata map[string]string, splitwiseMangager *SplitwiseManager) {
	// split the balance based on percentage from metadata.
}

type ExactExecutor struct {
}

func (p *ExactExecutor) ExecuteExpense(paidBy int, participants []int, totalAmount int, metadata map[string]string, splitwiseMangager *SplitwiseManager) {
	// split the balance based on percentage from metadata.
}

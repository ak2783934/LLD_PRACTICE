package main

type ExpenseExecutor interface {
	Execute(paidBy int, participants []int, totalAmount int, metadata interface{}, manager *SplitwiseManager)
}

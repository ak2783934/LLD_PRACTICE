package main

func main() {
	splitwiseManager := GetSplitwiseManager()
	splitwiseManager.AddUser("avinash")
	splitwiseManager.AddUser("Saurabh")
	splitwiseManager.AddUser("rajesh")
	splitwiseManager.AddUser("Aryan")
	splitwiseManager.ShowAllUsers()

	splitwiseManager.AddExpense(&EqualExecutor{}, 0, []int{0, 1, 2}, 30, nil)

	splitwiseManager.ShowAllBalances()
}

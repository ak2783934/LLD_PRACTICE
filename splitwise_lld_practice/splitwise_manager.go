package main

import (
	"fmt"
	"sync"
)

type SplitwiseManager struct {
	Users       map[int]*User
	Expenses    map[int]map[int]int
	userIDCount int
	mu          sync.Mutex
}

var (
	instance *SplitwiseManager
	once     sync.Once
)

func GetSplitwiseManager() *SplitwiseManager {
	once.Do(func() {
		instance = &SplitwiseManager{
			Users:       map[int]*User{},
			Expenses:    map[int]map[int]int{},
			userIDCount: 0,
		}
	})
	return instance
}

func (s *SplitwiseManager) AddUser(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user := &User{
		id:   s.userIDCount,
		name: name,
	}
	s.Users[s.userIDCount] = user
	s.userIDCount++
}

func (s *SplitwiseManager) updateBalance(from int, to int, amount int) {
	if from == to {
		return
	}
	// Always store debt from lower ID → higher ID
	minID, maxID := from, to
	sign := 1
	if from > to {
		minID, maxID = to, from
		sign = -1
	}

	if s.Expenses[minID] == nil {
		s.Expenses[minID] = make(map[int]int)
	}
	s.Expenses[minID][maxID] += amount * sign

	// Remove zero balance
	if s.Expenses[minID][maxID] == 0 {
		delete(s.Expenses[minID], maxID)
		if len(s.Expenses[minID]) == 0 {
			delete(s.Expenses, minID)
		}
	}
}

func (s *SplitwiseManager) AddExpense(executor ExpenseExecutor, paidBy int, participants []int, totalAmount int, metadata interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Users[paidBy]
	if !ok {
		return fmt.Errorf("paid by user doesn't exist")
	}
	for _, p := range participants {
		_, isParticipantExist := s.Users[p]
		if !isParticipantExist {
			return fmt.Errorf("participant doesn't exist ID: %v", p)
		}
	}

	executor.Execute(paidBy, participants, totalAmount, metadata, s)
	return nil
}

func (s *SplitwiseManager) ShowUserBalance(userID int) {
	total := 0

	_, ok := s.Users[userID]
	if !ok {
		fmt.Printf("User %d not found\n", userID)
		return
	}

	// Iterate all stored pairs
	for minID, inner := range s.Expenses {
		for maxID, val := range inner {
			if userID == minID {
				// stored val: min owes max (if val>0)
				total += val
			} else if userID == maxID {
				// stored val: min owes max, so max is owed => user receives => subtract
				total -= val
			}
		}
	}

	if total == 0 {
		fmt.Printf("User %d (%s) has no balance\n", userID, s.Users[userID].name)
	} else if total > 0 {
		fmt.Printf("User %d (%s) owes total %d\n", userID, s.Users[userID].name, total)
	} else {
		fmt.Printf("User %d (%s) should receive total %d\n", userID, s.Users[userID].name, -total)
	}
}

func (s *SplitwiseManager) ShowAllBalances() {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("----- All Balances -----")
	// for uid := range s.Users {
	// 	// reuse ShowUserBalance but avoid double locking; call internal printing
	// }
	// Instead of calling ShowUserBalance which locks, print all by computing totals here:
	totals := make(map[int]int)
	for uid := range s.Users {
		totals[uid] = 0
	}

	for minID, inner := range s.Expenses {
		for maxID, val := range inner {
			totals[minID] += val
			totals[maxID] -= val
		}
	}
	for uid, amt := range totals {
		if amt == 0 {
			fmt.Printf("User %d (%s) : No balance\n", uid, s.Users[uid].name)
		} else if amt > 0 {
			fmt.Printf("User %d (%s) : Owes %d\n", uid, s.Users[uid].name, amt)
		} else {
			fmt.Printf("User %d (%s) : Receives %d\n", uid, s.Users[uid].name, -amt)
		}
	}
	fmt.Println("------------------------")
}

func (s *SplitwiseManager) ShowAllUsers() {
	for _, user := range s.Users {
		fmt.Println("userID : ", user.id, " name : ", user.name)
	}
}

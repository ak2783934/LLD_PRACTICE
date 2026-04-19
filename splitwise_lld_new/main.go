package main

import (
	"errors"
	"sync"
)

type User struct {
	id    int
	name  string
	email string
}

type SplitwiseManager struct {
	Users     map[int]User
	userCount int
	mu        sync.Mutex
	Expenses  map[int]map[int]int // {1: {2: 20}, {3: 40}}, {2: }
}

/*
context of expenses table.

first is doner, second is taker.
and first should always be less, so that we don't have redundant data.

*/

func (s *SplitwiseManager) AddUser(name, email string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.userCount + 1
	s.userCount++
	s.Users[id] = User{
		id:    id,
		name:  name,
		email: email,
	}
}

func (s *SplitwiseManager) UpdateBalance(from, to, amount int) error {
	if from == to {
		return errors.New("to and from can't be same")
	}
	if from > to {
		amount = amount * -1
		from, to = to, from
	}

	// validate if the users exist or not?
	_, ok := s.Users[from]
	if !ok {
		return errors.New("User from doesn't exist")
	}
	_, ok = s.Users[to]
	if !ok {
		return errors.New("User to doesn't exist")
	}

	_, ok = s.Expenses[from]
	if !ok {
		s.Expenses[from] = make(map[int]int)
	}

	s.Expenses[from][to] += amount
	if s.Expenses[from][to] == 0 {
		delete(s.Expenses[from], to)
		if len(s.Expenses[from]) == 0 {
			delete(s.Expenses, from)
		}
	}
	return nil
}
func (s *SplitwiseManager) ShowBalance(userID int) (int, error) {
	_, ok := s.Users[userID]
	if !ok {
		return 0, errors.New("user not found")
	}

	// if found as from, add,
	// if found as to, decrease.
	total := 0
	for from := range s.Expenses {
		for to := range s.Expenses[from] {
			if from == userID {
				total += s.Expenses[from][to]
			}
			if to == userID {
				total -= s.Expenses[from][to]
			}
		}
	}
	return total, nil
}

func (s *SplitwiseManager) AddExpense(executor ExpenseExecutor, from int, participants []int, totalAmount int, metadata map[string]string) error {
	// validate if all users exist
	_, ok := s.Users[from]
	if !ok {
		return errors.New("from user doesn't exist")
	}

	for participant := range participants {
		_, ok = s.Users[participant]
		if !ok {
			return errors.New("participant does't exist")
		}
	}

	// call executor for executing the functions.
	err := executor.ExecuteExpense(from, participants, totalAmount, metadata, s)
	if err != nil {
		return err
	}
	return nil
}

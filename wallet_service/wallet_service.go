package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

const CurrencyINR = "INR"
const SuccessStatus = "SUCCESS"

type Wallet struct {
	walletID string
	balance  int
	userID   string
	currency string
	mu       sync.RWMutex
}

type TransactionData struct {
	transactionID string
	walletID      string
	amount        int
	status        string
	credit        bool
}

type WalletService struct {
	WalletIDMap        map[string]*Wallet // wallet id and wallet map
	UserIDWalletMap    map[string]*Wallet
	TransactionHistory []*TransactionData
}

func NewWalletService() *WalletService {
	return &WalletService{
		WalletIDMap:        make(map[string]*Wallet),
		UserIDWalletMap:    make(map[string]*Wallet),
		TransactionHistory: make([]*TransactionData, 0),
	}
}

func (w *WalletService) CreateWallet(userID string) (string, error) {
	// check if this user has wallet
	_, ok := w.UserIDWalletMap[userID]
	if ok {
		return "", errors.New(fmt.Sprintf("Wallet for user already exists: UserID: ", userID))
	}

	wallet := &Wallet{
		walletID: generate7CharUUID(),
		userID:   userID,
		balance:  0,
		currency: CurrencyINR,
	}

	// if no, create one and return walletID
	walletID := wallet.walletID
	w.UserIDWalletMap[userID] = wallet
	w.WalletIDMap[walletID] = wallet
	return walletID, nil
}

func (w *WalletService) Deposit(transactionID string, walletID string, amount int) (string, error) {
	wallet, ok := w.WalletIDMap[walletID]
	if !ok {
		return "", errors.New("wallet does not exist")
	}
	wallet.mu.Lock()
	defer wallet.mu.Unlock()
	wallet.balance += amount
	transaction := &TransactionData{
		transactionID: transactionID,
		walletID:      walletID,
		amount:        amount,
		credit:        true,
		status:        SuccessStatus,
	}
	w.TransactionHistory = append(w.TransactionHistory, transaction)
	return transaction.transactionID, nil
}

func (w *WalletService) Withdraw(transactionID string, walletID string, amount int) (string, error) {
	wallet, ok := w.WalletIDMap[walletID]
	if !ok {
		return "", errors.New("wallet does not exist")
	}
	wallet.mu.Lock()
	defer wallet.mu.Unlock()
	if wallet.balance < amount {
		return "", errors.New("Not enough balance in wallet")
	}

	wallet.balance -= amount
	transaction := &TransactionData{
		transactionID: transactionID,
		walletID:      walletID,
		amount:        amount,
		credit:        true,
		status:        SuccessStatus,
	}
	w.TransactionHistory = append(w.TransactionHistory, transaction)
	return transaction.transactionID, nil
}

func (w *WalletService) Transfer(fromWalletID string, toWalletID string, amount int) (string, error) {
	_, ok := w.WalletIDMap[fromWalletID]
	if !ok {
		return "", errors.New("sender wallet does not exist")
	}

	_, ok = w.WalletIDMap[toWalletID]
	if !ok {
		return "", errors.New("receiver wallet does not exist")
	}

	// atomic operation.
	transactionID := generate7CharUUID()

	_, err := w.Withdraw(transactionID, fromWalletID, amount)
	if err != nil {
		return "", err
	}

	w.Deposit(transactionID, toWalletID, amount)

	return transactionID, nil
}

func (w *WalletService) GetBalance(walletID string) (int, error) {
	wallet, ok := w.WalletIDMap[walletID]
	if !ok {
		return 0, errors.New("wallet does not exist")
	}
	wallet.mu.RLock()
	defer wallet.mu.RUnlock()
	return wallet.balance, nil
}

func generate7CharUUID() string {
	uuid := uuid.New().String()
	return uuid[0:7]
}

func main() {
	service := NewWalletService()

	// 1. Create wallets
	wallet1, err := service.CreateWallet("user1")
	if err != nil {
		fmt.Println(err)
		return
	}
	wallet2, err := service.CreateWallet("user2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created Wallets:", wallet1, wallet2)

	// 2. Deposit money
	txn1, _ := service.Deposit(generate7CharUUID(), wallet1, 1000)
	txn2, _ := service.Deposit(generate7CharUUID(), wallet2, 500)
	fmt.Println("Deposits done:", txn1, txn2)

	// 3. Get balances
	balance1, _ := service.GetBalance(wallet1)
	balance2, _ := service.GetBalance(wallet2)
	fmt.Println("Balances after deposit:", balance1, balance2)

	// 4. Withdraw money
	txn3, err := service.Withdraw(generate7CharUUID(), wallet1, 200)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	}
	fmt.Println("Withdraw txn:", txn3)

	// 5. Transfer money
	txn4, err := service.Transfer(wallet1, wallet2, 300)
	if err != nil {
		fmt.Println("Transfer error:", err)
	}
	fmt.Println("Transfer txn:", txn4)

	// 6. Balances after transfer
	balance1, _ = service.GetBalance(wallet1)
	balance2, _ = service.GetBalance(wallet2)
	fmt.Println("Balances after transfer:", balance1, balance2)

	// 7. Print transaction history
	fmt.Println("Transaction History:")
	for _, txn := range service.TransactionHistory {
		fmt.Printf("%+v\n", *txn)
	}
}

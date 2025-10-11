# Wallet Service - LLD Interview Question

## Scenario

You need to implement a basic wallet service where users can deposit, withdraw, and transfer money between wallets. The system must maintain transaction history (ledger) and be safe for concurrent operations.

---

## Requirements (Interview Scope)

### 1. Wallet

* Each wallet has a unique `walletID`, `userID`, and `balance`.
* Only supports one currency (e.g., USD) for simplicity.

### 2. Operations

* `CreateWallet(userID)` → returns `walletID`.
* `Deposit(walletID, amount)` → adds money to wallet.
* `Withdraw(walletID, amount)` → subtracts money; fail if insufficient balance.
* `Transfer(fromWalletID, toWalletID, amount)` → atomic debit + credit; fail if insufficient balance.
* `GetBalance(walletID)` → returns current balance.

### 3. Ledger

* Record every operation with `TransactionID`, `walletID(s)`, `type` (deposit, withdraw, transfer), `amount`, `status`.

### 4. Concurrency

* Multiple operations can occur on the same wallet simultaneously.
* Must ensure **thread-safe updates** to wallet balances.

### 5. Optional

* Print wallet balances and ledger for verification.
* Basic retry logic is optional.

---

## Scope Limits

* No database persistence (keep all data in memory).
* No multi-currency, no blockchain integration.
* Focus is **thread-safe operations, atomic transfer, and basic ledger**.
* Avoid over-engineering (e.g., no external queues or async workers).

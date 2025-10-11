# 🧠 Design Question: In-Memory Database with Basic Transactions

## Problem Statement

Design and implement an **in-memory database** that supports basic CRUD operations and simple transactional behavior.
The goal is to simulate a lightweight, Redis- or SQL-like system that can `INSERT`, `SELECT`, `UPDATE`, and `DELETE` records in memory, along with optional `BEGIN`, `COMMIT`, and `ROLLBACK` operations.

---

## ✅ Core Requirements

### 1. Data Model

* Each table can store multiple records.
* Each record can be represented as a key-value pair or a map of column name → value.
* Support multiple tables (e.g., `users`, `orders`).

### 2. Basic Operations

* `INSERT(table, recordID, recordData)` — insert a new record.
* `SELECT(table, recordID)` — fetch a record by ID.
* `UPDATE(table, recordID, updatedData)` — modify existing record.
* `DELETE(table, recordID)` — remove a record.

### 3. Transactions (Optional but Preferred)

* `BEGIN_TRANSACTION()` — start a transaction session.
* `COMMIT(transactionID)` — permanently apply all changes made in the transaction.
* `ROLLBACK(transactionID)` — discard all changes made in the transaction.

### 4. Thread Safety

* The design should be safe for concurrent reads/writes (use locks where necessary).

### 5. Optional Add-on (Bonus)

* Support simple **indexing** to make lookups by a non-primary field faster (e.g., fetch user by email).
* Auto-cleanup or expiration mechanism for records (like TTL).

---

## 💡 Expectations

You are expected to:

* Design modular structs/classes (like `Database`, `Table`, `Transaction`).
* Implement a **working code** that demonstrates at least:

  * CRUD operations
  * Commit/Rollback working correctly
* Keep it in-memory only (no persistence required).

---

## 🕐 Time Constraint

You should be able to:

* Explain your design in ~15–20 minutes.
* Write a **working implementation** and run test cases in ~40–45 minutes.


Entities that might be required for this:

Row{
    id string
    value map[string]string
}

Table{
    Rows []*Row
}

Database{
    tables map[string]*Table
}

Note: -> they should support passing the key and value, becuase in where clause we can pass both and it works for table. 
INSERT("tableName", map[string]string) -> it will generate one id for this entry. -> assuming that schema will remain same. 
SELECT("tableName", map[string]string) -> returns array, matching all response. Input can be many filters. 
UPDATE("tableName", key, value, map[sring]string) -> update the value after searching using key and value. 
DELETE("tableName", map[string]string) -> searches all the keys using the map provided and delete the values if matches. 

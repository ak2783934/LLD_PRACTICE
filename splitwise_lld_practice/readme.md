# 🧾 Splitwise System Design (LLD) - Problem Statement

## 🎯 Objective
Design and implement a simplified version of the **Splitwise** application that helps users track group expenses and balances between friends.

---

## 📌 Requirements

### 1. Users
- The system should support **adding users** with basic information (e.g., user ID, name, email or phone).

### 2. Add Expense
When a user adds an expense, the following details must be provided:
- **Amount**
- **Paid by** (single user who paid the full amount)
- **Split among** (list of users who should share the expense)
- **Split type**, which can be:
  - **EQUAL** – Split equally among users.
  - **EXACT** – Exact amount owed by each user is provided.
  - **PERCENT** – Percentage of the total owed by each user is provided.

### 3. Balances
The system should track **net balances between users**, showing who owes whom and how much.

- If `User A owes User B ₹50`, the system should store only this balance.
- Balances should not be duplicated (i.e., `User B owes User A -₹50` should not be stored).

---

## 📊 Functionality to Support

| Feature                 | Description |
|-------------------------|-------------|
| Add User               | Add a new user into the system. |
| Add Expense            | Record an expense with split logic. |
| Show User Balance      | Show how much a single user owes or is owed. |
| Show All Balances      | Show all balances for all users in the system. |

---

## 🔄 Split Methods Example

### ✅ Equal Split Example
- Total: ₹100
- Paid by: A
- Split among: A, B, C, D
- Each person owes ₹25

### ✅ Exact Split Example
- Total: ₹100
- Paid by: A
- Split among: A (0), B (20), C (30), D (50)

### ✅ Percent Split Example
- Total: ₹100
- Paid by: A
- Split among percentages: B(20%), C(30%), D(50%)

---

## 🎯 Goal of the Exercise
You must:
1. **Identify entities** (like User, Expense, Balance, etc.)
2. **Define interactions/flows**
3. **Design class structure**
4. **Implement core functionality** to handle splits and balances

---

## 💬 Next Step
Proceed like a system design interview:
- Start by asking **clarifying questions**
- Then identify **entities and relationships**

> Ask your first clarifying question when ready.




My own analysis
- Questions: Do we need to simplify the distribution as well? 
- If yes, how do we manage the split between different group of people who are not even directly related, ideally we should not do any split among them. 
- Assuming that we don't have the concept of groups as of now. 
- 


How do we settle these expenses:

Three methods to only add the expense. When we add, we should do simple logical calcuations, but during final balance calculation we should have the logic. 

Not clear about the logic for balance. 

think from graphs perspective?

A->B(money), lets not make a adjacency list, make a matrix graph, mat[a][b]=x means a owes b x amount.
lets consider the ids to be auto increment always. this will help us to sort the things better. Always lower pays to higher something like this?

And for each update, we will update that money value only and consider it a one directional array for always? 
basically based on some id? 
if payment is other way? do the subtraction, or else do the addition? 



Entities: 
```
User{
    name
    id 
}

Expense{
    id 
    participants []userID
    expenseMethod 
    paidBy userID
}
This will have three implementations. 

What does adding user does? it just puts it in our db. No entry is registered in expense table. 


```


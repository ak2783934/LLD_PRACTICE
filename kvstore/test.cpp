#include <iostream>
#include <vector>
#include <unordered_map>
#include <string>
#include <algorithm>

using namespace std;

struct Transaction {
   string id;
   long timestamp;     // Unix timestamp
   double amount;
   string type;   // "credit" or "debit"
   string status; // "success" or "failed"
};


bool isSelectedTransaction(const Transaction& transaction, const unordered_map<string, string>& filters) {
    for (const auto& [key, value] : filters) {
        if (key == "type" && transaction.type != value)
            return false;
        if (key == "status" && transaction.status != value)
            return false;
        if (key == "min_amount" && transaction.amount < stod(value))
            return false;
        if (key == "max_amount" && transaction.amount > stod(value))
            return false;
        if (key == "start_time" && transaction.timestamp < stol(value))
            return false;
        if (key == "end_time" && transaction.timestamp > stol(value))
            return false;
    }
    return true;
}


vector<Transaction> searchTransactions(
   const vector<Transaction>& transactions,
   const unordered_map<string, string>& filters,
   int limit,
   int offset
) {
    
    vector<Transaction>selectedTransactions;
    for(int i=0;i<transactions.size();i++){
        if(isSelectedTransaction(transactions[i], filters)){
            selectedTransactions.push_back(transactions[i]);
        }
    }

    
    return selectedTransactions;
}

void printTransactions(const vector<Transaction>& transactions) {
   for (const auto& tx : transactions) {
       cout << "ID: " << tx.id
                 << ", Timestamp: " << tx.timestamp
                 << ", Amount: " << tx.amount
                 << ", Type: " << tx.type
                 << ", Status: " << tx.status << endl;
   }
}

int main() {
   vector<Transaction> transactions = {
       {"tx1", 1719400000, 120.5, "credit", "success"},
       {"tx2", 1719403600, 75.0, "debit", "failed"},
       {"tx3", 1719407200, 200.0, "credit", "success"},
       {"tx4", 1719410800, 45.0, "debit", "success"},
       {"tx5", 1719414400, 300.0, "credit", "failed"}
   };

   unordered_map<string, string> filters = {
       {"type", "credit"},
       {"min_amount", "100"},
       {"status", "success"}
   };

   int limit = 2;
   int offset = 0;

   vector<Transaction> result = searchTransactions(transactions, filters, limit, offset);

   cout << "Filtered Transactions:" << endl;
   printTransactions(result);

   return 0;
}
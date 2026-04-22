1. Given a stream of transactions, write a function that aggregates them by merchant_id and returns the total amount per merchant in the last N minutes (moving window).



lets do bucketting at minute level? 
for each merchant id seperately? 


type Transaction struct {
    txn_id string
    merchant_id string 
    amount int
    timestamp time.Time
}

func processTxn(txn <-chan Transaction){
    switch txn{
        case <- done():

    }
}
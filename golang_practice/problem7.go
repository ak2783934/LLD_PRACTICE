package main

/*
🔴 Exercise 8: Avoid goroutine leak

Problem
Fix this leak:

func worker(ch <-chan int) {
    for {
        v := <-ch
        fmt.Println(v)
    }
}

we should be using range here, becuase when channel is closed, we won't know here that cannel is closed. 
or use context with switch

*/

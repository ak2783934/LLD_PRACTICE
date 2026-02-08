package main

/*

🟡 Exercise 7: select with timeout

Problem
Receive from a channel, but timeout after 1 second.

select {
case v := <-ch:
    fmt.Println(v)
case <-time.After(time.Second):
    fmt.Println("timeout")
}


Interviewer checks:

Do you know select?

Do you avoid blocking forever?


*/

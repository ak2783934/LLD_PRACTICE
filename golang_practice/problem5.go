package main

/*

🟡 Exercise 6: Fan-out / Fan-in

Problem

One goroutine generates numbers 1–10

- Multiple workers square them
- Collect results into one channel

Print results

Concepts:

Fan-out (jobs → workers)

Fan-in (results ← workers)

WaitGroup + close

*/

// func generateNumbers(numCh chan<- int, wg *sync.WaitGroup) {
// 	for i := 0; i < 20; i++ {
// 		numCh <- i
// 	}
// 	close(numCh)
// }

// // numCh for consuming
// // sqNumCh for publishing
// func generateSquareWorker(numCh <-chan int, sqNumCh chan<- int, workerID int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range numCh {
// 		sqNumCh <- num * num
// 	}
// }

// func main() {
// 	numCh := make(chan int, 10)
// 	sqNumCh := make(chan int, 10)
// 	var wg sync.WaitGroup

// 	// go routine to generate numbers
// 	wg.Add(4)

// 	// run 3 worker, consuming parallelly and publish the sqaures.
// 	for i := 0; i < 4; i++ {
// 		go generateSquareWorker(numCh, sqNumCh, i, &wg)
// 	}

// 	go generateNumbers(numCh, &wg)

// 	// run a for loop to listen to that topic and print the answers?

// 	go func() {
// 		wg.Wait()
// 		close(sqNumCh)
// 	}()

// 	for sq := range sqNumCh {
// 		fmt.Println(sq)
// 	}
// }

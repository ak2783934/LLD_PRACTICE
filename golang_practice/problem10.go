package main

/*
Problem 2: Ordered fan-in

(Harder than it looks)

Task

Same as Problem 1

Output must be printed in order (1², 2², 3²...)

Hints

Workers run concurrently

You must preserve order without blocking everything

Concepts tested
✔ Ordering guarantees
✔ Concurrency coordination
✔ Data structures + channels
*/

// func generateNumbers(numCh chan<- int) {
// 	for i := 0; i < 20; i++ {
// 		numCh <- i
// 	}
// 	close(numCh)
// }

// // numCh for consuming
// // sqNumCh for publishing
// func generateSquareWorker(numCh <-chan int, sqNumCh chan<- SqNumWithIndex, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range numCh {
// 		sqNumCh <- SqNumWithIndex{index: num, val: num * num}
// 	}
// }

// type SqNumWithIndex struct {
// 	index int
// 	val   int
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()
// 	numCh := make(chan int, 10)
// 	sqNumWithIndex := make(chan SqNumWithIndex, 10)
// 	var wg sync.WaitGroup

// 	// go routine to generate numbers
// 	wg.Add(4)

// 	// run 3 worker, consuming parallelly and publish the sqaures.
// 	for i := 0; i < 4; i++ {
// 		go generateSquareWorker(numCh, sqNumWithIndex, &wg)
// 	}

// 	go generateNumbers(numCh)

// 	// run a for loop to listen to that topic and print the answers?

// 	go func() {
// 		wg.Wait()
// 		close(sqNumWithIndex)
// 	}()

// 	// storageMap := map[int]int{}
// 	// nextExpected := 0
// 	// for sqWithIndex := range sqNumWithIndex {
// 	// 	index := sqWithIndex.index
// 	// 	val := sqWithIndex.val

// 	// 	storageMap[index] = val

// 	// 	for {
// 	// 		val, ok := storageMap[nextExpected]
// 	// 		if ok {
// 	// 			fmt.Println(val)
// 	// 			delete(storageMap, nextExpected)
// 	// 			nextExpected++
// 	// 		} else {
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }

// 	for {
// 		select {
// 		case sqWithIndex, ok := <-sqNumWithIndex:
// 			if !ok {
// 				return
// 			}
// 			time.Sleep(500 * time.Millisecond)
// 			val := sqWithIndex.val
// 			fmt.Println(val)
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }

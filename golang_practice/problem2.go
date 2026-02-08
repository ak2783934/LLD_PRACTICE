package main

/*

Print numbers 1–5 in order using goroutines


*/

// func consume(val <-chan int, wg *sync.WaitGroup) {
// 	for v := range val {
// 		fmt.Println(v)
// 		wg.Done()
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	intChan := make(chan int)

// 	wg.Add(5)
// 	go consume(intChan, &wg)
// 	for i := 0; i < 5; i++ {
// 		intChan <- i
// 	}
// 	close(intChan)
// 	wg.Wait()
// }

/*
Another good solution

func consume(ch <-chan int, done chan struct{}) {
	for v := range ch {
		fmt.Println(v)
	}
	close(done)
}

func main() {
	ch := make(chan int, 5)
	done := make(chan struct{})

	go consume(ch, done)

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	<-done
}

*/

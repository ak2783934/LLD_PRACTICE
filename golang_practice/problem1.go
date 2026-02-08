package main

import (
	"fmt"
	"sync"
)

/*

Print numbers 1 to 5 using goroutines, but ensure main waits for all goroutines.

*/

func printNumber(i int) {
	wg.done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Println(x)
			wg.Done()
		}(i)
	}
	fmt.Println("test")
	wg.Wait()
}

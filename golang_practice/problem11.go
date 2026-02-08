package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func main() {
	c := &Counter{}
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				ch <- 1
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		c.count += x
	}
	fmt.Println(c.count)
}

package main

import "fmt"

func main() {
	lruCache := CreateLRUCache(5)

	lruCache.set(5, 10)
	lruCache.Queue.PrintList()
	lruCache.set(23, 132)

	lruCache.set(231, 2345)
	lruCache.set(2342, 979)

	lruCache.set(231, 234234)
	lruCache.set(3412345, 979)
	lruCache.set(231, 2345)
	fmt.Println(lruCache.get(23))
	lruCache.set(2341234, 979)
	lruCache.set(231, 2345)
	lruCache.set(2342, 979)

	lruCache.Queue.PrintList()

	// fmt.Println(lruCache.get(234324))

	// lruCache.Queue.PrintList()
}

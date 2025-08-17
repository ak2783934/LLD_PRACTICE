package main

func main() {
	lruCache := NewLRUCache(5)

	lruCache.SetValue(5, 5)

	lruCache.SetValue(3, 2)

	lruCache.GetValue(5)

	lruCache.SetValue(1, 1)
	lruCache.SetValue(2, 2)
	lruCache.SetValue(0, 2)
	lruCache.SetValue(7, 2)
	lruCache.SetValue(11, 5)

	lruCache.PrintLRUCache()
}

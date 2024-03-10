package main

import "fmt"

func main() {
	// generate a new cache
	cache := NewCache(5, NewLRUReplacementPolicy())

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	cache.ViewCache()

	// get value from cache
	val, _ := cache.Get("key1")
	fmt.Println("value is ", val)

	cache.ViewCache()

	// put value in cache
	cache.Put("key3", "value3")
	cache.Put("key4", "value4")
	cache.Put("key5", "value5")
	cache.Put("key6", "value6")

	cache.ViewCache()
}

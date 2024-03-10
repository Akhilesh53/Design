package main

// IReplacementPolicy defines the interface for a cache replacement policy.
type IReplacementPolicy interface {
	// Evict is called when the cache needs to evict an item.
	Evict(c *Cache)

	// Get is called when the cache needs to retrieve an item.
	Get(c *Cache, key string) (string, error)

	// Put is called when the cache needs to store a new item.
	Put(c *Cache, key string, value string)

	// View the data associated with the replacement policy.
	View()
}

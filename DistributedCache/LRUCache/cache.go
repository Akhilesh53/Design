package main

type Cache struct {
	maxSize       int
	replacePolicy IReplacementPolicy
	cache         map[string]*Node
}

func NewCache(maxSize int, replacePolicy IReplacementPolicy) *Cache {
	return &Cache{
		maxSize:       maxSize,
		replacePolicy: replacePolicy,
		cache:         make(map[string]*Node),
	}
}

// write al getter setter for cahce
func (c *Cache) GetMaxSize() int {
	return c.maxSize
}

func (c *Cache) SetMaxSize(maxSize int) {
	c.maxSize = maxSize
}

func (c *Cache) GetReplacePolicy() IReplacementPolicy {
	return c.replacePolicy
}

func (c *Cache) SetReplacePolicy(replacePolicy IReplacementPolicy) {
	c.replacePolicy = replacePolicy
}

func (c *Cache) GetCache() map[string]*Node {
	return c.cache
}

func (c *Cache) SetCache(cache map[string]*Node) {
	c.cache = cache
}

func (c *Cache) SetValueInCache(key string, node *Node) {
	c.cache[key] = node
}

// get key from the cache, if present
func (c *Cache) Get(key string) (string, error) {
	return c.GetReplacePolicy().Get(c, key)
}

// put ndoe in the cache
func (c *Cache) Put(key string, value string) {
	c.GetReplacePolicy().Put(c, key, value)
}

// view dll
func (c *Cache) ViewCache() {
	c.GetReplacePolicy().View()
}

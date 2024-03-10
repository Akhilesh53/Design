package main

import "errors"

// LRUReplacementPolicy represents the Least Recently Used (LRU) replacement policy
// for a cache. It keeps track of the most recently used items using a doubly linked list.
type LRUReplacementPolicy struct {
	ll *DoublyLinkedList
}

// NewLRUReplacementPolicy creates a new LRUReplacementPolicy instance.
// It initializes and returns a pointer to an LRUReplacementPolicy struct,
// with an empty doubly linked list.
func NewLRUReplacementPolicy() *LRUReplacementPolicy {
	return &LRUReplacementPolicy{
		ll: NewDoublyLinkedList(),
	}
}

// GetLL returns the doubly linked list associated with the LRU replacement policy.
func (lru *LRUReplacementPolicy) GetLL() *DoublyLinkedList {
	return lru.ll
}

// MoveToLast moves the given node to the last position in the LRU cache.
// It returns the updated node after the move.
func (lru *LRUReplacementPolicy) MoveToLast(node *Node) *Node {
	return lru.GetLL().MoveToLast(node)
}

func (lru *LRUReplacementPolicy) Evict(c *Cache) {
	// if tail == head
	if lru.GetLL().GetTail() == lru.GetLL().GetHead() {
		node := lru.GetLL().GetTail()
		lru.GetLL().SetHead(nil)
		lru.GetLL().SetTail(nil)
		delete(c.GetCache(), node.GetKey())
		return
	}

	// Evict the least recently used item
	delNode := lru.GetLL().GetHead()
	lru.GetLL().SetHead(delNode.GetNext()).SetPrev(nil)
	delete(c.GetCache(), delNode.GetKey())
}

// Get retrieves the value associated with the given key from the cache.
// If the key is already in the cache, it moves the corresponding node to the last position in the cache.
// If the key is not in the cache, it does nothing.
func (lru *LRUReplacementPolicy) Get(c *Cache, key string) (string, error) {
	// if key is already in the cache
	if node, ok := c.GetCache()[key]; ok {
		node = lru.MoveToLast(node)
		c.SetValueInCache(key, node)
		return c.GetCache()[key].GetVal(), nil
	}
	return "", errors.New("key not found")
}

// Put adds a new key-value pair to the cache or updates the value if the key already exists.
// If the key already exists, the corresponding node is moved to the last position in the cache.
// If the cache is full, the least recently used node is evicted before adding the new node.
// Parameters:
//   - c: The cache object.
//   - key: The key of the new node.
//   - value: The value of the new node.
func (lru *LRUReplacementPolicy) Put(c *Cache, key string, value string) {

	// if value is already in the cache
	if node, ok := c.GetCache()[key]; ok {
		node.SetVal(value)
		node = lru.MoveToLast(node)
		c.SetValueInCache(key, node)
		return
	}

	// if cache is full
	if len(c.GetCache()) >= c.GetMaxSize() {
		lru.Evict(c)
	}

	// add new node to the cache
	node := NewNode(key, value)
	c.SetValueInCache(key, node)
	lru.GetLL().Add(node)
}

func (lru *LRUReplacementPolicy) View() {
	// print the doubly linked list
	lru.GetLL().PrintLL()
}

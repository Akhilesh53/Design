package main

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	key := "testKey"
	val := "testValue"
	node := NewNode(key, val)

	if node.key != key {
		t.Errorf("Expected key to be %s, got %s", key, node.key)
	}

	if node.val != val {
		t.Errorf("Expected value to be %s, got %s", val, node.val)
	}

	if node.next != nil {
		t.Errorf("Expected next to be nil, got %v", node.next)
	}

	if node.prev != nil {
		t.Errorf("Expected prev to be nil, got %v", node.prev)
	}
}

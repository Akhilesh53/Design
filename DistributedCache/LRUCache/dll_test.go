package main

import (
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()

	if list.head != nil {
		t.Errorf("Expected head to be nil, got %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("Expected tail to be nil, got %v", list.tail)
	}
}

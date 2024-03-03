package main

import (
	"fmt"
	"time"
)

func main() {
	window := NewSlidingWindowCounter(2, 1)

	fmt.Println(window.IsAllowed(NewPacket(1, 100))) //true
	fmt.Println(window.IsAllowed(NewPacket(2, 100))) //true
	fmt.Println(window.IsAllowed(NewPacket(3, 100))) //true
	time.Sleep(1 * time.Second)
	fmt.Println(window.IsAllowed(NewPacket(4, 100))) //true
	fmt.Println(window.IsAllowed(NewPacket(5, 100))) //true
	fmt.Println(window.IsAllowed(NewPacket(6, 100))) //true
	fmt.Println(window.IsAllowed(NewPacket(7, 100))) //false
}

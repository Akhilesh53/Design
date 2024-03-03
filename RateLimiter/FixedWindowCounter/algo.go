package main

import (
	"fmt"
	"time"
)

func main() {
	window := NewFixedWindowCounter(5, 1)

	fmt.Println(window.IsAllowed()) //true
	fmt.Println(window.IsAllowed()) //true
	fmt.Println(window.IsAllowed()) //true
	time.Sleep(1 * time.Second)
	fmt.Println(window.IsAllowed()) //true
	fmt.Println(window.IsAllowed()) //true
	fmt.Println(window.IsAllowed()) //true
	fmt.Println(window.IsAllowed()) //false
}

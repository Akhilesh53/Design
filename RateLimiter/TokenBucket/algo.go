package main

import (
	"fmt"
	"time"
)

func main() {
	bucket := NewBucket(3, 1*time.Minute)

	fmt.Println(bucket.HaveToken(1)) //true
	fmt.Println(bucket.HaveToken(1)) //true
	fmt.Println(bucket.HaveToken(1)) //true
	fmt.Println(bucket.HaveToken(1)) //false
}

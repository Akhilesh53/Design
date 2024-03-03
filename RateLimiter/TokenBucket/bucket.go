package main

import "time"

// if we have to make bucket related to a particular client,
// we can add a client id to the struct

type Bucket struct {
	capacity int
	tokens   int
	rate     time.Duration
	lastTime time.Time
}

// make a new bucket
func NewBucket(capacity int, rate time.Duration) *Bucket {
	return &Bucket{
		capacity: capacity,
		tokens:   capacity,
		rate:     rate,
		lastTime: time.Now(),
	}
}

// get a token from the bucket
func (b *Bucket) HaveToken(count int) bool {
	b.refill()
	if count > b.tokens {
		return false
	}
	b.tokens -= count
	return true
}

// write refil func
// if time elapsed exceeds the rate, refill the bucket
// o refill can be a cron job to fill be token after every rate time
func (b *Bucket) refill() {
	now := time.Now()
	elapsed := now.Sub(b.lastTime).Minutes()
	if elapsed < b.rate.Minutes() {
		return
	}
	b.lastTime = now
	b.tokens = b.capacity
}

func (b *Bucket) Capacity() int {
	return b.capacity
}

func (b *Bucket) Tokens() int {
	return b.tokens
}

func (b *Bucket) Rate() time.Duration {
	return b.rate
}

func (b *Bucket) SetRate(rate time.Duration) *Bucket {
	b.rate = rate
	return b
}

func (b *Bucket) SetCapacity(capacity int) *Bucket {
	b.capacity = capacity
	return b
}

func (b *Bucket) SetTokens(tokens int) *Bucket {
	b.tokens = tokens
	return b
}

package main

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	maxCapacity int
	buffer      []*Packet
	rate        int
}

func NewLeakyBucket(maxCapacity, rate int) *LeakyBucket {
	return &LeakyBucket{
		maxCapacity: maxCapacity,
		buffer:      make([]*Packet, 0),
		rate:        rate,
	}
}

// add packet
func (b *LeakyBucket) AddPacket(p *Packet) {
	if len(b.buffer) >= b.MaxCapacity() {
		fmt.Println("Buffer is full, packet id :: ", p.PacketId())
		return
	}
	b.buffer = append(b.buffer, p)
}

// get max capactiy
func (b *LeakyBucket) MaxCapacity() int {
	return b.maxCapacity
}

// transmit packet
func (b *LeakyBucket) retrieveRequests() {
	if len(b.buffer) == 0 {
		fmt.Println("Buffer is empty, no packet to transmit...")
		return
	}

	rate := b.Rate()
	for len(b.buffer) > 0 {
		packet := b.buffer[0]
		if packet.Size() > rate {
			fmt.Println("Packet size is greater than rate, cannot transmit...")
			break
		}
		fmt.Println("Transmitting packet with id: ", packet.PacketId())
		rate -= packet.Size()
		b.buffer = b.buffer[1:]
	}
}

// get rate
func (b *LeakyBucket) Rate() int {
	return b.rate
}

func (b *LeakyBucket) TransmitPacket() {
	for {
		b.retrieveRequests()
		time.Sleep(5 * time.Second)
	}
}

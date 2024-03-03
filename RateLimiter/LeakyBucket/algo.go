package main

func main() {
	bucket := NewLeakyBucket(5, 1000)
	bucket.AddPacket(NewPacket(1, 800))
	bucket.AddPacket(NewPacket(2, 700))
	bucket.AddPacket(NewPacket(3, 600))
	bucket.AddPacket(NewPacket(4, 500))
	bucket.AddPacket(NewPacket(5, 400))
	bucket.AddPacket(NewPacket(6, 300))
	bucket.AddPacket(NewPacket(7, 200))
	bucket.AddPacket(NewPacket(8, 100))
	bucket.TransmitPacket()
}

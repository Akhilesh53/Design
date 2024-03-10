package main

type IHash interface {
	// Hash returns the hash of the given key.
	Hash(key string) uint32
}

type ModHash struct {
	size uint32
}

// NewModHash creates a new ModHash.
func NewModHash(size uint32) *ModHash {
	return &ModHash{
		size: size,
	}
}

// Hash returns the hash of the given key.
func (m *ModHash) Hash(key string) uint32 {
	return uint32(len(key)) % uint32(m.size)
}

type RandomHash struct{}

func (rh *RandomHash) Hash(key string) uint32 {
	hash := uint32(7)
	for i := 0; i < len(key); i++ {
		hash = hash*31 + uint32(key[i])
	}
	return hash
}

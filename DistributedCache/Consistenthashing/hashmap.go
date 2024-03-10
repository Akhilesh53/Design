package main

type ConsistentHashMap struct {
	hashring map[uint32]*Server // sorted set
	replicas int                // to reduce the possibility of domino effect
	hashfunc IHash
}

// NewConsistentHashMap creates a new ConsistentHashMap.
func NewConsistentHashMap(replicas int, hashfunc IHash) *ConsistentHashMap {
	return &ConsistentHashMap{
		hashring: make(map[uint32]*Server),
		replicas: replicas,
		hashfunc: hashfunc,
	}
}

// wrte getter setter for all params
func (c *ConsistentHashMap) GetHashRing() map[uint32]*Server {
	return c.hashring
}

func (c *ConsistentHashMap) SetHashRing(hashring map[uint32]*Server) {
	c.hashring = hashring
}

func (c *ConsistentHashMap) SetInHashRing(key uint32, server *Server) {
	c.hashring[key] = server
}

func (c *ConsistentHashMap) GetFromHashRing(key uint32) *Server {
	return c.hashring[key]
}

func (c *ConsistentHashMap) GetReplicas() int {
	return c.replicas
}

func (c *ConsistentHashMap) SetReplicas(replicas int) {
	c.replicas = replicas
}

func (c *ConsistentHashMap) GetHashFunc() IHash {
	return c.hashfunc
}

func (c *ConsistentHashMap) SetHashFunc(hashfunc IHash) {
	c.hashfunc = hashfunc
}

// AddServer adds a new server to the hash ring.
func (c *ConsistentHashMap) AddServer(server *Server) {
	for i := 0; i < c.replicas; i++ {
		hash := c.GetHashFunc().Hash(server.GetIP() + string(rune(i)))
		c.hashring[hash] = server
	}
}

// RemoveServer removes a server from the hash ring.
func (c *ConsistentHashMap) RemoveServer(server *Server) {
	for i := 0; i < c.replicas; i++ {
		hash := c.GetHashFunc().Hash(server.GetIP() + string(rune(i)))
		serverToeDeleted := c.GetFromHashRing(hash)

		// reindex all keys after the removed server
		for _, entry := range serverToeDeleted.GetEntries() {
			c.put(entry.GetKey(), entry.GetValue())
		}
	}
}

// todo:  this func  can be modified to (n log n) considerign hash ring to be sorted map/ sorted set
// use binay search
// considering using this lib: https://github.com/umpc/go-sortedmap once i go through this, will integrate the same.
func (c *ConsistentHashMap) GetServer(key string) *Server {
	hash := c.GetHashFunc().Hash(key)
	for k, v := range c.hashring {
		if k >= hash {
			return v
		}
	}
	return c.hashring[0]
}

func (c *ConsistentHashMap) put(key string, value string) {
	server := c.GetServer(key)
	server.AddEntry(key, value)
}

func (c *ConsistentHashMap) get(key string) (string, error) {
	server := c.GetServer(key)
	return server.GetEntry(key)
}

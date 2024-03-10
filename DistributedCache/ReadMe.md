Design a distributed cache system
-------------------------------------

A distributed cache is basically a distributed hash table where we need to store key value objects.

Functional
Get(key)
Put(key, value )

Non-Funcational
Available
Scalable

- To store the data in hash table, we need hash function.
- Cache can be dedicated cluster or co-located cluster
  In co-located cluster, cache service and cache db are inplace - cannot store more amount of data
  In dedicated clsuter, cache db and service are independent and specify their own configuration.
- In a single hash table it is not easy to store large amount of data. Thus we need to shard the data.
- To shard the data across different clusters, hash func can be used. Limitation is if some node is down or we need to add some new host, we need to shuffle all keys which is again an overhead. To reduce the amount of key shuffling, we can use consistent hashing.
- If the size of  hash table reaches it max, we need to evict the data to make some free space. Thus, need of eviction strategies (LRU, LFU, FIFO)
- To maintain a list of configuration/ servers - which servers are part of distributed cache
  - each server has a file containeing list of servers - manually need to update file after certain time
  - Use configuration management service like zookeeper to maintain list of servers.

- Shards can become hot means some shards are processing more requests than others  
  To reduce this problem, we can introduce data replication strategy (easiest would be master slave)
  Write will be only on master
  Read will be distributed across master and slave node

- Other points:
  - Consistency: we are preferening availability and scalability over consistency. In async replication, thre might be case of data inconsistency. Soln to this is sync replication which negatively impacts performance by increasing latency.
  - Data Expiration: Eviction poilicies only removes data when cache is full, but there might be some keys which are present from such a long time. We need to introduce a time to live (TTL) mechanism to remove all the key which are expired after certain interval of time. There are two approaches for this:
    1) Which get  the data, we check if the lastupdated time of key is expired, return null result
    2) A separate thread will run over the cache and removes all expired keys. But traversilng over large number of keys will be a performance overhead. Thus uses probabilitic algo to remove the expired keys (random selection of keys)  
  - Local remote cache: we can also a maintain a local remote cache which will first check in local cache, if not found then call is made to distributed cache.
  - Security: Enalble firewall proxies to make connection to cache system.
  - Monitoring and logging

- Consistent hashing has 2 major flaws
  - Domino effect (a chain reaction of failures. If one node fails, all its load shifted to next server. due to load that might also gets failed). One possible soln is to add multiple number of entries of same server in the hash ring
  - hash ring is not splitted evenly. Modified versions of consistent hashing is in the market.

- Data Replication Protocol
  1) Probabilistic data repication protocol
     - Gossip, Epidemic Broadcast Trees, Bimodal Multicast (Provides Eventual Consistency)
  2) Consensus Protocols
     -  2 Phase/ 3 Phase commit, Raft, Paxos, Chain replication (Provides strong consistency)
  3) Master Slave Architecture

To Read:
- Redis Senetial
- Google Jump Hash Algorithm
- Yahoo proportional algorithm
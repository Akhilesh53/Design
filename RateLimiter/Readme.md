Design a Rate limiter
----------------------------

1) Rate limiter limits the number of hits to a url
2) If the number of hits exceed threshold, itwill throttle the system (throw an error that request cnnot be served)
3) All the requests with in threshold should be served properly

Types of Rate limiting
-----------------------
1) User level
2) Server level
3) Geographical level

Common Aglorithms
------------------

1) Token Bucket Algorithm
     Req -> If tokens: 
                - Yes : Process request
                - No : Discard request

            After a certain period of time, token bucket gets filled with particular number of tokens


2) Leaky Bucket: In this we have a bucket and a queue of fixed size. If bucket is full, all requests will be discarded.
                 From bucket all requests will be moved to queue from where it gets processed at a fixed rate.

3) Fixed Window Counter
   In this we divide the time into fixed slots/ time frame size. All the requests that came in particular slot, a counter for that slot will be increased. If counter exceeds threshold, requests will be discarded.
   Same goes for all other fixed time frame size windows.


4) Sliding Window Counter
   In this, what happens is rather than fixed window slots, we have sliding window slots. for each slot we have a counter. If counter exceeds threshold, request will be discarded.
   


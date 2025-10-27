# ZDIFF
This command is similar to ZDIFFSTORE, but instead of storing the resulting sorted set, it is returned to the client.

## Syntax
```
ZDIFF numkeys key [key ...] [WITHSCORES]
```

## Time Complexity
O(L + (N-K)log(N)) worst case where L is the total number of elements in all the sets, N is the size of the first set, and K is the size of the result set.

## Return Value
**Array reply**: the result of the difference including, optionally, scores when the WITHSCORES option is used.

## Examples
```bash
redis> ZADD zset1 1 "one"
(integer) 1
redis> ZADD zset1 2 "two"
(integer) 1
redis> ZADD zset1 3 "three"
(integer) 1
redis> ZADD zset2 1 "one"
(integer) 1
redis> ZADD zset2 2 "two"
(integer) 1
redis> ZDIFF 2 zset1 zset2
1) "three"
redis> ZDIFF 2 zset1 zset2 WITHSCORES
1) "three"
2) "3"
```
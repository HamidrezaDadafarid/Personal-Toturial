# ZDIFFSTORE
Computes the difference between the first and all successive input sorted sets and stores the result in destination. The total number of input keys is specified by numkeys.

Keys that do not exist are considered to be empty sets.

If destination already exists, it is overwritten.

## Syntax
```
ZDIFFSTORE destination numkeys key [key ...]
```

## Time Complexity
O(L + (N-K)log(N)) worst case where L is the total number of elements in all the sets, N is the size of the first set, and K is the size of the result set.

## Return Value
**Integer reply**: the number of members in the resulting sorted set at destination.

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
redis> ZDIFFSTORE out 2 zset1 zset2
(integer) 1
redis> ZRANGE out 0 -1 WITHSCORES
1) "three"
2) "3"
```
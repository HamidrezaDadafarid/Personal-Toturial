# ZINTER
This command is similar to ZINTERSTORE, but instead of storing the resulting sorted set, it is returned to the client.

## Syntax
```
ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]]
[AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
```

## Time Complexity
O(N*K)+O(M*log(M)) worst case with N being the smallest input sorted set, K being the number of input sorted sets and M being the number of elements in the resulting sorted set.

## Return Value
**Array reply**: the result of the intersection including, optionally, scores when the WITHSCORES option is used.

## Examples
```bash
redis> ZADD zset1 1 "one"
(integer) 1
redis> ZADD zset1 2 "two"
(integer) 1
redis> ZADD zset2 1 "one"
(integer) 1
redis> ZADD zset2 2 "two"
(integer) 1
redis> ZADD zset2 3 "three"
(integer) 1
redis> ZINTER 2 zset1 zset2
1) "one"
2) "two"
redis> ZINTER 2 zset1 zset2 WITHSCORES
1) "one"
2) "2"
3) "two"
4) "4"
```
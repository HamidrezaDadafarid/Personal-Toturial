# ZUNION
This command is similar to ZUNIONSTORE, but instead of storing the resulting sorted set, it is returned to the client.

## Syntax
```
ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]]
[AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
```

## Time Complexity
O(N)+O(M*log(M)) with N being the sum of the sizes of the input sorted sets, and M being the number of elements in the resulting sorted set.

## Return Value
**Array reply**: the result of the union with, optionally, their scores when WITHSCORES is used.

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
redis> ZUNION 2 zset1 zset2
1) "one"
2) "three"
3) "two"
redis> ZUNION 2 zset1 zset2 WITHSCORES
1) "one"
2) "2"
3) "three"
4) "3"
5) "two"
6) "4"
```
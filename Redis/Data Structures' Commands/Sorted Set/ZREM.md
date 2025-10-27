# ZREM
Removes the specified members from the sorted set stored at key. Non existing members are ignored.

An error is returned when key exists and does not hold a sorted set.

## Syntax
```
ZREM key member [member ...]
```

## Time Complexity
O(M*log(N)) with N being the number of elements in the sorted set and M the number of elements to be removed.

## Return Value
**Integer reply**: the number of members removed from the sorted set, not including non-existing members.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZADD myzset 3 "three"
(integer) 1
redis> ZREM myzset "two"
(integer) 1
redis> ZRANGE myzset 0 -1 WITHSCORES
1) "one"
2) "1"
3) "three"
4) "3"
```
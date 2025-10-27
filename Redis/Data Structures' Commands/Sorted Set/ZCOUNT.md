# ZCOUNT
Returns the number of elements in the sorted set at key with a score between min and max.

The min and max arguments have the same semantic as described for ZRANGEBYSCORE.

**Note**: the command has a complexity of just O(log(N)) because it uses elements ranks (see ZRANK) to get an idea of the range. Because of this there is no need to do a work proportional to the size of the range.

## Syntax
```
ZCOUNT key min max
```

## Time Complexity
O(log(N)) with N being the number of elements in the sorted set.

## Return Value
**Integer reply**: the number of members in the specified score range.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZADD myzset 3 "three"
(integer) 1
redis> ZCOUNT myzset -inf +inf
(integer) 3
redis> ZCOUNT myzset (1 3
(integer) 2
redis> ZCOUNT myzset 1 (3
(integer) 2
```
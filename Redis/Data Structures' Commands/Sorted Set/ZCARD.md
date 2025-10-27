# ZCARD
Returns the sorted set cardinality (number of elements) of the sorted set stored at key.

## Syntax
```
ZCARD key
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the cardinality (number of members) of the sorted set, or 0 if the key doesn't exist.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZCARD myzset
(integer) 2
```
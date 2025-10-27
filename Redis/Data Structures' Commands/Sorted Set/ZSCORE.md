# ZSCORE
Returns the score of member in the sorted set at key.

If member does not exist in the sorted set, or key does not exist, nil is returned.

## Syntax
```
ZSCORE key member
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Bulk string reply**: the score of the member (a double-precision floating point number), represented as a string.
2. **Nil reply**: if member does not exist in the sorted set, or the key does not exist.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZSCORE myzset "one"
"1"
```
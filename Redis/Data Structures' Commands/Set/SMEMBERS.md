# SMEMBERS
Returns all the members of the set value stored at key.

This has the same effect as running SINTER with one argument key.

## Syntax
```
SMEMBERS key
```

## Time Complexity
O(N) where N is the set cardinality.

## Return Value
**Array reply**: an array with all the members of the set.

## Examples
```bash
redis> SADD myset "Hello"
(integer) 1
redis> SADD myset "World"
(integer) 1
redis> SMEMBERS myset
1) "Hello"
2) "World"
```
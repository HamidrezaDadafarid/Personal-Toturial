# SISMEMBER
Returns if member is a member of the set stored at key.

## Syntax
```
SISMEMBER key member
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Integer reply**: 0 if the element is not a member of the set, or when the key does not exist.
2. **Integer reply**: 1 if the element is a member of the set.

## Examples
```bash
redis> SADD myset "one"
(integer) 1
redis> SISMEMBER myset "one"
(integer) 1
redis> SISMEMBER myset "two"
(integer) 0
```
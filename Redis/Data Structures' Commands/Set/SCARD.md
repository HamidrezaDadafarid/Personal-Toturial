# SCARD
Returns the set cardinality (number of elements) of the set stored at key.

## Syntax
```
SCARD key
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the cardinality (number of elements) of the set, or 0 if the key does not exist.

## Examples
```bash
redis> SADD myset "Hello"
(integer) 1
redis> SADD myset "World"
(integer) 1
redis> SCARD myset
(integer) 2
```
# HVALS
Returns all values in the hash stored at key.

## Syntax
```
HVALS key
```

## Time Complexity
O(N) where N is the size of the hash.

## Return Value
**Array reply**: a list of values in the hash, or an empty list when the key does not exist

## Examples
```bash
redis> HSET myhash field1 "Hello"
(integer) 1
redis> HSET myhash field2 "World"
(integer) 1
redis> HVALS myhash
1) "Hello"
2) "World"
```
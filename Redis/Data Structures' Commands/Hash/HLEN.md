# HLEN
Returns the number of fields contained in the hash stored at key.

## Syntax
```
HLEN key
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the number of fields in the hash, or 0 when the key does not exist.

## Examples
```bash
redis> HSET myhash field1 "Hello"
(integer) 1
redis> HSET myhash field2 "World"
(integer) 1
redis> HLEN myhash
(integer) 2
```
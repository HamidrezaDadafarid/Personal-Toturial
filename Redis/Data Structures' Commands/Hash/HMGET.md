# HMGET
Returns the values associated with the specified fields in the hash stored at key.

For every field that does not exist in the hash, a nil value is returned. Because non-existing keys are treated as empty hashes, running HMGET against a non-existing key will return a list of nil values.

## Syntax
```
HMGET key field [field ...]
```

## Time Complexity
O(N) where N is the number of fields being requested.

## Return Value
**Array reply**: a list of values associated with the given fields, in the same order as they are requested.

## Examples
```bash
redis> HSET myhash field1 "Hello"
(integer) 1
redis> HSET myhash field2 "World"
(integer) 1
redis> HMGET myhash field1 field2 nofield
1) "Hello"
2) "World"
3) (nil)
```
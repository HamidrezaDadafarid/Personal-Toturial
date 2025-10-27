# HGETALL
Returns all fields and values of the hash stored at key. In the returned value, every field name is followed by its value, so the length of the reply is twice the size of the hash.

## Syntax
```
HGETALL key
```

## Time Complexity
O(N) where N is the size of the hash.

## Return Value
**Array reply**: a list of fields and their values, or an empty list when key does not exist.

## Examples
```bash
redis> HSET myhash field1 "Hello"
(integer) 1
redis> HSET myhash field2 "World"
(integer) 1
redis> HGETALL myhash
1) "field1"
2) "Hello"
3) "field2"
4) "World"
```
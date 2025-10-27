# HSET
Sets the specified fields to their respective values in the hash stored at key.

This command overwrites the values of specified fields that exist in the hash. If key doesn't exist, a new key holding a hash is created.

## Syntax
```
HSET key field value [field value ...]
```

## Time Complexity
O(1) for each field/value pair added, so O(N) to add N field/value pairs when the command is called with multiple field/value pairs.

## Return Value
**Integer reply**: the number of fields that were added.

## Examples
```bash
redis> HSET myhash field1 "Hello"
(integer) 1
redis> HGET myhash field1
"Hello"
redis> HSET myhash field2 "Hi" field3 "World"
(integer) 2
redis> HGET myhash field2
"Hi"
redis> HGET myhash field3
"World"
redis> HGETALL myhash
1) "field1"
2) "Hello"
3) "field2"
4) "Hi"
5) "field3"
6) "World"
```
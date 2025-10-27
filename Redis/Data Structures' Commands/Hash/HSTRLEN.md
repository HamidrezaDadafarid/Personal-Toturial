# HSTRLEN
Returns the string length of the value associated with field in the hash stored at key. If the key or the field do not exist, 0 is returned.

## Syntax
```
HSTRLEN key field
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the string length of the value associated with the field, or zero when the field isn't present in the hash or the key doesn't exist at all.

## Examples
```bash
redis> HSET myhash f1 HelloWorld f2 99 f3 -256
(integer) 3
redis> HSTRLEN myhash f1
(integer) 10
redis> HSTRLEN myhash f2
(integer) 2
redis> HSTRLEN myhash f3
(integer) 4
```
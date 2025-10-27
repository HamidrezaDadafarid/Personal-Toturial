# HINCRBY
Increments the number stored at field in the hash stored at key by increment. If key does not exist, a new key holding a hash is created. If field does not exist the value is set to 0 before the operation is performed.

The range of values supported by HINCRBY is limited to 64 bit signed integers.

## Syntax
```
HINCRBY key field increment
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the value of the field after the increment operation.

## Examples
```bash
redis> HSET myhash field 5
(integer) 1
redis> HINCRBY myhash field 1
(integer) 6
redis> HINCRBY myhash field -1
(integer) 5
redis> HINCRBY myhash field -10
(integer) -5
```
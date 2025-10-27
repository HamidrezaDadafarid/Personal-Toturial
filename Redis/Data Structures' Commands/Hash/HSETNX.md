# HSETNX
Sets field in the hash stored at key to value, only if field does not yet exist. If key does not exist, a new key holding a hash is created. If field already exists, this operation has no effect.

## Syntax
```
HSETNX key field value
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Integer reply**: 0 if the field already exists in the hash and no operation was performed.
2. **Integer reply**: 1 if the field is a new field in the hash and the value was set.

## Examples
```bash
redis> HSETNX myhash field "Hello"
(integer) 1
redis> HSETNX myhash field "World"
(integer) 0
redis> HGET myhash field
"Hello"
```
# HEXISTS
Returns if field is an existing field in the hash stored at key.

## Syntax
```
HEXISTS key field
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Integer reply**: 0 if the hash does not contain the field, or the key does not exist.
2. **Integer reply**: 1 if the hash contains the field.

## Examples
```bash
redis> HSET myhash field1 "foo"
(integer) 1
redis> HEXISTS myhash field1
(integer) 1
redis> HEXISTS myhash field2
(integer) 0
```
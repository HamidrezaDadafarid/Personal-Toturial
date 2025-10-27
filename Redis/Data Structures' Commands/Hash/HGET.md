# HGET
Returns the value associated with field in the hash stored at key.

## Syntax
```
HGET key field
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Bulk string reply**: The value associated with the field.
2. **Nil reply**: If the field is not present in the hash or key does not exist.

## Examples
```bash
redis> HSET myhash field1 "foo"
(integer) 1
redis> HGET myhash field1
"foo"
redis> HGET myhash field2
(nil)
```
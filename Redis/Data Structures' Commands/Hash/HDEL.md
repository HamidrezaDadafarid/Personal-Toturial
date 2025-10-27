# HDEL
Removes the specified fields from the hash stored at key. Specified fields that do not exist within this hash are ignored. Deletes the hash if no fields remain. If key does not exist, it is treated as an empty hash and this command returns 0.

## Syntax
```
HDEL key field [field ...]
```

## Time Complexity
O(N) where N is the number of fields to be removed.

## Return Value
**Integer reply**: the number of fields that were removed from the hash, excluding any specified but non-existing fields.

## Examples
```bash
redis> HSET myhash field1 "foo"
(integer) 1
redis> HDEL myhash field1
(integer) 1
redis> HDEL myhash field2
(integer) 0
```
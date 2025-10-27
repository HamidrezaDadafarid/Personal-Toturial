# SINTERSTORE
This command is equal to SINTER, but instead of returning the resulting set, it is stored in destination.

If destination already exists, it is overwritten.

## Syntax
```
SINTERSTORE destination key [key ...]
```

## Time Complexity
O(N*M) worst case where N is the cardinality of the smallest set and M is the number of sets.

## Return Value
**Integer reply**: the number of elements in the resulting set.

## Examples
```bash
redis> SADD key1 "a"
(integer) 1
redis> SADD key1 "b"
(integer) 1
redis> SADD key1 "c"
(integer) 1
redis> SADD key2 "c"
(integer) 1
redis> SADD key2 "d"
(integer) 1
redis> SADD key2 "e"
(integer) 1
redis> SINTERSTORE key key1 key2
(integer) 1
redis> SMEMBERS key
1) "c"
```
# SUNIONSTORE
This command is equal to SUNION, but instead of returning the resulting set, it is stored in destination.

If destination already exists, it is overwritten.

## Syntax
```
SUNIONSTORE destination key [key ...]
```

## Time Complexity
O(N) where N is the total number of elements in all given sets.

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
redis> SUNIONSTORE key key1 key2
(integer) 5
redis> SMEMBERS key
1) "a"
2) "b"
3) "c"
4) "d"
5) "e"
```
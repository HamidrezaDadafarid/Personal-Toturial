# SDIFF
Returns the members of the set resulting from the difference between the first set and all the successive sets.

## Syntax
```
SDIFF key [key ...]
```

## Time Complexity
O(N) where N is the total number of elements in all given sets.

## Return Value
**Array reply**: a list with members of the resulting set.

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
redis> SDIFF key1 key2
1) "b"
2) "a"
```
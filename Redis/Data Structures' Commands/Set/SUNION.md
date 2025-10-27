# SUNION
Returns the members of the set resulting from the union of all the given sets.

## Syntax
```
SUNION key [key ...]
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
redis> SUNION key1 key2
1) "c"
2) "a"
3) "e"
4) "b"
5) "d"
```
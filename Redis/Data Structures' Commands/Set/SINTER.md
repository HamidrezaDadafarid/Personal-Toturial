# SINTER
Returns the members of the set resulting from the intersection of all the given sets.

## Syntax
```
SINTER key [key ...]
```

## Time Complexity
O(N*M) worst case where N is the cardinality of the smallest set and M is the number of sets.

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
redis> SINTER key1 key2
1) "c"
```
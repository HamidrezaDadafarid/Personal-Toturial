# ZPOPMAX
Removes and returns up to count members with the highest scores in the sorted set stored at key.

When left unspecified, the default value for count is 1. Specifying a count value that is higher than the sorted set's cardinality will not produce an error. When returning multiple elements, the one with the highest score will be the first, followed by the elements with lower scores.

## Syntax
```
ZPOPMAX key [count]
```

## Time Complexity
O(log(N)*M) with N being the number of elements in the sorted set, and M being the number of elements popped.

## Return Value
**Array reply**: a list of popped elements and scores.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZADD myzset 3 "three"
(integer) 1
redis> ZPOPMAX myzset
1) "three"
2) "3"
```
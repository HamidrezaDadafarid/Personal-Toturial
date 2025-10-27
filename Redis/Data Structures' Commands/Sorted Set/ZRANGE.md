# ZRANGE
Returns the specified range of elements in the sorted set stored at `key`.

ZRANGE can perform different types of range queries: by index (rank), by the score, or by lexicographical order.

## Syntax
```
ZRANGE key start stop [BYSCORE | BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
```

## Time Complexity
O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements returned.

## Options
The order of elements is from the lowest to the highest score. Elements with the same score are ordered lexicographically.

The optional REV argument reverses the ordering, so elements are ordered from highest to lowest score, and score ties are resolved by reverse lexicographical ordering.

The optional LIMIT argument can be used to obtain a sub-range from the matching elements (similar to SELECT LIMIT offset, count in SQL). A negative `count` returns all elements from the `offset`. Keep in mind that if `offset` is large, the sorted set needs to be traversed for `offset` elements before getting to the elements to return, which can add up to O(N) time complexity.

The optional WITHSCORES argument supplements the command's reply with the scores of elements returned. The returned list contains value1,score1,...,valueN,scoreN instead of value1,...,valueN. Client libraries are free to return a more appropriate data type (suggestion: an array with (value, score) arrays/tuples).

## Return Value
**Array reply**: a list of members in the specified range with, optionally, their scores when the WITHSCORES option is given.

## Examples
```bash
redis> ZADD myzset 1 "one" 2 "two" 3 "three"
(integer) 3
redis> ZRANGE myzset 0 -1
1) "one"
2) "two"
3) "three"
redis> ZRANGE myzset 2 3
1) "three"
redis> ZRANGE myzset -2 -1
1) "two"
2) "three"
```

The following example using WITHSCORES shows how the command returns always an array, but this time, populated with element_1, score_1, element_2, score_2, ..., element_N, score_N.
```bash
redis> ZADD myzset 1 "one" 2 "two" 3 "three"
(integer) 3
redis> ZRANGE myzset 0 1 WITHSCORES
1) "one"
2) "1"
3) "two"
4) "2"
```

This example shows how to query the sorted set by score, excluding the value 1 and up to infinity, returning only the second element of the result:
```bash
redis> ZADD myzset 1 "one" 2 "two" 3 "three"
(integer) 3
redis> ZRANGE myzset (1 +inf BYSCORE LIMIT 1 1
1) "three"
```
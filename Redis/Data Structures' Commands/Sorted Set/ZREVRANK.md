# ZREVRANK
Returns the rank of member in the sorted set stored at key, with the scores ordered from high to low. The rank (or index) is 0-based, which means that the member with the highest score has rank 0.

The optional WITHSCORE argument supplements the command's reply with the score of the element returned.

Use ZRANK to get the rank of an element with the scores ordered from low to high.

## Syntax
```
ZREVRANK key member [WITHSCORE]
```

## Time Complexity
O(log(N))

## Return Value
One of the following:

1. **Nil reply**: if the key does not exist or the member does not exist in the sorted set.
2. **Integer reply**: The rank of the member when WITHSCORE is not used.
3. **Array reply**: The rank and score of the member when WITHSCORE is used.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZADD myzset 3 "three"
(integer) 1
redis> ZREVRANK myzset "one"
(integer) 2
redis> ZREVRANK myzset "four"
(nil)
redis> ZREVRANK myzset "three" WITHSCORE
1) (integer) 0
2) "3"
redis> ZREVRANK myzset "four" WITHSCORE
(nil)
```
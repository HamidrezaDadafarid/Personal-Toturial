# ZRANDMEMBER
When called with just the key argument, return a random element from the sorted set value stored at key.

If the provided count argument is positive, return an array of distinct elements. The array's length is either count or the sorted set's cardinality (ZCARD), whichever is lower.

If called with a negative count, the behavior changes and the command is allowed to return the same element multiple times. In this case, the number of returned elements is the absolute value of the specified count.

The optional WITHSCORES modifier changes the reply so it includes the respective scores of the randomly selected elements from the sorted set.

## Syntax
```
ZRANDMEMBER key [count [WITHSCORES]]
```

## Time Complexity
O(N) where N is the number of members returned

## Return Value
**Bulk string reply**: without the additional count argument, the command returns a randomly selected member, or Nil reply when key doesn't exist. Array reply: when the additional count argument is passed, the command returns an array of members, or an empty array when key doesn't exist. If the WITHSCORES modifier is used, the reply is a list of members and their scores from the sorted set.

## Examples
```bash
redis> ZADD dadi 1 uno 2 due 3 tre 4 quattro 5 cinque 6 sei
(integer) 6
redis> ZRANDMEMBER dadi
"tre"
redis> ZRANDMEMBER dadi
"sei"
redis> ZRANDMEMBER dadi -5 WITHSCORES
1) "uno"
2) "1"
3) "quattro"
4) "4"
5) "sei"
6) "6"
7) "cinque"
8) "5"
9) "uno"
10) "1"
```
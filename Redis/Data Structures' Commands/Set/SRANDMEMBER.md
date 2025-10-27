# SRANDMEMBER
When called with just the key argument, return a random element from the set value stored at key.

If the provided count argument is positive, return an array of distinct elements. The array's length is either count or the set's cardinality (SCARD), whichever is lower.

If called with a negative count, the behavior changes and the command is allowed to return the same element multiple times. In this case, the number of returned elements is the absolute value of the specified count.

## Syntax
```
SRANDMEMBER key [count]
```

## Time Complexity
Without the count argument O(1), otherwise O(N) where N is the absolute value of the passed count.

## Return Value
One of the following:

1. **Bulk string reply**: without the additional count argument, the command returns a randomly selected member, or a Nil reply when key doesn't exist.
2. **Array reply**: when the optional count argument is passed, the command returns an array of members, or an empty array when key doesn't exist.

## Examples
```bash
redis> SADD myset one two three
(integer) 3
redis> SRANDMEMBER myset
"one"
redis> SRANDMEMBER myset 2
1) "two"
2) "three"
redis> SRANDMEMBER myset -5
1) "one"
2) "three"
3) "three"
4) "one"
5) "three"
```
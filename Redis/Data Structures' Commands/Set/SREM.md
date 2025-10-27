# SREM
Remove the specified members from the set stored at key. Specified members that are not a member of this set are ignored. If key does not exist, it is treated as an empty set and this command returns 0.

An error is returned when the value stored at key is not a set.

## Syntax
```
SREM key member [member ...]
```

## Time Complexity
O(N) where N is the number of members to be removed.

## Return Value
**Integer reply**: the number of members that were removed from the set, not including non existing members.

## Examples
```bash
redis> SADD myset "one"
(integer) 1
redis> SADD myset "two"
(integer) 1
redis> SADD myset "three"
(integer) 1
redis> SREM myset "one"
(integer) 1
redis> SREM myset "four"
(integer) 0
redis> SMEMBERS myset
1) "two"
2) "three"
```
# SPOP
Removes and returns one or more random members from the set value store at key.

This operation is similar to SRANDMEMBER, that returns one or more random elements from a set but does not remove it.

By default, the command pops a single member from the set. When provided with the optional count argument, the reply will consist of up to count members, depending on the set's cardinality.

## Syntax
```
SPOP key [count]
```

## Time Complexity
Without the count argument O(1), otherwise O(N) where N is the value of the passed count.

## Return Value
One of the following:

1. **Nil reply**: if the key does not exist.
2. **Bulk string reply**: when called without the count argument, the removed member.
3. **Array reply**: when called with the count argument, a list of the removed members.

## Examples
```bash
redis> SADD myset "one"
(integer) 1
redis> SADD myset "two"
(integer) 1
redis> SADD myset "three"
(integer) 1
redis> SPOP myset
"two"
redis> SMEMBERS myset
1) "one"
2) "three"
redis> SADD myset "four"
(integer) 1
redis> SADD myset "five"
(integer) 1
redis> SPOP myset 3
1) "one"
2) "three"
3) "four"
redis> SMEMBERS myset
1) "five"
```
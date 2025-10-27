# SADD
Add the specified members to the set stored at key. Specified members that are already a member of this set are ignored. If key does not exist, a new set is created before adding the specified members.

An error is returned when the value stored at key is not a set.

## Syntax
```
SADD key member [member ...]
```

## Time Complexity
O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.

## Return Value
**Integer reply**: the number of elements that were added to the set, not including all the elements already present in the set.

## Examples
```bash
redis> SADD myset "Hello" "World"
(integer) 2
redis> SADD myset "World"
(integer) 0
redis> SMEMBERS myset
1) "Hello"
2) "World"
```
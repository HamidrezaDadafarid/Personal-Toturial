# LREM
Removes the first count occurrences of elements equal to element from the list stored at key. The count argument influences the operation in the following ways:

- count > 0: Remove elements equal to element moving from head to tail.
- count < 0: Remove elements equal to element moving from tail to head.
- count = 0: Remove all elements equal to element.
For example, LREM list -2 "hello" will remove the last two occurrences of "hello" in the list stored at list.

Note that non-existing keys are treated like empty lists, so when key does not exist, the command will always return 0.

## Syntax
```
LREM key count element
```

## Time Complexity
O(N+M) where N is the length of the list and M is the number of elements removed.

## Return Value
**Integer reply**: the number of removed elements.

## Examples
```bash
redis> RPUSH mylist "hello"
(integer) 1
redis> RPUSH mylist "hello"
(integer) 2
redis> RPUSH mylist "foo"
(integer) 3
redis> RPUSH mylist "hello"
(integer) 4
redis> LREM mylist -2 "hello"
(integer) 2
redis> LRANGE mylist 0 -1
1) "hello"
2) "foo"
```
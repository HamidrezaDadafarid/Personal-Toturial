# BLPOP
BLPOP is a blocking list pop primitive. It is the blocking version of LPOP because it blocks the connection when there are no elements to pop from any of the given lists. An element is popped from the head of the first list that is non-empty, with the given keys being checked in the order that they are given.

## Syntax
```
BLPOP key [key ...] timeout
```

## Time Complexity
O(N) where N is the number of provided keys.

## Return Value
One of the following:

1. **Nil reply**: no element could be popped and the timeout expired
2. **Array reply**: the key from which the element was popped and the value of the popped element.

## Examples
```bash
redis> DEL list1 list2
(integer) 0
redis> RPUSH list1 a b c
(integer) 3
redis> BLPOP list1 list2 0
1) "list1"
2) "a"
```
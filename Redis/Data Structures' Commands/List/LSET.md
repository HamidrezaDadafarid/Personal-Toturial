# LSET
Sets the list element at index to element. For more information on the index argument, see LINDEX.

An error is returned for out of range indexes.

## Syntax
```
LSET key index element
```

## Time Complexity
O(N) where N is the length of the list. Setting either the first or the last element of the list is O(1).

## Return Value
**Simple string reply**: OK.

## Examples
```bash
redis> RPUSH mylist "one"
(integer) 1
redis> RPUSH mylist "two"
(integer) 2
redis> RPUSH mylist "three"
(integer) 3
redis> LSET mylist 0 "four"
"OK"
redis> LSET mylist -2 "five"
"OK"
redis> LRANGE mylist 0 -1
1) "four"
2) "five"
3) "three"
```
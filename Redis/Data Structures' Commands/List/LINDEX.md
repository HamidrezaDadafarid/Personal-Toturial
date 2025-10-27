# LINDEX
Returns the element at index index in the list stored at key. The index is zero-based, so 0 means the first element, 1 the second element and so on. Negative indices can be used to designate elements starting at the tail of the list. Here, -1 means the last element, -2 means the penultimate and so forth.

When the value at key is not a list, an error is returned.

## Syntax
```
LINDEX key index
```

## Time Complexity
O(N) where N is the number of elements to traverse to get to the element at index. This makes asking for the first or the last element of the list O(1).

## Return Value
One of the following:

1. **Nil reply**: when index is out of range.
2. **Bulk string reply**: the requested element.

## Examples
```bash
redis> LPUSH mylist "World"
(integer) 1
redis> LPUSH mylist "Hello"
(integer) 2
redis> LINDEX mylist 0
"Hello"
redis> LINDEX mylist -1
"World"
redis> LINDEX mylist 3
(nil)
```
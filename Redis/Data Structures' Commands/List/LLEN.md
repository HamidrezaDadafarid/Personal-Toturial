# LLEN
Returns the length of the list stored at key. If key does not exist, it is interpreted as an empty list and 0 is returned. An error is returned when the value stored at key is not a list.

## Syntax
```
LLEN key
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the length of the list.

## Examples
```bash
redis> LPUSH mylist "World"
(integer) 1
redis> LPUSH mylist "Hello"
(integer) 2
redis> LLEN mylist
(integer) 2
```
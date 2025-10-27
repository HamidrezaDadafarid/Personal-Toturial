# LPUSHX
Inserts specified values at the head of the list stored at key, only if key already exists and holds a list. In contrary to LPUSH, no operation will be performed when key does not yet exist.

## Syntax
```
LPUSHX key element [element ...]
```

## Time Complexity
O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.

## Return Value
**Integer reply**: the length of the list after the push operation.

## Examples
```bash
redis> LPUSH mylist "World"
(integer) 1
redis> LPUSHX mylist "Hello"
(integer) 2
redis> LPUSHX myotherlist "Hello"
(integer) 0
redis> LRANGE mylist 0 -1
1) "Hello"
2) "World"
redis> LRANGE myotherlist 0 -1
(empty array)
```
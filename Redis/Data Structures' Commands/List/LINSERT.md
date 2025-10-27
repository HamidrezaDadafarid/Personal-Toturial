# LINSERT
Inserts element in the list stored at key either before or after the reference value pivot.

When key does not exist, it is considered an empty list and no operation is performed.

An error is returned when key exists but does not hold a list value.

## Syntax
```
LINSERT key <BEFORE | AFTER> pivot element
```

## Time Complexity
O(N) where N is the number of elements to traverse before seeing the value pivot. This means that inserting somewhere on the left end on the list (head) can be considered O(1) and inserting somewhere on the right end (tail) is O(N).

## Return Value
One of the following:

1. **Integer reply**: the list length after a successful 2. insert operation.
2. **Integer reply**: 0 when the key doesn't exist.
3. **Integer reply**: -1 when the pivot wasn't found.

## Examples
```bash
redis> RPUSH mylist "Hello"
(integer) 1
redis> RPUSH mylist "World"
(integer) 2
redis> LINSERT mylist BEFORE "World" "There"
(integer) 3
redis> LRANGE mylist 0 -1
1) "Hello"
2) "There"
3) "World"
```
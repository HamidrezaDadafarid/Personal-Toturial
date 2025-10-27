# RPOP
Removes and returns the last elements of the list stored at key.

By default, the command pops a single element from the end of the list. When provided with the optional `count` argument, the reply will consist of up to count elements, depending on the list's length.

## Syntax
```
RPOP key [count]
```

## Time Complexity
O(N) where N is the number of elements returned

## Return Value
One of the following:

1. **Nil reply**: if the key does not exist.
2. **Bulk string reply**: when called without the count argument, the value of the first element.
3. **Array reply**: when called with the count argument, a list of popped elements.

## Examples
```bash
redis> RPUSH mylist "one" "two" "three" "four" "five"
(integer) 5
redis> RPOP mylist
"five"
redis> RPOP mylist 2
1) "four"
2) "three"
redis> LRANGE mylist 0 -1
1) "one"
2) "two"
```
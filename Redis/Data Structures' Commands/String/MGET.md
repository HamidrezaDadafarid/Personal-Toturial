# MGET
Returns the values of all specified keys. For every key that does not hold a string value or does not exist, the special value nil is returned. Because of this, the operation never fails.

## Syntax
```
MGET key [key ...]
```

## Time Complexity
O(N) where N is the number of keys to retrieve.

## Return Value
**Array reply**: a list of values at the specified keys.

## Examples
```bash
redis> SET key1 "Hello"
"OK"
redis> SET key2 "World"
"OK"
redis> MGET key1 key2 nonexisting
1) "Hello"
2) "World"
3) (nil)
```
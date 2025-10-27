# PERSIST
Remove the existing timeout on key, turning the key from volatile (a key with an expire set) to persistent (a key that will never expire as no timeout is associated).

## Syntax
```
PERSIST key
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Integer reply**: 0 if key does not exist or does not have an associated timeout.
2. **Integer reply**: 1 if the timeout has been removed.

## Examples
```bash
redis> SET mykey "Hello"
"OK"
redis> EXPIRE mykey 10
(integer) 1
redis> TTL mykey
(integer) 10
redis> PERSIST mykey
(integer) 1
redis> TTL mykey
(integer) -1
```
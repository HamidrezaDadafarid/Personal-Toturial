# GETSET
Atomically sets key to value and returns the old value stored at key. Returns an error when key exists but does not hold a string value. Any previous time to live associated with the key is discarded on successful SET operation.

## Syntax
```
GETSET key value
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Bulk string reply**: the old value stored at the key.
2. **Nil reply**: if the key does not exist.

## Examples
```bash
redis> SET mykey "Hello"
"OK"
redis> GETSET mykey "World"
"Hello"
redis> GET mykey
"World"
```
# PTTL
Like TTL this command returns the remaining time to live of a key that has an expire set, with the sole difference that TTL returns the amount of remaining time in seconds while PTTL returns it in milliseconds.

## Syntax
```
PTTL key
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Integer reply**: TTL in milliseconds.
2. **Integer reply**: -1 if the key exists but has no associated expiration.
3. **Integer reply**: -2 if the key does not exist.

## Examples
```bash
redis> SET mykey "Hello"
"OK"
redis> EXPIRE mykey 1
(integer) 1
redis> PTTL mykey
(integer) 999
```
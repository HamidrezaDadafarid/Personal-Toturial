# TTL
Returns the remaining time to live of a key that has a timeout. This introspection capability allows a Redis client to check how many seconds a given key will continue to be part of the dataset.

## Syntax
```
TTL key
```

## Time Complexity
O(1)

## Return Value
One of the following:

- Integer reply: TTL in seconds.
- Integer reply: -1 if the key exists but has no associated expiration.
- Integer reply: -2 if the key does not exist.

## Examples
```bash
> SET mykey "Hello"
"OK"
> EXPIRE mykey 10
(integer) 1
> TTL mykey
(integer) 10
```
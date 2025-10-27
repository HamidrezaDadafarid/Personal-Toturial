# EXISTS
Returns if key exists.

The user should be aware that if the same existing key is mentioned in the arguments multiple times, it will be counted multiple times. So if somekey exists, EXISTS somekey somekey will return 2.

## Syntax
```
EXISTS key [key ...]
```

## Time Complexity
O(N) where N is the number of keys to check.

## Return Value
Integer reply: the number of keys that exist from those specified as arguments.

## Examples
```bash
redis> SET key1 "Hello"
"OK"
redis> EXISTS key1
(integer) 1
redis> EXISTS nosuchkey
(integer) 0
redis> SET key2 "World"
"OK"
redis> EXISTS key1 key2 nosuchkey
(integer) 2
```
# SET
Set key to hold the **string** value. If key already holds a value, it is overwritten, regardless of its type. Any previous time to live associated with the key is discarded on successful SET operation. If the string we set is more than 1 word, we have to put it in double quotes.

## Syntax
```
SET key value [NX | XX] [GET] [EX seconds | PX milliseconds |
EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]
```

## Time Complexity
O(1)

## Options
- EX seconds -- Set the specified expire time, in seconds (a positive integer).

## Return Value
If GET was not specified, any of the following:
1. **Nil reply**: Operation was aborted (conflict with one of the XX/NX options). The key was not set.
2. **Simple string reply**: OK: The key was set.

If GET was specified, any of the following:
1. **Nil reply**: The key didn't exist before the SET. If XX was specified, the key was not set. Otherwise, the key was set.
2. **Bulk string reply**: The previous value of the key. If NX was specified, the key was not set. Otherwise, the key was set.

## Examples
```bash
redis> SET name hamid reza
(error) ERR syntax error
redis> SET mykey Hello
"OK"
redis> GET mykey
"Hello"
redis> SET anotherkey "will expire in a minute" EX 60
"OK"
redis>
```
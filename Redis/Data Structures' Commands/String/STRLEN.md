# STRLEN
Returns the length of the string value stored at key. An error is returned when key holds a non-string value.

## Syntax
```
STRLEN key
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the length of the string stored at key, or 0 when the key does not exist.

## Examples
```bash
redis> SET mykey "Hello world"
"OK"
redis> STRLEN mykey
(integer) 11
redis> STRLEN nonexisting
(integer) 0
```
# INCRBY
Increments the number stored at key by increment. If the key does not exist, it is set to 0 before performing the operation. An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer. This operation is limited to 64 bit signed integers.

## Syntax
```
INCRBY key increment
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the value of the key after the increment.

## Examples
```bash
redis> SET mykey "10"
"OK"
redis> INCRBY mykey 5
(integer) 15
```
# DECR
Decrements the number stored at key by one. If the key does not exist, it is set to 0 before performing the operation. An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer. This operation is limited to 64 bit signed integers.

## Syntax
```
DECR key
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the value of the key after decrementing it.

## Examples
```bash
redis> SET mykey "10"
"OK"
redis> DECR mykey
(integer) 9
redis> SET mykey "234293482390480948029348230948"
"OK"
redis> DECR mykey
(error) value is not an integer or out of range
```
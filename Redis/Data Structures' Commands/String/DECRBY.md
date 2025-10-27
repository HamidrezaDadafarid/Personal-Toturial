# DECRBY
The DECRBY command reduces the value stored at the specified key by the specified decrement. If the key does not exist, it is initialized with a value of 0 before performing the operation. If the key's value is not of the correct type or cannot be represented as an integer, an error is returned. This operation is limited to 64-bit signed integers.

## Syntax
```
DECRBY key decrement
```

## Time Complexity
O(1)

## Return Value
**Integer reply**: the value of the key after decrementing it.

## Examples
```bash
redis> SET mykey "10"
"OK"
redis> DECRBY mykey 3
(integer) 7
```
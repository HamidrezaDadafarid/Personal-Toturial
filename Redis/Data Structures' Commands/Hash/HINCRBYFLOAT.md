# HINCRBYFLOAT
Increment the specified field of a hash stored at key, and representing a floating point number, by the specified increment. If the increment value is negative, the result is to have the hash field value decremented instead of incremented. If the field does not exist, it is set to 0 before performing the operation. An error is returned if one of the following conditions occur:

The key contains a value of the wrong type (not a hash).
The current field content or the specified increment are not parsable as a double precision floating point number.

## Syntax
```
HINCRBYFLOAT key field increment
```

## Time Complexity
O(1)

## Return Value
**Bulk string reply**: the value of the field after the increment operation.

## Examples
```bash
redis> HSET mykey field 10.50
(integer) 1
redis> HINCRBYFLOAT mykey field 0.1
"10.6"
redis> HINCRBYFLOAT mykey field -5
"5.6"
redis> HSET mykey field 5.0e3
(integer) 0
redis> HINCRBYFLOAT mykey field 2.0e2
"5200"
```
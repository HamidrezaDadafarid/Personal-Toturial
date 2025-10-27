# EXPIREAT
EXPIREAT has the same effect and semantic as EXPIRE, but instead of specifying the number of seconds representing the TTL (time to live), it takes an absolute Unix timestamp (seconds since January 1, 1970). A timestamp in the past will delete the key immediately.

## Syntax
```
EXPIREAT key unix-time-seconds [NX | XX | GT | LT]
```

## Time Complexity
O(1)

## Options
The EXPIREAT command supports a set of options:

- NX -- Set expiry only when the key has no expiry
- XX -- Set expiry only when the key has an existing expiry
- GT -- Set expiry only when the new expiry is greater than current one
- LT -- Set expiry only when the new expiry is less than current one

A non-volatile key is treated as an infinite TTL for the purpose of GT and LT. The GT, LT and NX options are mutually exclusive.

## Return Value
One of the following:

1. **Integer reply**: 0 if the timeout was not set; for example, the key doesn't exist, or the operation was skipped because of the provided arguments.
2. **Integer reply**: 1 if the timeout was set.

## Examples
```bash
redis> SET mykey "Hello"
"OK"
redis> EXISTS mykey
(integer) 1
redis> EXPIREAT mykey 1293840000
(integer) 1
redis> EXISTS mykey
(integer) 0
```
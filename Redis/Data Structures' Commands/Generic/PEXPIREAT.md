# PEXPIREAT
PEXPIREAT has the same effect and semantic as EXPIREAT, but the Unix time at which the key will expire is specified in milliseconds instead of seconds.

## Syntax
```
PEXPIREAT key unix-time-milliseconds [NX | XX | GT | LT]
```

## Time Complexity
O(1)

## Options
The PEXPIREAT command supports a set of options since Redis 7.0:

- NX -- Set expiry only when the key has no expiry
- XX -- Set expiry only when the key has an existing expiry
- GT -- Set expiry only when the new expiry is greater than current one
- LT -- Set expiry only when the new expiry is less than current one

A non-volatile key is treated as an infinite TTL for the purpose of GT and LT. The GT, LT and NX options are mutually exclusive.

## Return Value
One of the following:

1. **Integer reply**: 0 if the timeout was not set. For example, if the key doesn't exist, or the operation was skipped due to the provided arguments.
2. **Integer reply**: 1 if the timeout was set.

## Examples
```bash
redis> SET mykey "Hello"
"OK"
redis> PEXPIREAT mykey 1555555555005
(integer) 1
redis> TTL mykey
(integer) -2
redis> PTTL mykey
(integer) -2
```
# PEXPIRE
This command works exactly like EXPIRE but the time to live of the key is specified in milliseconds instead of seconds.

## Syntax
```
PEXPIRE key milliseconds [NX | XX | GT | LT]
```

## Time Complexity
O(1)

## Options
The PEXPIRE command supports a set of options since Redis 7.0:

- NX -- Set expiry only when the key has no expiry
- XX -- Set expiry only when the key has an existing expiry
- GT -- Set expiry only when the new expiry is greater than current one
- LT -- Set expiry only when the new expiry is less than current one

A non-volatile key is treated as an infinite TTL for the purpose of GT and LT. The GT, LT and NX options are mutually exclusive.

## Return Value
One of the following:

1. **Integer reply**: 0 if the timeout was not set. For example, if the key doesn't exist, or the operation skipped because of the provided arguments.
2. **Integer reply**: 1 if the timeout was set.

## Examples
```bash
redis> SET mykey "Hello"
"OK"
redis> PEXPIRE mykey 1500
(integer) 1
redis> TTL mykey
(integer) 1
redis> PTTL mykey
(integer) 1498
redis> PEXPIRE mykey 1000 XX
(integer) 1
redis> TTL mykey
(integer) 1
redis> PEXPIRE mykey 1000 NX
(integer) 0
redis> TTL mykey
(integer) 1
```
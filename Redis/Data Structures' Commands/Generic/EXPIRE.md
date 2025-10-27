# EXPIRE
Set a timeout on key. After the timeout has expired, the key will automatically be deleted. A key with an associated timeout is often said to be volatile in Redis terminology.

## Syntax
```
EXPIRE key seconds [NX | XX | GT | LT]
```

## Time Complexity
O(1)

## Options
The EXPIRE command supports a set of options:

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
> SET mykey "Hello"
"OK"
> EXPIRE mykey 10
(integer) 1
> TTL mykey
(integer) 10
> SET mykey "Hello World"
"OK"
> TTL mykey
(integer) -1
> EXPIRE mykey 10 XX
(integer) 0
> TTL mykey
(integer) -1
> EXPIRE mykey 10 NX
(integer) 1
> TTL mykey
(integer) 10
```
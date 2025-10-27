# GET
Get the value of key. If the key does not exist the special value nil is returned. An error is returned if the value stored at key is not a string, because GET only handles string values.

## Syntax
```
GET key
```

## Time Complexity
O(1)

## Return Value
One of the following:

- Bulk string reply: the value of the key.
- Nil reply: if the key does not exist.

## Examples
```bash
redis> GET nonexisting
(nil)
redis> SET mykey "Hello"
"OK"
redis> GET mykey
"Hello"
redis> 
```
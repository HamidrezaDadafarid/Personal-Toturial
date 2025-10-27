# MSET
Sets the given keys to their respective values. MSET replaces existing values with new values, just as regular SET. See MSETNX if you don't want to overwrite existing values.

MSET is atomic, so all given keys are set at once. It is not possible for clients to see that some of the keys were updated while others are unchanged.

## Syntax
```
MSET key value [key value ...]
```

## Time Complexity
O(N) where N is the number of keys to set.

## Return Value
**Simple string reply**: always OK because MSET can't fail.

## Examples
```bash
redis> MSET key1 "Hello" key2 "World"
"OK"
redis> GET key1
"Hello"
redis> GET key2
"World"
```
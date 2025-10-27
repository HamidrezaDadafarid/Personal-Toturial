# KEYS
Returns all keys matching pattern.

While the time complexity for this operation is O(N), the constant times are fairly low. For example, Redis running on an entry level laptop can scan a 1 million key database in 40 milliseconds.

## Syntax
```
KEYS pattern
```

## Time Complexity
O(N) with N being the number of keys in the database, under the assumption that the key names in the database and the given pattern have limited length.

## Return Value
Array reply: a list of keys matching pattern.

## Examples
```bash
redis> MSET firstname Jack lastname Stuntman age 35
"OK"
redis> KEYS *name*
1) "firstname"
2) "lastname"
redis> KEYS a??
1) "age"
redis> KEYS *
1) "firstname"
2) "age"
3) "lastname"
```
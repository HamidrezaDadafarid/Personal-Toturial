# RANDOMKEY
Return a random key from the currently selected database.

## Syntax
```
RANDOMKEY
```

## Time Complexity
O(1)

## Return Value
One of the following:

1. **Nil reply**: when the database is empty.
2. **Bulk string reply**: a random key in database.
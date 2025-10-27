# PFADD
Adds all the element arguments to the HyperLogLog data structure stored at the variable name specified as first argument.

As a side effect of this command the HyperLogLog internals may be updated to reflect a different estimation of the number of unique items added so far (the cardinality of the set).

If the approximated cardinality estimated by the HyperLogLog changed after executing the command, PFADD returns 1, otherwise 0 is returned. The command automatically creates an empty HyperLogLog structure (that is, a Redis String of a specified length and with a given encoding) if the specified key does not exist.

To call the command without elements but just the variable name is valid, this will result into no operation performed if the variable already exists, or just the creation of the data structure if the key does not exist (in the latter case 1 is returned).

## Syntax
```
PFADD key [element [element ...]]
```

## Time Complexity
O(1) to add every element.

## Return Value
One of the following:

1. **Integer reply**: 1 if at least one HyperLogLog internal register was altered.
2. **Integer reply**: 0 if no HyperLogLog internal registers were altered.

## Examples
```bash
redis> PFADD hll a b c d e f g
(integer) 1
redis> PFCOUNT hll
(integer) 7
```
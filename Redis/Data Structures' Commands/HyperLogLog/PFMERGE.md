# PMERGE
Merge multiple HyperLogLog values into a unique value that will approximate the cardinality of the union of the observed Sets of the source HyperLogLog structures.

The computed merged HyperLogLog is set to the destination variable, which is created if does not exist (defaulting to an empty HyperLogLog).

If the destination variable exists, it is treated as one of the source sets and its cardinality will be included in the cardinality of the computed HyperLogLog.

## Syntax
```
PFMERGE destkey [sourcekey [sourcekey ...]]
```

## Time Complexity
O(N) to merge N HyperLogLogs, but with high constant times.

## Return Value
**Simple string reply**: OK.

## Examples
```bash
redis> PFADD hll1 foo bar zap a
(integer) 1
redis> PFADD hll2 a b c foo
(integer) 1
redis> PFMERGE hll3 hll1 hll2
"OK"
redis> PFCOUNT hll3
(integer) 6
```
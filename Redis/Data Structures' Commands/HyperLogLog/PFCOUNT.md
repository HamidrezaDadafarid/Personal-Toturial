# PCOUNT
When called with a single key, returns the approximated cardinality computed by the HyperLogLog data structure stored at the specified variable, which is 0 if the variable does not exist.

When called with multiple keys, returns the approximated cardinality of the union of the HyperLogLogs passed, by internally merging the HyperLogLogs stored at the provided keys into a temporary HyperLogLog.

The HyperLogLog data structure can be used in order to count unique elements in a set using just a small constant amount of memory, specifically 12k bytes for every HyperLogLog (plus a few bytes for the key itself).

The returned cardinality of the observed set is not exact, but approximated with a standard error of 0.81%.

For example in order to take the count of all the unique search queries performed in a day, a program needs to call PFADD every time a query is processed. The estimated number of unique queries can be retrieved with PFCOUNT at any time.

**Note**: as a side effect of calling this function, it is possible that the HyperLogLog is modified, since the last 8 bytes encode the latest computed cardinality for caching purposes. So PFCOUNT is technically a write command.

## Syntax
```
PFCOUNT key [key ...]
```

## Time Complexity
O(1) with a very small average constant time when called with a single key. O(N) with N being the number of keys, and much bigger constant times, when called with multiple keys.

## Return Value
**Integer reply**: the approximated number of unique elements observed via PFADD.

## Examples
```bash
redis> PFADD hll foo bar zap
(integer) 1
redis> PFADD hll zap zap zap
(integer) 0
redis> PFADD hll foo bar
(integer) 0
redis> PFCOUNT hll
(integer) 3
redis> PFADD some-other-hll 1 2 3
(integer) 1
redis> PFCOUNT hll some-other-hll
(integer) 6
```
# ZADD
Adds all the specified members with the specified scores to the sorted set stored at key. It is possible to specify multiple score / member pairs. If a specified member is already a member of the sorted set, the score is updated and the element reinserted at the right position to ensure the correct ordering.

If key does not exist, a new sorted set with the specified members as sole members is created, like if the sorted set was empty. If the key exists but does not hold a sorted set, an error is returned.

The score values should be the string representation of a double precision floating point number. +inf and -inf values are valid values as well.

## Syntax
```
ZADD key [NX | XX] [GT | LT] [CH] [INCR] score member [score member...]
```

## Time Complexity
O(log(N)) for each item added, where N is the number of elements in the sorted set.

## Options
ZADD supports a list of options, specified after the name of the key and before the first score argument. Options are:

- **XX**: Only update elements that already exist. Don't add new elements.
- **NX**: Only add new elements. Don't update already existing elements.
- **LT**: Only update existing elements if the new score is `less than` the current score. This flag doesn't prevent adding new elements.
- **GT**: Only update existing elements if the new score is `greater than` the current score. This flag doesn't prevent adding new elements.
- **CH**: Modify the return value from the number of new elements added, to the total number of elements changed (CH is an abbreviation of changed). Changed elements are `new elements added` and elements already existing for which `the score was updated`. So elements specified in the command line having the same score as they had in the past are not counted. Note: normally the return value of ZADD only counts the number of new elements added.
- **INCR**: When this option is specified ZADD acts like ZINCRBY. Only one score-element pair can be specified in this mode.

**Note**: The **GT**, **LT** and **NX** options are mutually exclusive.

## Return Value
Any of the following:

1. **Nil reply**: if the operation was aborted because of a conflict with one of the XX/NX/LT/GT options.
2. **Integer reply**: the number of new members when the CH option is not used.
3. **Integer reply**: the number of new or updated members when the CH option is used.
4. **Bulk string reply**: the updated score of the member when the INCR option is used.

## Examples
```bash
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 1 "uno"
(integer) 1
redis> ZADD myzset 2 "two" 3 "three"
(integer) 2
redis> ZRANGE myzset 0 -1 WITHSCORES
1) "one"
2) "1"
3) "uno"
4) "1"
5) "two"
6) "2"
7) "three"
8) "3"
```
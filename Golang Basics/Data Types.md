# Data Types in Go

## Primitive types

### numbers
- int(s): int (word size integer), int8, int16, int32, int64 (ranges from $-2^{n-1}$ to $2^{n-1} -1$)
- uint(s): uint (word size unsigned integer), uint8, uint16, uint32, uint64 (ranges from $0$ to $2^{n - 1}$)
- float(s): float32, float64 (`IEEE 754`)
- byte (alias for uint8)
- rune (alias for int32): usually used to store utf-8 characters
- Complex(s): complex64, complex128

### strings
### boolean
### Pointers
- unitptr (word size pointer): uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
- *uint
- *int32
- *string
- ... 

**NOTE: raw pointers have their own static type!**

## Concrete (Composite) Types

### Array: fixed size lists
### Slice: Not fixed size lists
### Map: key-value structure
### Struct
### Interface
### Channels
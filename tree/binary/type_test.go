package binary

// numericType is a generic type constrains consisting of Go numerics.
// NOTE: This is only for demonstration purpose only. There are a number
// of equivalent constraints in standard packages. One example is the
// `constraints` package. You can also use an alias to numeric constraints
// call `comparable`.
// NOTE: int has been deliberate left of the list
type numericType interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

type Node[N numericType] struct {
	value N
	left  *Node[N]
	right *Node[N]
}

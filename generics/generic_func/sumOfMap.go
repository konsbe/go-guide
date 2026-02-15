package generic_func

import "go-guide/generics/types"

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
// K is a type parameter that must be comparable (it can be used as a map key).
// comparable is a constraint that allows any type that supports the == and != operators. (strings, ints, etc. - required for map keys)
// V is a type parameter that can be either int64 or float64. (for map values)
// Constraint - V must be either int64 OR float64 (union type)
// Function parameter - a map with keys of type K and values of type V
func SumIntsOrFloats[K comparable, V types.Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
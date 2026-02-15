package main

import (
	"fmt"
	"go-guide/generics/non_generic_func"
	"go-guide/generics/generic_func"
)

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

	// Initialize a map of float64 values and a map of int64 values, each with two entries.
	// Call the two non-generic functions you declared earlier to find the sum of each map’s values.
    fmt.Printf("Non-Generic Sums: %v and %v\n",
        non_generic_func.SumInts(ints),
        non_generic_func.SumFloats(floats),
	)

	// Call the generic function we just declared, passing each of the maps entries.
	// Specify type arguments – the type names in square brackets 
	// – to be clear about the types that should replace type parameters in the function you’re calling.
	fmt.Printf("Generic Sums: %v and %v\n",
		generic_func.SumIntsOrFloats[string, int64](ints),
		generic_func.SumIntsOrFloats[string, float64](floats),
	)
}

## [Getting started with generics](https://go.dev/doc/tutorial/generics)

The basics of generics in Go. With generics, we can declare and use functions or types that are written to work with any of a set of types provided by calling code.

We will declare two simple non-generic functions, then capture the same logic in a single generic function.

The progress through the following sections:

- Create a folder for our code.
- Add non-generic functions.
- Add a generic function to handle multiple types.
- Remove type arguments when calling the generic function.
- Declare a type constraint.

With generics, we can write one function instead of two. Next, we will add a single generic function for maps containing either integer or float values.

a single generic function that can receive a map containing either integer or float values, effectively replacing the two functions we just wrote with a single function.

To support values of either type, that single function will need a way to declare what types it supports. Calling code, on the other hand, will need a way to specify whether it is calling with an integer or float map.

Each type parameter has a type constraint that acts as a kind of meta-type for the type parameter. Each type constraint specifies the permissible type arguments that calling code can use for the respective type parameter.

While a type parameter’s constraint typically represents a set of types, at compile time the type parameter stands for a single type – the type provided as a type argument by the calling code. If the type argument’s type isn’t allowed by the type parameter’s constraint, the code won’t compile.

Keep in mind that a type parameter must support all the operations the generic code is performing on it. For example, if our function’s code were to try to perform string operations (such as indexing) on a type parameter whose constraint included numeric types, the code wouldn’t compile.


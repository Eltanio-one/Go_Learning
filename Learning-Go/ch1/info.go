// package main

// // var Versus :=
// // most verbose way to state this
// var x int = 10

// // as above, state var, the name, the type and then assign a value

// // but if the value on the right is the expected type, then we can leave off
// // the explicit type
// var x = 10

// // if we want to just declare a variable and assign it zero.
// var x int

// // and declaring multiple variables at once is similar to python
// var x, y int = 10, 20

// // multiple variables of different types
// var x, y = 10, "hello"

// // we can also wrap multiple variable declarations in a declaration list
// var (
// 	a    int
// 	b        = 20
// 	z    int = 30
// 	d, e     = 40, "hello"
// 	f, g string
// )

// // go also supports short declaration formatting
// // using the walrus operator
// var x = 10
// x := 10

// // and using walrus for multiple declarations
// x, y := 10, "hello"

// // and the walrus operator also allows us to overwrite variables
// x := 10
// x, y := 20, "hello"

// // N.B> IF CREATING A PACKAGE, WE MUST USE VAR AS IT ISNT LEGAL TO
// // USE WALRUS OPERATOR OUTSIDE OF FUNCTIONS
// // THE TIMES WHEN TO AVOID :=
// // - WHEN ASSIGNING TO ZERO
// // - WHEN ASSIGNING AN UNTYPED CONSTANT OR LITERAL TO A VARIABLE AND THE DEFAULT TYPE FOR THE
// // CONSTANT OR LITERAL ISNT THE TYPE YOU WANT FOR THE VARIABLE.
// // E.G. NOT x := byte(20), USE var x byte = 20

// // Typed and Untyped Constants
// const x = 10

// // this is untyped, and allows the constant to be assigned to different
// // types

// var y int = x
// var z float64 = x
// var d byte = x

// const typedX int = 10
// // and this can only be assigned to an int variable;
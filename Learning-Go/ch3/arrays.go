// all elements within an array must be of the type declared
var x [3]int

// the above creates an array of 3 ints, but as no values were specified
// all ints are initialised to zero (this is the int zero value).
// if we have initial values, we state these using an array literal
var x = [3]int{10, 20, 30}

// if we have a sparse array (where most elements are set to their zero value)
// we can specify only the indices with values in the array literal
var x = [12]int{1, 5: 4, 6, 10: 100, 15}
// this creates the below array...
[1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15].

// when using an array literal to initialise an array,
// we can leave out the number and replace it with ...
var x = [...]int{10, 20, 30}

// we can use == and != to compare arrays
var x = [...]int{1, 2, 3}
var y = [3]int{1, 2, 3}
fmt.Println(x == y) 
// this would print true

// multidimensional arrays follow the same syntax as python 
var x [2][3]int

// what makes Go arrays lesser used is that if arrays are declared at different
// sizes, they are considered different types and there is no way to 
// type convert arrays of different sizes to the same size and type.
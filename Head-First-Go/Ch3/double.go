package main

import (
	"fmt"
)

// func main() {
// 	// declare an amount to double
// 	amount := 6
// 	// pass a pointer to the declared amount
// 	// this is to ensure that the "pass-by-value" Go functionality doesn't
// 	// prevent the original value not being doubled!
// 	double(&amount)
// 	//
// 	fmt.Println(amount)
// }

func double(number *int) {
	// expect a pointer instead of a copy of a value
	// and double the value at that pointer.
	*number *= 2
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

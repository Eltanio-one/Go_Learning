package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name,
// checking for a valid return value
func TestHelloName(t *testing.T) {
	// set name var equal to Dan
	name := "Dan"
	// assign the want variable to be name, but passed through the MustCompile regexp function
	// to ensure this variable can be compiled
	want := regexp.MustCompile(`\b` + name + `\b`)
	// run the Hello function using name as the argument
	// pass the return values to the variables msg and err
	msg, err := Hello("Dan")
	// if want != msg or the error isnt empty,
	if !want.MatchString(msg) || err != nil {
		// return an error message using the t parameters FatalF method
		t.Fatalf(`Hello("Dan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

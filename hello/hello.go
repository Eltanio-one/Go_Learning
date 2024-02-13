// package main

// import (
// 	"fmt"
// 	"log"

// 	"example.com/greetings"
// )

// func main() {
// 	// set the properties of the predefined logger, including:
// 	// the log entry prefix and a flag to disable printing
// 	// the time, source file and line number
// 	log.SetPrefix("greetings: ")
// 	log.SetFlags(0)

// 	// a slice of names to pass to Hellos
// 	names := []string{"Dan", "Charlotte", "Warren"}

// 	// request greeting messages for the names
// 	messages, err := greetings.Hellos(names)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// get a greeting message and print
// 	// message, err := greetings.Hello("Warren")
// 	// if an error was return, print it to the console and exit the program

// 	// if no error, then print the returned message
// 	fmt.Println(messages)
// }

package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!", "ligma")
}

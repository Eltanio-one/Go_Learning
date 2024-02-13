package main

import (
	"fmt"
	"log"
	"readinputstring"
)

func main() {
	input, err := readinputstring.ReadInputString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

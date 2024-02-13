// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// ask for input but keep the cursor on the same line not using Println
// 	fmt.Print("Enter a grade: ")
// 	// create a bufio.Reader that will read the users standard input and store
// 	// this in the reader variable
// 	reader := bufio.NewReader(os.Stdin)
// 	// create an input variable that stores what the user typed as a string
// 	// until enter was pressed (a new line was made using the \n rune)
// 	input, err := reader.ReadString('\n')
// 	// print input
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	input = strings.TrimSpace(input)
// 	grade, err := strconv.ParseFloat(input, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var status string
// 	if grade >= 60 {
// 		status = "Pass!"
// 	} else {
// 		status = "Fail!"
// 	}
// 	fmt.Println("A grade of", grade, "is a", status)
// }

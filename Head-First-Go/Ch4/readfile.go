// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	// declare the array we will populate
// 	var numbers []float64
// 	// attempt to open the text file into the variable "file"
// 	// this returns a pointer to the opened file, file and an error if generated
// 	file, err := os.Open("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// initialise a scanner that is set to scan the file
// 	scanner := bufio.NewScanner(file)
// 	// iterate through the lines in the scanner and print the text on each line
// 	for scanner.Scan() {
// 		// parse the float
// 		number, err := strconv.ParseFloat(scanner.Text(), 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// append to the numbers slice
// 		numbers = append(numbers, number)
// 	}
// 	// close the file, will return nil if successful
// 	err = file.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// if there was an error when the scanner tried to scan, then report it
// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}
// 	// print array
// 	fmt.Println(numbers)
// }

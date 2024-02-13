package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	// seed the random number generator, this will change each time
	// the main function is run as the time will have changed
	rand.Seed(seconds)
	// create target variable and assign it a number between 1 - 100
	// as rand.Intn will generate a number between 0 and the number passed
	target := rand.Intn(100) + 1
	var count int
	for count = 10; count > 0; count-- {
		fmt.Println("The count is", count)
		// get the current date and time as an integer
		// the Unix() function will convert the Time value
		// into a Unix time format which is an integer with the number
		// of seconds since jan 1 1970.
		fmt.Print("Guess a number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess == target {
			fmt.Println("You win!")
			break
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		}
	}
	if count == 0 {
		fmt.Println("You lose!")
	}
}

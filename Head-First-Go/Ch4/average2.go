package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	floats := os.Args[1:]
	var sum float64
	for _, numb := range floats {
		number, err := strconv.ParseFloat(numb, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("Average: %0.2f\n", sum/float64(len(floats)))
}

func ave(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))

}

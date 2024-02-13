package main

import (
	"fmt"
)

func area(width float64, height float64) (area float64, err error) {
	if width <= 0 || height <= 0 {
		return 0, fmt.Errorf("no negative values allowed")
	}
	area = width * height
	return area / 10.0, nil
}

func main() {
	var amount, total float64
	amount, err := area(4.2, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f litres needed\n", amount)
	}
	total += amount
	amount, err = area(5.2, 3.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f litres needed\n", amount)
	}
	total += amount
	fmt.Printf("Total: %.2f litres\n", total)
}

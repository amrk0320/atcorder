package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	product := a * b
	if remainder := product % 2; remainder != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}

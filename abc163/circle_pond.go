package main

import (
	"fmt"
)

func main() {
	var r float32
	fmt.Scanf("%f", &r)
	const pie float32 = 3.14
	fmt.Printf("%f\n", 2*r*pie)
}

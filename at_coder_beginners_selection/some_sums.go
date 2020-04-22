package main

import (
	"fmt"
	"strconv"
)

func someDigitsSum(n int) int {
	var sum int
	s := strconv.Itoa(n)
	for _, r := range s {
		i := string(r)
		if num, err := strconv.Atoi(i); err == nil {
			sum += num
		}
	}
	return sum
}

func main() {
	var sums int
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	for i := 1; i <= n; i++ {
		sum := someDigitsSum(i)
		if a <= sum && sum <= b {
			sums += i
		}
	}
	fmt.Println(sums)
}

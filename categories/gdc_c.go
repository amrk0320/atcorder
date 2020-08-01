package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	var i int
	if sc.Scan() {
		if num, err := strconv.Atoi(sc.Text()); err == nil {
			i = num
		}
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	var steps = []int{}
	n, x := nextInt(), nextInt()
	for i := 0; i < n; i++ {
		city := nextInt()
		steps = append(steps, absTwo(city, x))
	}
	fmt.Println(allGcd(steps))
}

func absTwo(a, b int) int {
	v := a - b
	return int(math.Abs(float64(v)))
}

func allGcd(slice []int) int {
	a := slice[0]
	for i := 0; i < len(slice)-1; i++ {
		b := slice[i+1]
		a = gcd(a, b)
	}
	return a
}

func gcd(a, b int) int {
	if a%b != 0 {
		return gcd(b, a%b)
	} else {
		return b
	}
}

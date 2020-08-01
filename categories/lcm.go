package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt64() int64 {
	var i int64
	if sc.Scan() {
		if num, err := strconv.ParseInt(sc.Text(), 10, 64); err == nil {
			i = num
		}
	}
	return i
}

func nextInt() int {
	var i int
	if sc.Scan() {
		if num, err := strconv.Atoi(sc.Text()); err == nil {
			i = num
		}
	}
	return i
}

func allLcm(slice []int64) int64 {
	a := slice[0]
	for i := 0; i < len(slice)-1; i++ {
		b := slice[i+1]
		a = lcm(a, b)
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func gcd(a, b int64) int64 {
	if a%b != 0 {
		return gcd(b, a%b)
	} else {
		return b
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	var clocks = []int64{}
	n := nextInt()
	for i := 0; i < n; i++ {
		t := nextInt64()
		clocks = append(clocks, t)
	}
	fmt.Println(allLcm(clocks))
}

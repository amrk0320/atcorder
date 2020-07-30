package main

import (
	"bufio"
	"fmt"
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
	n := nextInt()
	var point = map[int]int{}
	for i := 0; i < n; i++ {
		point[i+1] = nextInt()
	}
	var min int
	var initialized bool
	for p := 0; p < 100; p++ {
		mtg := p + 1
		var currMin int
		for i := 0; i < len(point); i++ {
			diff := point[i+1] - mtg
			diff *= diff
			currMin += diff
		}
		if !initialized && min == 0 {
			min = currMin
			initialized = true
		}
		if currMin < min {
			min = currMin
		}
	}
	fmt.Println(min)
}

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
	var set = map[int]bool{}
	var a = []int{}
	var n = nextInt()
	for i := 0; i < n; i++ {
		ai := nextInt()
		set[ai] = true
		a = append(a, ai)
	}
	if len(set) == len(a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

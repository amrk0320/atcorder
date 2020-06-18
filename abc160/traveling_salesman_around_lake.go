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

func abs(i int) int {
	if i < 0 {
		i = i * -1
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	k, n := nextInt(), nextInt()
	var maxLoad int
	var firstHome = nextInt()
	var preHome = firstHome

	for i := 1; i <= n-1; i++ {
		nowHome := nextInt()
		load := nowHome - preHome
		if maxLoad <= load {
			maxLoad = load
		}
		preHome = nowHome
	}
	topLoad := (k - preHome) + firstHome
	if maxLoad <= topLoad {
		maxLoad = topLoad
	}
	fmt.Println(k - maxLoad)
}

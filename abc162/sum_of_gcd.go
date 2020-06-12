package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
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

func gcd(m, n int) int {
	
}

func threeGcd(a, b, c int) int {
	r := a%b
	if 
}

var sum int

func dfs(a, b, c int) {
	for x := 1; x <= c; x++ {
		for y := 1; x <= b; y++ {
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	k := nextInt()
	dfs(k, k, k)
	fmt.Println(sum)
}

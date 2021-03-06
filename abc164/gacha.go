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

func nextString() string {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	var gachaMap = map[string]bool{}
	n := nextInt()
	for i := 0; i < n; i++ {
		s := nextString()
		if 1 <= len(s) && len(s) <= 10 {
			gachaMap[s] = true
		}
	}
	fmt.Println(len(gachaMap))
}

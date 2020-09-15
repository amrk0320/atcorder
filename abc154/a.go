package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
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

func main() {
	sc.Split(bufio.ScanWords)
	s, _ := nextString(), nextString()
	a, b := nextInt(), nextInt()
	u := nextString()
	if u == s {
		fmt.Printf("%d %d\n", a-1, b)
	} else {
		fmt.Printf("%d %d\n", a, b-1)
	}
}

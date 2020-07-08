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
	if n%1000 == 0 {
		fmt.Println(0)
	} else {
		b := n/1000 + 1
		c := b*1000 - n
		fmt.Println(c)
	}
}

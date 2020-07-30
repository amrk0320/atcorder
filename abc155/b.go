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
	for i := 0; i < n; i++ {
		a := nextInt()
		if a%2 == 0 {
			if a%3 == 0 {
				continue
			}
			if a%5 == 0 {
				continue
			}
			fmt.Println("DENIED")
			return
		}
	}
	fmt.Println("APPROVED")
}

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
	var employees = nextInt()
	var subordinateNum = map[int]int{}
	for empNum := 1; empNum <= employees; empNum++ {
		subordinateNum[empNum] = 0
	}
	for empNum := 2; empNum <= employees; empNum++ {
		managerNo := nextInt()
		subordinateNum[managerNo]++
	}
	for empNum := 1; empNum <= employees; empNum++ {
		fmt.Println(subordinateNum[empNum])
	}
}

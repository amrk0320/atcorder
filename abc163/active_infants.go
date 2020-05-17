package main

import (
	"bufio"
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
	var childCount = nextInt()
	var activeInfants []int
	for n := 1; n <= childCount; n++ {
		activeInfants = append(activeInfants, nextInt())
	}

}

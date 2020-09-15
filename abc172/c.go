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
	var reads, maxA, maxB int
	var a = map[int]int{}
	var b = map[int]int{}
	var n, m, k = nextInt(), nextInt(), nextInt()
	for i := 0;i < n; i++{
		a[i+1] = nextInt() + a[i]
	}
	for i := 0;i < m; i++{
		b[i+1] = nextInt() + b[i]
	}
	for i := 0; i < n;i++{
		if k <= a[i+1] {
			break
		}
		maxA = i + 1
	}
	for i := 0;  i < m;i++{
		if k <= b[i+1] {
			break
		}
		maxB = i + 1
	}
	if maxA < maxB {
		reads += maxB
		for i := 0;i < n; i++{
			if b[maxB] + a[i+1] <= k {
				reads++
			} else {
				fmt.Println(reads)
				return
			}
		}
	} else {
		reads += maxA
		for i := 0;i < m; i++{
			if a[maxA] + b[i+1] <= k {
				reads++
			} else {
				fmt.Println(reads)
				return
			}
		}
	}
	fmt.Println(reads)
}

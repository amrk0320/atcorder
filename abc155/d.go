package main

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)
	var values = []int64{}
	var q = []int{}
	n, k := nextInt(), nextInt()
	for i := 0; i < n; i++ {
		q = append(q, nextInt())
	}
	sort.Ints(q)
	var count int
	for {
		if len(q) < 2 {
			break
		}
		top := q[0]
		culc := (count + (len(q) - 1))
		if k <= culc {
			for i := 1; i < len(q); i++ {
				v := q[i]
				mul := int64(top * v)
				values = append(values, mul)
				count++
				if count == k {
					fmt.Println(mul)
					return
				}
			}
		}
		count += len(q) - 1
		q = q[1:]
	}
}

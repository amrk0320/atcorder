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

func nextString() string {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	var outStrings = map[string]int{}
	var max int

	for i := 0; i < n; i++ {
		s := nextString()
		outStrings[s]++
		count := outStrings[s]
		if max < count {
			max = count
		}
	}
	slice := []string{}
	for k, v := range outStrings {
		if v == max {
			slice = append(slice, k)
		}
	}
	sort.Strings(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

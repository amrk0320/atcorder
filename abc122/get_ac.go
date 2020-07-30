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
	_, loop := nextInt(), nextInt()
	s := nextString()
	var cSum = map[int]int{}
	for i := 0; i <= len(s); i++ {
		str := s[:i]
		if i == 0 {
			cSum[i] = 0
		} else {
			var count int
			for j := 0; j < len(str); j++ {
				if 0 < j {
					if string(str[j]) == "C" && string(str[j-1]) == "A" {
						count++
					}
				}
			}
			cSum[i] = count
		}
	}
	for i := 0; i < loop; i++ {
		var times int64
		l, r := nextInt(), nextInt()
		times = int64(cSum[r]) - int64(cSum[l])
		fmt.Println(times)
	}
}

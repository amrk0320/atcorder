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

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	var allSunukeCount, okahiCount = nextInt(), nextInt()
	var okashiSunukes []int
	for i := 0; i < okahiCount; i++ {
		sunukeCount := nextInt()
		for i := 0; i < sunukeCount; i++ {
			sunukeNum := nextInt()
			if !contains(okashiSunukes, sunukeNum) {
				okashiSunukes = append(okashiSunukes, sunukeNum)
			}
		}
	}
	fmt.Println(allSunukeCount - len(okashiSunukes))
}

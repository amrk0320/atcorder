package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func main() {
	sc.Split(bufio.ScanWords)
	var gachaMap = map[string]bool{}
	n := nextInt()
	for i := 0; i < n; i++ {
		gachaMap[nextString()] = true
	}
	fmt.Println(len(gachaMap))
}

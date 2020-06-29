package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func reverseS(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func lastAppend(s []string, str string) []string {
	return append(s, str)
}

func firstAppend(s []string, str string) []string {
	array := []string{str}
	return append(array, s...)
}

func main() {
	sc.Split(bufio.ScanWords)
	s := []string{nextString()}
	q := nextInt()
	var reverse bool = false

	for i := 0; i < q; i++ {
		t := nextInt()
		if t == 1 {
			reverse = !reverse
		} else {
			f := nextInt()
			c := nextString()
			if f == 1 {
				// 先頭
				if reverse == false {
					s = firstAppend(s, c)
				} else {
					s = lastAppend(s, c)
				}
			} else {
				// 末尾
				if reverse == false {
					s = lastAppend(s, c)
				} else {
					s = firstAppend(s, c)
				}
			}
		}
	}
	if reverse == false {
		fmt.Println(strings.Join(s, ""))
	} else {
		fmt.Println(reverseS(strings.Join(s, "")))
	}
}

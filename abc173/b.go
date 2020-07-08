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
	var ac, wa, tle, re int
	for i := 0; i < n; i++ {
		c := nextString()
		if c == "AC" {
			ac++
		} else if c == "WA" {
			wa++
		} else if c == "TLE" {
			tle++
		} else if c == "RE" {
			re++
		}
	}
	fmt.Println("AC x " + strconv.Itoa(ac))
	fmt.Println("WA x " + strconv.Itoa(wa))
	fmt.Println("TLE x " + strconv.Itoa(tle))
	fmt.Println("RE x " + strconv.Itoa(re))
}

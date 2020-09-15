package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()
	var result string
	for range s {
		result += "x"
	}
	fmt.Println(result)
}

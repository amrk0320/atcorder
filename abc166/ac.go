package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextText() string {
	var i string
	if sc.Scan() {
		i = sc.Text()
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	var s = nextText()
	if s == "ABC" {
		fmt.Println("ARC")
	} else {
		fmt.Println("ABC")
	}
}

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
	n := nextString()
	for _, str := range n {
		if string([]rune{str}) == "7" {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

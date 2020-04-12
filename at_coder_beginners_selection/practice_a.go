package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var sum int
var s string

func main() {
	for sc.Scan() {
		a := strings.Split(sc.Text(), " ")
		for _, item := range a {
			if num, err := strconv.Atoi(item); err == nil {
				sum += num
			} else {
				s = item
			}
		}
	}
	fmt.Printf("%d %s\n", sum, s)
}

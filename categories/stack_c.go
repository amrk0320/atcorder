package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func isSingelColor(nokori string) bool {
	return strings.Count("1", nokori) == 0 || strings.Count("1", nokori) == 1
}

func search(s string) int {
	var count int
	var stack []string
	for _, r := range s {
		cube := string([]rune{r})
		if len(stack) == 0 {
			stack = append(stack, cube)
		} else if stack[len(stack)-1] != cube {
			count += 2
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, cube)
		}
	}
	nokori := strings.Join(stack, "")
	if len(nokori) == 0 || isSingelColor(nokori) {
		return count
	} else {
		return count + search(nokori)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()
	fmt.Println(search(s))
}

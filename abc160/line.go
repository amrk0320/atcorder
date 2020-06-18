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

func main() {
	sc.Split(bufio.ScanWords)
	n, x, y := nextInt(), nextInt(), nextInt()
	var lines = map[int]int{}
	var shortcut = (x - 1) + (n - y) + 1
	var shortcutPoint = shortcut + 1
	var mountain = y - x

	for k := 1; k <= n-1; k++ {
		lines[k] = n - k
		if k <= shortcut {
			lines[k] += shortcutPoint - k
			if k <= x {
				lines[k] -= x - k
			}
			if k <= n-y {
				lines[k] -= (n - y) - k
			}
			if k <= mountain {
				lines[k] -= mountain + 1 - k
			}
		}
		if shortcut < k {
			lines[k] = 0
		}
	}

	for k := 1; k <= len(lines); k++ {
		fmt.Println(lines[k])
	}
}

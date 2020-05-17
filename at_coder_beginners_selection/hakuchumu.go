package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scan() string {
	var sc = bufio.NewScanner(os.Stdin)
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func join(s []string) string {
	return strings.Join(s, "")
}

func insertT(s string) string {
	phrases := []string{"dream", "dreamer", "erase", "eraser"}
	var t []string

	for _, p := range phrases {
		if strings.Index(s, p) != 0 {
			continue
		}
		t = append(t, p)
		if s == join(t) || len(s) < len(t) {
			return join(t)
		}
		tSufix := insertT(strings.TrimPrefix(s, p))
		t = append(t, tSufix)
		if s == join(t) {
			return join(t)
		} else {
			t = []string{}
		}
	}
	return join(t)
}

func main() {
	var s = scan()

	t := insertT(s)

	if s == t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

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

func insertT(s string) []string {
	phrases := []string{"dream", "dreamer", "erase", "eraser"}
	var t []string

	for _, v := range phrases {
		if strings.Index(s, v) == 0 {
			trimed := strings.TrimPrefix(s, v)
			t = append(t, v)
			if trimed == "" {
				return t
			} else {
				secondT := insertT(trimed)
				if 0 < len(secondT) {
					t = append(t, secondT...)
				}
			}
			if join(t) == s {
				return t
			} else {
				t = []string{}
			}
		}
	}
	return t
}

func main() {
	var s = scan()
	if len(s) < 1 {
		fmt.Println("NO")
		return
	}

	t := insertT(s)

	if s == join(t) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

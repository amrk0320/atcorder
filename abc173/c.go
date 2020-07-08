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
	var blackMap = map[int]map[int]int{
		1: map[int]int{},
		2: map[int]int{},
		3: map[int]int{},
		4: map[int]int{},
		5: map[int]int{},
		6: map[int]int{},
	}

	sc.Split(bufio.ScanWords)
	height, width, k := nextInt(), nextInt(), nextInt()
	var count int
	for i := 0; i < height; i++ {
		mark := nextString()
		for j, c := range mark {
			if string([]rune{c}) == "#" {
				blackMap[i+1][j+1] = 1
			}
		}
	}

	n := width + height
	// 11111
	for bits := 0; bits < (1 << uint64(n)); bits++ {
		var heightMap = make(map[int]bool)
		var widthMap = make(map[int]bool)
		var kcount int
		for i := 0; i < n; i++ {
			if (bits>>uint64(i))&1 == 1 {
				if i < height {
					heightMap[i+1] = true
				} else {
					widthMap[i+1-height] = true
				}
			}
		}
		for i := 0; i < height; i++ {
			if heightMap[i+1] == true {
				continue
			}
			for j := 0; j < width; j++ {
				if widthMap[j+1] == true {
					continue
				}
				// [height][width]
				if blackMap[i+1][j+1] == 1 {
					kcount++
				}
			}
		}
		if kcount == k {
			count++
		}
	}
	fmt.Println(count)
}

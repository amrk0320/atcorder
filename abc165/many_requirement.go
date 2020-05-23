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

func maxAmap(aLength int, maxA int) map[int]int {
	for n := 1; n <= aLength; n++ {
		aMap[n] = maxA
	}
	return aMap
}

func addPoint(a int, b int, c int, d int, points int) int {
	if c == (aMap[b] - aMap[a]) {
		points += d
	}
	return points
}

func updateMaxPoint(aMap map[int]int, currentMaxPoint int) int {
	addPoint
	return currentMaxPoint
}

var aMap = map[int]int{}
var maxPoint int

func main() {
	sc.Split(bufio.ScanWords)
	aLength, maxA, tryCount := nextInt(), nextInt(), nextInt()

	aMap = maxAmap(aLength, maxA)

	for i := 1; i <= tryCount; i++ {
		a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	}

	// make A
	for maxValue := maxA; 1 < maxValue; maxValue-- {
		for reduceOffset := 1; reduceOffset <= aLength; reduceOffset++ {
			aMap[reduceOffset] = aMap[reduceOffset] - 1
			updateMaxPoint(aMap, currentMaxPoint)
		}
	}

	fmt.Println(currentMaxPoint)
}

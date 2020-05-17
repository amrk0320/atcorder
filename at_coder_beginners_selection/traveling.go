package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func scanN() int {
	var sc = bufio.NewScanner(os.Stdin)
	var n string
	if sc.Scan() {
		n = sc.Text()
	}
	if i, err := strconv.Atoi(n); err == nil {
		return i
	} else {
		return -1
	}
}

func scanT(n int) [][]int {
	var sc = bufio.NewScanner(os.Stdin)
	var t [][]int
	for i := 1; i <= n && sc.Scan(); i++ {
		v := sc.Text()
		plan := []int{}
		for _, s := range strings.Split(v, " ") {
			if num, err := strconv.Atoi(s); err == nil {
				plan = append(plan, num)
			}
		}
		t = append(t, plan)
	}
	return t
}

func walk(xStep int, yStep int, time int) bool {
	x := math.Abs(float64(xStep))
	y := math.Abs(float64(yStep))
	remain := time - int(x+y)
	if 0 <= remain && remain%2 == 0 {
		return true
	} else {
		return false
	}
}

func traveling(t [][]int) bool {
	var currentTime int
	var x int
	var y int
	for _, plan := range t {
		time := plan[0] - currentTime
		xGoal := plan[1]
		yGoal := plan[2]
		ok := walk(xGoal-x, yGoal-y, time)
		if ok {
			currentTime = plan[0]
			x = xGoal
			y = yGoal
		} else {
			return false
		}
	}
	return true
}

func main() {
	var n int = scanN()
	var t [][]int = scanT(n)

	ok := traveling(t)

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

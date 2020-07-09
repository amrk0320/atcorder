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
	// [y][x]
	var card = map[int]map[int]int{
		1: map[int]int{},
		2: map[int]int{},
		3: map[int]int{},
	}
	var result = map[int]map[int]bool{
		1: map[int]bool{},
		2: map[int]bool{},
		3: map[int]bool{},
	}
	for i := 0; i < 3; i++ {
		card[1][i+1] = nextInt()
	}
	for i := 0; i < 3; i++ {
		card[2][i+1] = nextInt()
	}
	for i := 0; i < 3; i++ {
		card[3][i+1] = nextInt()
	}

	nums := nextInt()
	for i := 0; i < nums; i++ {
		num := nextInt()
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				if card[x+1][y+1] == num {
					result[x+1][y+1] = true
				}
			}
		}
	}

	for y := 0; y < 3; y++ {
		var ok int
		for x := 0; x < 3; x++ {
			if result[x+1][y+1] == true {
				ok++
			}
		}
		if ok == 3 {
			fmt.Println("Yes")
			return
		}
	}
	for x := 0; x < 3; x++ {
		var ok int
		for y := 0; y < 3; y++ {
			if result[x+1][y+1] == true {
				ok++
			}
		}
		if ok == 3 {
			fmt.Println("Yes")
			return
		}
	}

	if result[1][1] == true && result[2][2] == true && result[3][3] {
		fmt.Println("Yes")
		return
	}
	if result[1][3] == true && result[2][2] == true && result[3][1] {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

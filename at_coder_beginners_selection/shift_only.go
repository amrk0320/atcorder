package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allItemEqualToEven(array []int) bool {
	for _, num := range array {
		if num%2 != 0 {
			return false
		}
	}
	return true
}

func divideTwoAllItem(array []int) []int {
	var dividedArray []int
	for _, item := range array {
		dividedItem := item / 2
		dividedArray = append(dividedArray, dividedItem)
	}
	return dividedArray
}

func scanBlackBorad() bool {
	var blackBoradSize int
	fmt.Scan(&blackBoradSize)
	if 0 < blackBoradSize && sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		// 整数値のslice
		for _, item := range inputs {
			if num, err := strconv.Atoi(item); err == nil {
				blackBorad = append(blackBorad, num)
			}
		}
	}
	if 0 < len(blackBorad) {
		return true
	} else {
		return false
	}
}

var sc = bufio.NewScanner(os.Stdin)
var tryCount int
var blackBorad []int

func main() {
	var scanSuccess bool = scanBlackBorad()
	var try bool
	if scanSuccess {
		try = true
	}
	for try {
		even := allItemEqualToEven(blackBorad)
		if even {
			blackBorad = divideTwoAllItem(blackBorad)
			try = true
			tryCount++
		} else {
			try = false
		}
	}
	fmt.Println(tryCount)
}

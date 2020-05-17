package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int
var a []int
var alice int
var bob int
var diff int

func scan() (int, []int) {
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		v := sc.Text()
		if i, err := strconv.Atoi(v); err == nil {
			n = i
		}
	}
	if sc.Scan() {
		v := sc.Text()
		array := strings.Fields(v)
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(array[i])
			a = append(a, num)
		}
	}
	return n, a
}

func selectCards(a []int) (int, int) {
	for i := range a {
		if i%2 == 0 {
			selectMaxPointCard(&a, &alice)
		} else {
			selectMaxPointCard(&a, &bob)
		}
	}
	return alice, bob
}

func selectMaxPointCard(a *[]int, person *int) {
	var max int
	var n int
	for i, v := range *a {
		if max <= v {
			max = v
			n = i
		}
	}
	*person += max
	*a = append((*a)[:n], (*a)[n+1:]...)
}

func main() {
	n, a = scan()
	alice, bob = selectCards(a)
	fmt.Println(alice - bob)
}

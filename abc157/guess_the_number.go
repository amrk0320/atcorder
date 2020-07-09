package main

import (
	"bufio"
	"fmt"
	"math"
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
	digits, loop := nextInt(), nextInt()
	var num = make(map[int]string)

	if loop == 0 {
		if digits == 1 {
			fmt.Println(0)
			return
		} else if digits == 2 {
			fmt.Println(10)
			return
		} else if digits == 3 {
			fmt.Println(100)
			return
		}
	}

	for i := 0; i < loop; i++ {
		s, c := nextInt(), nextInt()
		if 1 < digits && s == 1 && c == 0 {
			fmt.Println(-1)
			return
		}
		if digits < s {
			fmt.Println(-1)
			return
		}
		str := strconv.Itoa(c)
		if num[s] == "" || num[s] == str {
			num[s] = str
		} else {
			fmt.Println(-1)
			return
		}
	}
	var ans int
	for i := 0; i < digits; i++ {
		digit := digits - i
		if num[i+1] != "" {
			val, _ := strconv.Atoi(num[i+1])
			ans += int(math.Pow(float64(10), float64(digit-1))) * val
		} else if num[i+1] == "" && i == 0 {
			ans += int(math.Pow(float64(10), float64(digit-1)))
		}
	}
	fmt.Println(ans)
}

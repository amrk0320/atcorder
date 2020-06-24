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

// O(N^2)
// 配列Aを作成
// 数値に対して、値の出現数を記録したハッシュテーブルを作成
// ハッシュテーブルの値に対して、組み合わせ数を記録したハッシュテーブルを作成

func combinationTwo(a int) int64 {
	var n int64 = int64(a)
	return n * (n - 1) / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	var a = make([]int, n)
	var values = make(map[int]int)
	var combinations = make(map[int]int64)
	var all int64
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		values[a[i]]++
	}
	for k, v := range values {
		combi := combinationTwo(v)
		combinations[k] = combi
		all += combi
	}
	for i := 0; i < n; i++ {
		k := a[i]
		balls := values[k] - 1
		combi := combinations[k]
		newCombi := combinationTwo(balls)
		pattern := all + newCombi - combi
		fmt.Println(pattern)
	}
}

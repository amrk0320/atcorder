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

// 0.08かけして切り捨ててA円
// 0.1かけして切り捨ててB円
// 0.1でわると1桁目が切り捨てられる
// 0.08でわると整数部が税額
// 1. 1 から1250までループ
// 2. ループ値を0.08でわった整数部がaと一致するものを探す
// 3. その値を ループ値を0.1かけした整数部が一致すればその値を出力して終了

func main() {
	sc.Split(bufio.ScanWords)
	a, b := nextInt(), nextInt()
	for i := 0; i < 1250; i++ {
		noTaxPrice := i + 1
		taxEight := float32(noTaxPrice) * 0.08
		if int(taxEight) == a {
			if int(float32(noTaxPrice)*0.1) == b {
				fmt.Println(noTaxPrice)
				return
			}
		}
	}
	fmt.Println(-1)
}

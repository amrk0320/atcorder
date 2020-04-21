package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var combinationCount int

func scanCoinsAndPrice(sc *bufio.Scanner) (int, int, int, int) {
	var coins []int
	for i := 0; sc.Scan() && i < 4; i++ {
		v, _ := strconv.Atoi(sc.Text())
		coins = append(coins, v)
	}
	var fiveHundredCoin int = coins[0]
	var oneHundredCoin int = coins[1]
	var fiftyCoin int = coins[2]
	var price int = coins[3]
	return fiveHundredCoin, oneHundredCoin, fiftyCoin, price
}

func countCombinationByFiftyCoin(fiftyCoin int, price int) {
	if price <= fiftyCoin*50 {
		combinationCount++
	}
}

func countCombinationByOneHundredAndFiftyCoin(oneHundredCoin int, fiftyCoin int, price int) {
	for i := 0; i <= oneHundredCoin; i++ {
		balance := price - i*100
		if 0 < balance {
			countCombinationByFiftyCoin(fiftyCoin, balance)
		} else if balance == 0 {
			combinationCount++
		}
	}
}

func countCombinationByAllCoins(fiveHundredCoin int, oneHundredCoin int, fiftyCoin, price int) {
	for i := 0; i <= fiveHundredCoin; i++ {
		balance := price - i*500
		if 0 < balance {
			countCombinationByOneHundredAndFiftyCoin(oneHundredCoin, fiftyCoin, balance)
		} else if balance == 0 {
			combinationCount++
		}
	}
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	fiveHundredCoin, oneHundredCoin, fiftyCoin, expectSumPrice := scanCoinsAndPrice(sc)
	countCombinationByAllCoins(fiveHundredCoin, oneHundredCoin, fiftyCoin, expectSumPrice)
	fmt.Println(combinationCount)
}

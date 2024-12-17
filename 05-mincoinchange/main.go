package main

import (
	"fmt"
	"labalg/mincoinchange"
)

func main() {
	amount := 36
	coins := []int{1, 5, 10, 25}
	result := mincoinchange.MinCoinChange(coins, amount)
	fmt.Println("result: ", result)
}

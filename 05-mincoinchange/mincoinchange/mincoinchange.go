package mincoinchange

func MinCoinChange(coins []int, amount int) []int {
	cache := make(map[int][]int)
	return makeChange(amount, cache, coins)
}

func makeChange(amount int, cache map[int][]int, coins []int) []int {
	if amount <= 0 {
		return nil
	}

	amountCache, ok := cache[amount]
	if ok && len(amountCache) > 0 {
		return cache[amount]
	}

	min := []int{}
	newMin := []int{}
	newAmount := 0

	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		newAmount = amount - coin

		if newAmount >= 0 {
			newMin = makeChange(newAmount, cache, coins)
		}

		if newAmount >= 0 &&
			(len(newMin) < len(min)-1 || len(min) == 0) &&
			(len(newMin) > 0 || newAmount == 0) {
			min = append(newMin, coin)
		}
	}

	cache[amount] = min
	return cache[amount]
}

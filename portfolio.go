package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(getMaxProfit([]int{-4, -9, 3, 4, -2}))
}

func getMaxProfit(profits []int) int {
	sort.Slice(profits, func(i, j int) bool {
		return profits[i] > profits[j]
	})
	maxs := int(math.Inf(-1))
	for priority := len(profits); priority >= 1; priority-- {
		sum := 0
		p := priority
		for i := 0; i < priority; i++ {
			sum += profits[i] * p
			p--
		}
		if sum > maxs {
			maxs = sum
		}
	}
	if maxs < 0 {
		return 0
	}
	return maxs
}

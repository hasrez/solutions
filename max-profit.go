package main

import "fmt"

func main() {
	fmt.Println(calc(2, []int{3, 2, 6, 5, 0, 3}))
}

// todo for DP
var dp [][][]int

// todo for DP

func calc(k int, prices []int) int {
	if k > len(prices)/2 {
		k = len(prices) / 2
	}
	profits := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		var diffs []int
		for j := 0; j < len(prices); j++ {
			diffs = append(diffs, prices[j]-prices[i])
		}
		profits[i] = diffs
	}
	// todo for DP
	dp = make([][][]int, k+1)
	for s := 0; s <= k; s++ {
		dp[s] = make([][]int, len(prices))
		for i := 0; i < len(prices); i++ {
			var diffs []int
			for j := 0; j < len(prices); j++ {
				diffs = append(diffs, -1)
			}
			dp[s][i] = diffs
		}
	}
	// todo for DP
	maxp := 0
	for i := 0; i < k; i++ {
		p := getMaxProfit(0, 1, profits, i+1)
		if p > maxp {
			maxp = p
		}
	}
	return maxp
}

func getMaxProfit(i, j int, profits [][]int, state int) int {
	m := len(profits)
	// todo for DP
	l, h := i, j
	// todo for DP
	if i >= m || j >= m || state == 0 {
		return 0
	}
	// todo for DP
	if dp[state][i][j] > 0 {
		return dp[state][i][j]
	}
	// todo for DP
	p := 0
	for i < m {
		pp := profits[i][j] + getMaxProfit(j, j, profits, state-1)
		if pp > p {
			p = pp
		}
		j++
		if j >= m {
			i++
			j = i
		}
	}
	// todo for DP
	dp[state][l][h] = p
	// todo for DP
	return p
}

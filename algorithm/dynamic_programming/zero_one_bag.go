package dynamic_programming

func ZeroOneBag(weights []int, values []int, W int) int {
	dp := make([][]int, len(weights)+1)
	for i := range W + 1 {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i < len(weights); i++ {
		for j := 1; j <= W; j++ {
			if j >= weights[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(weights)][W]
}

func ZeroOneBag2(weight []int, values []int, W int) int {
	n := len(weight)
	dp := make([]int, W+1)
	for i := 0; i < n; i++ {
		for j := W; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+values[i])
		}
	}
	return dp[W]
}

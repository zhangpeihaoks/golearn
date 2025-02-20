package dynamic_programming

import "fmt"

func arrPlus(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])
	fmt.Println(dp)
	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
		fmt.Println(dp)
	}
	return dp[len(dp)-1]
}

func countSavingWays(target int) int {
	coins := []int{1, 5}
	dp := make([]int, target+1)
	dp[0] = 1
	for amount := 1; amount <= target; amount++ {
		for _, coin := range coins {
			if coin <= amount {
				dp[amount] += dp[amount-coin]
			}
		}
	}
	return dp[target]
}

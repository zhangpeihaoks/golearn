package leetcode

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
//
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
//
// Related Topics 数组 动态规划 👍 3188 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	return rob1(nums)
}
func rob2(nums []int) int {
	n := len(nums)
	ans := make([]int, n+2)
	ans[0] = nums[0]
	ans[1] = max(nums[0], nums[1])
	for i := 0; i < n; i++ {
		ans[i+2] = max(ans[i+1], ans[i]+nums[i])
	}
	return ans[n+1]
}
func rob1(nums []int) int {
	cache := make([]int, len(nums))
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		ans := max(dfs(i-1), dfs(i-2)+nums[i])
		cache[i] = ans
		return ans
	}
	return dfs(len(nums) - 1)
}
func rob3(nums []int) int {
	data := make([][]int, len(nums))
	for i := range data {
		data[i] = make([]int, 2)
	}
	data[0][0] = 0
	data[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		data[i][0] = max(data[i-1][0], data[i-1][1])
		data[i][1] = data[i-1][0] + nums[i]
	}
	return max(data[len(nums)-1][0], data[len(nums)-1][1])
}

//leetcode submit region end(Prohibit modification and deletion)

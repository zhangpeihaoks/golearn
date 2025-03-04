package leetcode

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5548 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	return trap2(height)
}
func trap1(height []int) int {
	ans := 0
	n := len(height)
	prevMax := make([]int, n)
	prevMax[0] = height[0]
	suffixMax := make([]int, n)
	suffixMax[n-1] = height[n-1]
	prevIndex := 1
	sufIndex := n - 2
	for prevIndex < n && sufIndex >= 0 {
		prevMax[prevIndex] = max(height[prevIndex], prevMax[prevIndex-1])
		suffixMax[sufIndex] = max(suffixMax[sufIndex+1], height[sufIndex])
		prevIndex += 1
		sufIndex -= 1
	}

	for i := 0; i < n; i++ {
		ans += min(prevMax[i], suffixMax[i]) - height[i]
	}
	return ans
}

func trap2(height []int) int {
	ans := 0
	n := len(height)
	left, right := 0, n-1
	preMax := 0
	sufMax := 0
	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

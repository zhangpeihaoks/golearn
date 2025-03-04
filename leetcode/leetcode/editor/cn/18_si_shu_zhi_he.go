package leetcode

import (
	"slices"
)

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 双指针 排序 👍 2033 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for a := range n - 3 {
		x := nums[a]
		if a > 0 && x == nums[a-1] {
			continue
		}
		sum := x + nums[a+1] + nums[a+2] + nums[a+3]
		if sum > target {
			break
		}
		sum = x + nums[n-1] + nums[n-2] + nums[n-3]
		if sum < target {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			y := nums[b]
			if b > a+1 && y == nums[b-1] {
				continue
			}
			if x+y+nums[b+1]+nums[b+2] > target {
				break
			}
			if x+y+nums[n-1]+nums[n-2] < target {
				continue
			}
			c, d := b+1, n-1
			for c < d {
				sum = x + y + nums[c] + nums[d]
				if sum == target {
					ans = append(ans, []int{x, y, nums[c], nums[d]})
					for c++; c < d && nums[c] == nums[c-1]; c++ {

					}
					for d--; d > c && nums[d] == nums[d-1]; d-- {
					}
				}
				if sum > target {
					d--
				} else if sum < target {
					c++
				}
			}
		}

	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

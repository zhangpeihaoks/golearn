package leetcode

import (
	"math"
	"slices"
)

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 双指针 排序 👍 1701 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	minDiff := math.MaxInt
	n := len(nums)
	ans := 0
	for i, num := range nums[:n-2] {
		if i > 0 && num == nums[i-1] {
			continue
		}

		sum := num + nums[i+1] + nums[i+2]
		if sum > target {
			if sum-target < minDiff {
				ans = sum
			}
			break
		}

		sum = num + nums[n-1] + nums[n-2]
		if sum < target {
			if target-sum < minDiff {
				minDiff = target - sum
				ans = sum
			}
			continue
		}

		j, k := i+1, n-1
		for j < k {
			sum = num + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if sum < target {
				if target-sum < minDiff {
					minDiff = target - sum
					ans = sum
				}
				j++
			} else {
				if sum-target < minDiff {
					minDiff = sum - target
					ans = sum
				}
				k--
			}
		}

	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

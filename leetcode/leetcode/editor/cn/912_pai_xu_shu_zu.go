package leetcode

//给你一个整数数组 nums，请你将该数组升序排列。
//
// 你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//解释：数组排序后，某些数字的位置没有改变（例如，2 和 3），而其他数字的位置发生了改变（例如，1 和 5）。
//
//
// 示例 2：
//
//
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//解释：请注意，nums 的值不一定唯一。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -5 * 10⁴ <= nums[i] <= 5 * 10⁴
//
//
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 1173 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	//slices.Sort(nums)
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, i, j int) {
	if i >= j {
		return
	}
	l, r := i, j
	pivot := nums[(i+j)/2]
	for l <= r {
		for nums[l] < pivot {
			l++
		}
		for nums[r] > pivot {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	if i < r {
		quickSort(nums, i, r)
	}
	if l < j {
		quickSort(nums, l, j)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

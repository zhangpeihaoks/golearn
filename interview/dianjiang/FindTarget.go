package dianjiang

import "sort"

func FindTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	left, right := 0, len(nums)-1

	for left < right {
		currentSum := nums[left] + nums[right]
		if currentSum == target {
			result = append(result, []int{nums[left], nums[right]})
			// 跳过左边相同的元素
			for left < right && nums[left] == nums[left+1] {
				left++
			}
			left++
			// 跳过右边相同的元素
			for left < right && nums[right] == nums[right-1] {
				right--
			}
			right--
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return result
}

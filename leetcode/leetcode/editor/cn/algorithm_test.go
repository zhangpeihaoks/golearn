package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	theLen := lengthOfLongestSubstring("au")
	fmt.Println(theLen)
}

func TestRemoveDuplicates2(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 3}
	data := removeDuplicates2(arr)
	fmt.Println(data)
}

func TestMinimumSum(t *testing.T) {
	minimumSum(5, 4)
}

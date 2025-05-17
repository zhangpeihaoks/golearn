package leetcode

//给定 m 个数组，每个数组都已经按照升序排好序了。
//
// 现在你需要从两个不同的数组中选择两个整数（每个数组选一个）并且计算它们的距离。两个整数 a 和 b 之间的距离定义为它们差的绝对值 |a-b| 。
//
// 返回最大距离。
//
// 示例 1：
//
//
//输入：[[1,2,3],[4,5],[1,2,3]]
//输出：4
//解释：
//一种得到答案 4 的方法是从第一个数组或者第三个数组中选择 1，同时从第二个数组中选择 5 。
//
//
// 示例 2：
//
//
//输入：arrays = [[1],[1]]
//输出：0
//
//
//
//
// 提示：
//
//
// m == arrays.length
// 2 <= m <= 10⁵
// 1 <= arrays[i].length <= 500
// -10⁴ <= arrays[i][j] <= 10⁴
// arrays[i] 以 升序 排序。
// 所有数组中最多有 10⁵ 个整数。
//
//
//
//
// Related Topics 贪心 数组 👍 161 👎 0
import (
	"math"
)

func maxDistance(arrays [][]int) int {
	minNum, maxNum := math.MaxInt/2, math.MinInt/2
	ans := 0
	for _, array := range arrays {
		tmpMinNum, tmpMaxNum := array[0], array[len(array)-1]
		ans = max(ans, tmpMaxNum-minNum, maxNum-tmpMinNum)
		maxNum = max(maxNum, tmpMaxNum)
		minNum = min(minNum, tmpMinNum)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

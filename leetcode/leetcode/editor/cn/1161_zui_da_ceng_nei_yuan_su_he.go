package leetcode

//给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
//
// 请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,7,0,7,-8,null,null]
//输出：2
//解释：
//第 1 层各元素之和为 1，
//第 2 层各元素之和为 7 + 0 = 7，
//第 3 层各元素之和为 7 + -8 = -1，
//所以我们返回第 2 层的层号，它的层内元素之和最大。
//
//
// 示例 2：
//
//
//输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
//输出：2
//
//
//
//
// 提示：
//
//
// 树中的节点数在
// [1, 10⁴]范围内
//
// -10⁵ <= Node.val <= 10⁵
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 146 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) (ans int) {
	result := []int{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if level == len(result) {
			result = append(result, root.Val)
		} else {

			result[level] += root.Val
		}
		if root.Left != nil {
			dfs(root.Left, level+1)
		}
		if root.Right != nil {
			dfs(root.Right, level+1)
		}
	}
	dfs(root, 0)
	for index, data := range result {
		if data > result[ans] {
			ans = index
		}
	}
	return ans + 1
}

//leetcode submit region end(Prohibit modification and deletion)

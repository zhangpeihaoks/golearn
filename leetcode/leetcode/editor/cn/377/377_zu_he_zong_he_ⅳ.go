package leetcode

/**
<p>给你一个由 <strong>不同</strong> 整数组成的数组 <code>nums</code> ，和一个目标整数 <code>target</code> 。请你从 <code>nums</code> 中找出并返回总和为 <code>target</code> 的元素组合的个数。</p>

<p>题目数据保证答案符合 32 位整数范围。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3], target = 4
<strong>输出：</strong>7
<strong>解释：</strong>
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [9], target = 3
<strong>输出：</strong>0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
 <li><code>1 &lt;= nums.length &lt;= 200</code></li>
 <li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
 <li><code>nums</code> 中的所有元素 <strong>互不相同</strong></li>
 <li><code>1 &lt;= target &lt;= 1000</code></li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？</p>

<div><div>Related Topics</div><div><li>数组</li><li>动态规划</li></div></div><br><div><li>👍 1140</li><li>👎 0</li></div>
*/

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum4(nums []int, target int) int {
	memo := make([]int, target+1)
	for i := range memo {
		memo[i] = -1 // -1表示未计算
	}
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == 0 {
			return 1
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		for _, num := range nums {
			if num <= i {
				res += dfs(i - num)
			}
		}
		return
	}
	return dfs(target)
}

func combinationSum42(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

//leetcode submit region end(Prohibit modification and deletion)

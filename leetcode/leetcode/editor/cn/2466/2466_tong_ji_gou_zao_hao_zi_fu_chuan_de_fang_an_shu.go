package leetcode

/**
<p>给你整数&nbsp;<code>zero</code>&nbsp;，<code>one</code>&nbsp;，<code>low</code>&nbsp;和&nbsp;<code>high</code>&nbsp;，我们从空字符串开始构造一个字符串，每一步执行下面操作中的一种：</p>

<ul>
 <li>将&nbsp;<code>'0'</code>&nbsp;在字符串末尾添加&nbsp;<code>zero</code>&nbsp; 次。</li>
 <li>将&nbsp;<code>'1'</code>&nbsp;在字符串末尾添加&nbsp;<code>one</code>&nbsp;次。</li>
</ul>

<p>以上操作可以执行任意次。</p>

<p>如果通过以上过程得到一个 <strong>长度</strong>&nbsp;在&nbsp;<code>low</code> 和&nbsp;<code>high</code>&nbsp;之间（包含上下边界）的字符串，那么这个字符串我们称为&nbsp;<strong>好</strong>&nbsp;字符串。</p>

<p>请你返回满足以上要求的 <strong>不同</strong>&nbsp;好字符串数目。由于答案可能很大，请将结果对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>low = 3, high = 3, zero = 1, one = 1
<b>输出：</b>8
<b>解释：</b>
一个可能的好字符串是 "011" 。
可以这样构造得到："" -&gt; "0" -&gt; "01" -&gt; "011" 。
从 "000" 到 "111" 之间所有的二进制字符串都是好字符串。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>low = 2, high = 3, zero = 1, one = 2
<b>输出：</b>5
<b>解释：</b>好字符串为 "00" ，"11" ，"000" ，"110" 和 "011" 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
 <li><code>1 &lt;= low&nbsp;&lt;= high&nbsp;&lt;= 10<sup>5</sup></code></li>
 <li><code>1 &lt;= zero, one &lt;= low</code></li>
</ul>

<div><div>Related Topics</div><div><li>动态规划</li></div></div><br><div><li>👍 130</li><li>👎 0</li></div>
*/

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodStrings1(low int, high int, zero int, one int) int {
	const mod = 1_000_000_007
	dp := make([]int, high+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 1
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return 1
		}
		p := &dp[i]
		if *p == -1 {
			*p = (dfs(i-zero) + dfs(i-one)) % mod
		}
		return *p
	}
	count := 0
	for i := low; i <= high; i++ {
		count += dfs(i)
	}
	return count % mod
}

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	const mod = 1_000_000_007
	ans := 0
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = dp[i-zero]
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
		if i >= low {
			ans = (ans + dp[i]) % mod
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

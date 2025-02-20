package leetcode

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10601 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	start, maxLen := 0, 0           // 初始化窗口起点和最大长度
	lastIndex := make(map[byte]int) // 哈希表记录字符最后出现的位置

	for i := 0; i < len(s); i++ {
		c := s[i]
		// 如果字符已存在且在窗口内，移动窗口起点
		if idx, ok := lastIndex[c]; ok && idx >= start {
			start = idx + 1
		}
		lastIndex[c] = i // 更新字符最后位置
		// 计算当前窗口长度并更新最大值
		if currentLen := i - start + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)

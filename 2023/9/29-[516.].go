package leetcode

//给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
//
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：s = "bbbab"
//输出：4
//解释：一个可能的最长回文子序列为 "bbbb" 。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出：2
//解释：一个可能的最长回文子序列为 "bb" 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 1116 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindromeSubseq(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, length)
			}
			if i == j {
				dp[i][j] = 1
			}
		}
	}

	//dp[i][j] i~j 最长回文字符串
	// s[i]==s[i]  dp[i][j]=dp[i+1][j-1]+2
	// !=    dp[i][j]=max(dp[i+1][j],dp[i][j-1])

	for i := length - 1; i >= 0; i-- { //*
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max0(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][length-1]
}

func max0(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

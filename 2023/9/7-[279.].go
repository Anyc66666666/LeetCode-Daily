package leetcode

import "math"

//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
//
//
//
// 示例 1：
//
//
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//
// 示例 2：
//
//
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
//
//
// Related Topics 广度优先搜索 数学 动态规划 👍 1807 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	//x:=math.Pow(float64(n),0.5)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min2(dp[j-i*i]+1, dp[j])
		}
	}

	return dp[n]

}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

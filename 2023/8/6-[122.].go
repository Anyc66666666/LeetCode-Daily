package leetcode

//给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。
//
//
//
// 示例 1：
//
//
//输入：prices = [7,1,5,3,6,4]
//输出：7
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
//     总利润为 4 + 3 = 7 。
//
// 示例 2：
//
//
//输入：prices = [1,2,3,4,5]
//输出：4
//解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
//     总利润为 4 。
//
// 示例 3：
//
//
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 3 * 10⁴
// 0 <= prices[i] <= 10⁴
//
//
// Related Topics 贪心 数组 动态规划 👍 2195 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	//var ans int
	//for i:=1;i<len(prices);i++{
	//	ans+=max1(0,prices[i]-prices[i-1])
	//}
	//return ans

	//-------
	//n := len(prices)
	//dp := make([][2]int, n)
	//dp[0][1] = -prices[0]
	//for i := 1; i < n; i++ {
	//	dp[i][0] = max1(dp[i-1][0], dp[i-1][1]+prices[i])
	//	dp[i][1] = max1(dp[i-1][1], dp[i-1][0]-prices[i])
	//}
	//return dp[n-1][0]

	//---------
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0 = max1(dp0, dp1+prices[i])
		dp1 = max1(dp1, dp0-prices[i])
	}
	return dp0

	//作者：力扣官方题解
	//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/solutions/476791/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
	//来源：力扣（LeetCode）
	//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

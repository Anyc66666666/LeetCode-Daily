package leetcode

//给定一个整数数组
// prices，其中第 prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//
//
//输入: prices = [1]
//输出: 0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1586 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit5(prices []int) int {
	//我们目前持有一支股票，对应的「累计最大收益」记为 f[i][0]
	//
	//我们目前不持有任何股票，并且处于冷冻期中，对应的「累计最大收益」记为 f[i][1]
	//
	//我们目前不持有任何股票，并且不处于冷冻期中，对应的「累计最大收益」记为 f[i][2】
	//
	//作者：力扣官方题解
	//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/solutions/323509/zui-jia-mai-mai-gu-piao-shi-ji-han-leng-dong-qi-4/
	//来源：力扣（LeetCode）
	//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益

	f := make([][3]int, n)

	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max4(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max4(f[i-1][1], f[i-1][2])
	}
	return max4(f[n-1][1], f[n-1][2])

}
func max4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

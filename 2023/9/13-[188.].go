package leetcode

import "math"

//给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1：
//
//
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//
// 示例 2：
//
//
//输入：k = 2, prices = [3,2,6,5,0,3]
//输出：7
//解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3
//。
//
//
//
// 提示：
//
//
// 1 <= k <= 100
// 1 <= prices.length <= 1000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1017 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = min(k, n/2)
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64 / 2
		sell[i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[0] = max1(buy[0], sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = max1(buy[j], sell[j]-prices[i])
			sell[j] = max1(sell[j], buy[j-1]+prices[i])
		}
	}
	return max1(sell...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

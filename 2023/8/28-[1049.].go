package leetcode

//有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
//
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//
//
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
//
//
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。
//
//
//
// 示例 1：
//
//
//输入：stones = [2,7,4,1,8,1]
//输出：1
//解释：
//组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
//组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
//组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
//组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
//
//
// 示例 2：
//
//
//输入：stones = [31,26,33,21,40]
//输出：5
//
//
//
//
// 提示：
//
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 738 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func lastStoneWeightII(stones []int) int {
	//sum := 0
	//for _, v := range stones {
	//	sum += v
	//}
	//n, m := len(stones), sum/2
	//dp := make([][]bool, n+1)
	//for i := range dp {
	//	dp[i] = make([]bool, m+1)
	//}
	//dp[0][0] = true
	//for i, weight := range stones {
	//	for j := 0; j <= m; j++ {
	//		if j < weight {
	//			dp[i+1][j] = dp[i][j]
	//		} else {
	//			dp[i+1][j] = dp[i][j] || dp[i][j-weight]
	//		}
	//	}
	//}
	//for j := m; ; j-- {
	//	if dp[n][j] {
	//		return sum - 2*j
	//	}
	//}
	//
	////作者：力扣官方题解
	////链接：https://leetcode.cn/problems/last-stone-weight-ii/solutions/817930/zui-hou-yi-kuai-shi-tou-de-zhong-liang-i-95p9/
	////来源：力扣（LeetCode）
	////著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	//-------------

	var sum int
	for _, v := range stones {
		sum += v
	}
	m := sum / 2
	dp := make([]int, m+1)
	for _, stone := range stones {
		for j := m; j >= stone; j-- {
			dp[j] = max3(dp[j], dp[j-stone]+stone)
		}
	}

	return sum - 2*dp[m]
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

/*
// 0.1 如何计算：
// 我们可以先求出这一堆石头的总质量sum，
// 而sum = heap1 + heap2    （heap1 > heap2）
// heap1 - heap2 = sum - 2 * heap2
// 要求heap1 - heap2 的最小值，就可以转化成求sum - 2 * heap2 的最小值，
// 也就转化成了求 2 * heap2 的最大值，也就是求heap2的最大值（前提：sum - 2 * heap2 >= 0 等价于 heap2 <= sum / 2）
// 那么就转化成了01背包问题：背包的最大容量为 sum / 2


// 1.状态表示：f[i][j]
// 1.1 集合划分：i表示前i个石头，j表示背包最大容量；f[i][j]就表示前i个数，容量为j的背包能装下的石头最大质量
// （注意：此题的石头体积与石头质量都是stones[i]）
// 1.2 属性（一般为最大值、最小值、数量）：最大值


// 2.状态计算：
// 2.1 f[i][j]可以表示成f[i - 1][j]，表示不加上当前第i块石头的背包最大质量
// 2.2 f[i][j]还可以表示成f[i - 1][j - stones[i]] + stones[i]，表示加上第i块石头后背包最大质量
// 2.3 所以最后总和为j的所有方案为：f[i][j] = max(f[i - 1][j] + f[i - 1][j - stones[i]] + stones[i])


// 3.优化：
// 通过滚动数组可以将二维数组优化成一维数组：f[i][j] --> f[j]

*/

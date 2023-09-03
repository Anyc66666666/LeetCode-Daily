func change(amount int, coins []int) int {
dp := make([]int, amount+1)
dp[0] = 1
for _, coin := range coins {
for i := coin; i <= amount; i++ {
dp[i] += dp[i-coin]
}
}
return dp[amount]
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/coin-change-ii/solutions/821278/ling-qian-dui-huan-ii-by-leetcode-soluti-f7uh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
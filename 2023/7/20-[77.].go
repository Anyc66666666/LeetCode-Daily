package leetcode

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//
// 示例 2：
//
//
//输入：n = 1, k = 1
//输出：[[1]]
//
//
//
// 提示：
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// Related Topics 回溯 👍 1429 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	// 1 2 3 4 // 12 13 14 23 24 34
	var ans [][]int
	a := make([]int, 0, k)
	var dfs func(now int)
	dfs = func(now int) {
		if len(a) == k {
			temp := make([]int, k)
			copy(temp, a)
			ans = append(ans, temp)
			return
		}

		for i := now; i <= n; i++ {
			//比如填满k还需要4个，但是i到n到距离不足4个，此时可以直接return
			if n-i+1 < k-len(a) {
				return
			}

			a = append(a, i)
			dfs(i + 1)
			a = a[:len(a)-1]

		}

	}

	dfs(1)

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

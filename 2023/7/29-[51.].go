package leetcode

import "strings"

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 1838 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	res := [][]string{}
	helper(0, bd, &res, n, cols, diag1, diag2)
	return res

}

func helper(r int, bd [][]string, res *[][]string, n int, cols, diag1, diag2 map[int]bool) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if !cols[c] && !diag1[r+c] && !diag2[r-c] {
			bd[r][c] = "Q"
			cols[c] = true
			diag1[r+c] = true
			diag2[r-c] = true
			helper(r+1, bd, res, n, cols, diag1, diag2)
			bd[r][c] = "."
			cols[c] = false
			diag1[r+c] = false
			diag2[r-c] = false
		}
	}
}

//作者：xiao_ben_zhu
//链接：https://leetcode.cn/problems/n-queens/solution/shou-hua-tu-jie-cong-jing-dian-de-nhuang-hou-wen-t/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
//leetcode submit region end(Prohibit modification and deletion)

package leetcode

//给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
//
// 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
//
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[1,4,7],[2,5,8],[3,6,9]]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3],[4,5,6]]
//输出：[[1,4],[2,5],[3,6]]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 1000
// 1 <= m * n <= 10⁵
// -10⁹ <= matrix[i][j] <= 10⁹
//
//
// Related Topics 数组 矩阵 模拟 👍 243 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func transpose(matrix [][]int) [][]int {
	var a [][]int
	m := make(map[int][]int)
	for _, r := range matrix {
		for k, c := range r {
			m[k] = append(m[k], c)
		}
	}
	for i := 0; i < len(m); i++ {
		a = append(a, m[i])
	}
	return a

}

//leetcode submit region end(Prohibit modification and deletion)

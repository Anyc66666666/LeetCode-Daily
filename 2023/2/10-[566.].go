package leetcode

//在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个 m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。
//
//
// 给你一个由二维数组 mat 表示的 m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。
//
// 重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。
//
// 如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
//
//
//
// 示例 1：
//
//
//输入：mat = [[1,2],[3,4]], r = 1, c = 4
//输出：[[1,2,3,4]]
//
//
// 示例 2：
//
//
//输入：mat = [[1,2],[3,4]], r = 2, c = 4
//输出：[[1,2],[3,4]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 100
// -1000 <= mat[i][j] <= 1000
// 1 <= r, c <= 300
//
//
// Related Topics 数组 矩阵 模拟 👍 377 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	var a []int
	for _, rol := range mat {
		for _, col := range rol {
			a = append(a, col)
		}
	}
	if len(a) != r*c {
		return mat
	}

	var mat1 [][]int
	var b int
	for i := 0; i < r; i++ {
		m := make([]int, c, c)
		for j := 0; j < c; j++ {
			m[j] = a[b]
			b++
		}
		mat1 = append(mat1, m)
	}

	return mat1

}

//leetcode submit region end(Prohibit modification and deletion)
/*
func matrixReshape(nums [][]int, r int, c int) [][]int {
    n, m := len(nums), len(nums[0])
    if n*m != r*c {
        return nums
    }
    ans := make([][]int, r)
    for i := range ans {
        ans[i] = make([]int, c)
    }
    for i := 0; i < n*m; i++ {
        ans[i/c][i%c] = nums[i/m][i%m]
    }
    return ans
}

*/

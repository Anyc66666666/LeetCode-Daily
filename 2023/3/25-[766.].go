package leetcode

//给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。
//
// 如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵 。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
//输出：true
//解释：
//在上述矩阵中, 其对角线为:
//"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]"。
//各条对角线上的所有元素均相同, 因此答案是 True 。
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2],[2,2]]
//输出：false
//解释：
//对角线 "[1, 2]" 上的元素不同。
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 20
// 0 <= matrix[i][j] <= 99
//
//
//
//
// 进阶：
//
//
// 如果矩阵存储在磁盘上，并且内存有限，以至于一次最多只能将矩阵的一行加载到内存中，该怎么办？
// 如果矩阵太大，以至于一次只能将不完整的一行加载到内存中，该怎么办？
//
//
// Related Topics 数组 矩阵 👍 288 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isToeplitzMatrix(matrix [][]int) bool {
	type Loc struct {
		Val int
		X   int
		Y   int
	}
	var left, right []Loc
	y := len(matrix)
	x := len(matrix[0])
	for k, row := range matrix {
		if k == 0 {
			for xx, col := range row {
				right = append(right, Loc{
					Val: col,
					X:   xx,
					Y:   0,
				})
			}
			continue
		}
		left = append(left, Loc{
			Val: row[0],
			X:   0,
			Y:   k,
		})
	}

	for _, v := range left {
		for x1, y1 := v.X, v.Y; x1 < x; x1++ {
			if y1 < y {
				//fmt.Println(matrix[y1][x1],v.Val)
				if matrix[y1][x1] == v.Val {
					y1++
					continue
				} else {
					return false
				}
			}
			y1++

		}

	}

	for _, v := range right {
		for x1, y1 := v.X, v.Y; x1 < x; x1++ {
			if y1 < y {
				//fmt.Println(matrix[y1][x1],v.Val)
				if matrix[y1][x1] == v.Val {
					y1++
					continue
				} else {
					return false
				}
			}
			y1++

		}

	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)
/*

根据定义，当且仅当矩阵中每个元素都与其左上角相邻的元素（如果存在）相等时，
该矩阵为托普利茨矩阵。因此，我们遍历该矩阵，将每一个元素和它左上角的元素相比对即可。

func isToeplitzMatrix(matrix [][]int) bool {
    n, m := len(matrix), len(matrix[0])
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] != matrix[i-1][j-1] {
                return false
            }
        }
    }
    return true
}



作者：力扣官方题解
链接：https://leetcode.cn/problems/toeplitz-matrix/solutions/613732/tuo-pu-li-ci-ju-zhen-by-leetcode-solutio-57bb/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

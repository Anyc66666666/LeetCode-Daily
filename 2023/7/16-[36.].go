package leetcode

//请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
//
//
// 注意：
//
//
// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 空白格用 '.' 表示。
//
//
//
//
// 示例 1：
//
//
//输入：board =
//[["5","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//输出：true
//
//
// 示例 2：
//
//
//输入：board =
//[["8","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//输出：false
//解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无
//效的。
//
//
//
// 提示：
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字（1-9）或者 '.'
//
//
// Related Topics 数组 哈希表 矩阵 👍 1118 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	//for _,r:=range board{
	//	m:=make(map[byte]struct{})
	//	for _,v:=range r{
	//		if v=='.'{
	//			continue
	//		}
	//		if _,ok:=m[v];ok{
	//			return false
	//		}else {
	//			m[v]= struct{}{}
	//		}
	//	}
	//}
	//
	//for c:=0;c<len(board[0]);c++{
	//	m:=make(map[byte]struct{})
	//	for _,v:=range board{
	//		if v[c]=='.'{
	//			continue
	//		}
	//		if _,ok:=m[v[c]];ok{
	//			return false
	//		}else {
	//			m[v[c]]= struct{}{}
	//		}
	//	}
	//}
	//
	////0-3 3-6 6-9
	////0-3 3-6 6-9
	//
	//
	//
	//for x:=0;x<9;x+=3{
	//	for y:=0;y<9;y+=3{
	//		m:=make(map[byte]struct{})
	//		for i:=x;i<x+3;i++{
	//			for j:=y;j<y+3;j++{
	//
	//				xx:=board[i][j]
	//				if xx=='.'{
	//					continue
	//				}
	//				if _,ok:=m[xx];ok{
	//					return false
	//				}else {
	//					m[xx]= struct{}{}
	//				}
	//
	//			}
	//		}
	//
	//	}
	//}
	//
	//return true

	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true

	//作者：LeetCode-Solution
	//链接：https://leetcode.cn/problems/valid-sudoku/solution/you-xiao-de-shu-du-by-leetcode-solution-50m6/
	//来源：力扣（LeetCode）
	//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

//leetcode submit region end(Prohibit modification and deletion)

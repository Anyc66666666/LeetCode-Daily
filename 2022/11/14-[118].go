package leetcode

//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
//
// 示例 2:
//
//
//输入: numRows = 1
//输出: [[1]]
//
//
//
//
// 提示:
//
//
// 1 <= numRows <= 30
//
//
// Related Topics 数组 动态规划 👍 868 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	o := []int{1}
	t := []int{1, 1}
	two := make([][]int, numRows)
	if numRows == 1 {
		two[0] = o
		return two
	}
	if numRows == 2 {
		two[0] = o
		two[1] = t
		return two
	}
	two[0] = o
	two[1] = t

	for i := 2; i < numRows; i++ {
		s := o
		for j := 0; j < len(two[i-1])-1; j++ {
			m := two[i-1][j] + two[i-1][j+1]
			s = append(s, m)
		}
		s = append(s, 1)
		two[i] = s
	}
	return two

}

//leetcode submit region end(Prohibit modification and deletion)

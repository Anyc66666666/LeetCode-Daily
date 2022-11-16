package leetcode

//给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
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
//输入: rowIndex = 3
//输出: [1,3,3,1]
//
//
// 示例 2:
//
//
//输入: rowIndex = 0
//输出: [1]
//
//
// 示例 3:
//
//
//输入: rowIndex = 1
//输出: [1,1]
//
//
//
//
// 提示:
//
//
// 0 <= rowIndex <= 33
//
//
//
//
// 进阶：
//
// 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
//
// Related Topics 数组 动态规划 👍 439 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	o := []int{1}
	t := []int{1, 1}
	if rowIndex == 0 {
		return o
	}
	if rowIndex == 1 {
		return t
	}
	tt := getRow(rowIndex - 1)
	m := o

	for k := 0; k < len(tt)-1; k++ {
		tmp := tt[k] + tt[k+1]
		m = append(m, tmp)

	}
	m = append(m, 1)
	return m
}

//leetcode submit region end(Prohibit modification and deletion)

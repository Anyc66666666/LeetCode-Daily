package leetcode

//给你一个 n * n 的网格 grid ，上面放置着一些 1 x 1 x 1 的正方体。每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单
//元格 (i, j) 上。
//
// 放置好正方体后，任何直接相邻的正方体都会互相粘在一起，形成一些不规则的三维形体。
//
// 请你返回最终这些形体的总表面积。
//
// 注意：每个形体的底面也需要计入表面积中。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,2],[3,4]]
//输出：34
//
//
// 示例 2：
//
//
//输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
//输出：32
//
//
// 示例 3：
//
//
//输入：grid = [[2,2,2],[2,1,2],[2,2,2]]
//输出：46
//
//
//
//
// 提示：
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 50
// 0 <= grid[i][j] <= 50
//
//
// Related Topics 几何 数组 数学 矩阵 👍 171 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func surfaceArea(grid [][]int) int {
	area := 0
	for i, row := range grid {
		for j, level := range row {
			if level > 0 {
				area += level*4 + 2
				if i > 0 {
					area -= findMin(grid[i-1][j], level) * 2
				}
				if j > 0 {
					area -= findMin(grid[i][j-1], level) * 2
				}
			}
		}
	}
	return area

}
func findMin(i, j int) int {
	if i > j {
		return j
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

/*
先算每个点不计算相邻点的侧面重合的表面积之和，然后减去重叠的部分，注意这里重叠的部分要乘以二，因为只计算和上一排的同一个位置和当前方块的左边方块的重合，

*/

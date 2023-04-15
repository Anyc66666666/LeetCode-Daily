package leetcode

//在
// n x n 的网格
// grid 中，我们放置了一些与 x，y，z 三轴对齐的
// 1 x 1 x 1 立方体。
//
// 每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。
//
// 现在，我们查看这些立方体在 xy 、yz 和 zx 平面上的投影。
//
// 投影 就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。
//
// 返回 所有三个投影的总面积 。
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：[[1,2],[3,4]]
//输出：17
//解释：这里有该形体在三个轴对齐平面上的三个投影(“阴影部分”)。
//
//
// 示例 2:
//
//
//输入：grid = [[2]]
//输出：5
//
//
// 示例 3：
//
//
//输入：[[1,0],[0,2]]
//输出：8
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 1 <= n <= 50
// 0 <= grid[i][j] <= 50
//
//
// Related Topics 几何 数组 数学 矩阵 👍 147 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func projectionArea(grid [][]int) int {
	//y:=len(grid)
	//x:=len(grid[0])
	//s1:=x*y
	//var s2,s3 int
	//max82:=make([]int,x,x)
	//
	//for _,rol:=range grid{
	//	var max81 int
	//	for k,col:=range rol{
	//		max81=max88(max81,col)
	//		max82[k]=max88(max82[k],col)
	//
	//	}
	//
	//	s2=s2+max81
	//}
	//for _,v:=range max82{
	//	s3+=v
	//}
	//return s1+s2+s3
	var xyArea, yzArea, zxArea int
	for i, row := range grid {
		yzHeight, zxHeight := 0, 0
		for j, v := range row {
			if v > 0 {
				xyArea++
			}
			yzHeight = max8(yzHeight, grid[j][i])
			zxHeight = max8(zxHeight, v)
		}
		yzArea += yzHeight
		zxArea += zxHeight
	}
	return xyArea + yzArea + zxArea

}
func max8(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

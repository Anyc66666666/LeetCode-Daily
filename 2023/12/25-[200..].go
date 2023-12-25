package leetcode

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2386 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(y, x int)
	color := byte('1')
	dfs = func(y, x int) {
		if y >= m || y < 0 ||
			x >= n || x < 0 ||
			grid[y][x] != '1' {
			return
		}
		grid[y][x] = color
		dfs(y, x+1)
		dfs(y, x-1)
		dfs(y+1, x)
		dfs(y-1, x)
	}
	for y := 0; y < m; y++ {
		//3)全扫描
		for x := 0; x < n; x++ {
			if grid[y][x] != '1' {
				continue
			}
			color++ // 先加后操作
			dfs(y, x)
		}
	}
	return int(color - '1') //取巧规避第一个颜色
}

//leetcode submit region end(Prohibit modification and deletion)

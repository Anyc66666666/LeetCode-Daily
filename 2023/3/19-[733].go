package leetcode

import "fmt"

//有一幅以 m x n 的二维整数数组表示的图画 image ，其中 image[i][j] 表示该图画的像素值大小。
//
// 你也被给予三个整数 sr , sc 和 newColor 。你应该从像素 image[sr][sc] 开始对图像进行 上色填充 。
//
// 为了完成 上色工作 ，从初始像素开始，记录初始坐标的 上下左右四个方向上 像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对
//应 四个方向上 像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为 newColor 。
//
// 最后返回 经过上色渲染后的图像 。
//
//
//
// 示例 1:
//
//
//
//
//输入: image = [[1,1,1],[1,1,0],[1,0,1]]，sr = 1, sc = 1, newColor = 2
//输出: [[2,2,2],[2,2,0],[2,0,1]]
//解析: 在图像的正中间，(坐标(sr,sc)=(1,1)),在路径上所有符合条件的像素点的颜色都被更改成2。
//注意，右下角的像素没有更改为2，因为它不是在上下左右四个方向上与初始点相连的像素点。
//
//
// 示例 2:
//
//
//输入: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
//输出: [[2,2,2],[2,2,2]]
//
//
//
//
// 提示:
//
//
// m == image.length
// n == image[i].length
// 1 <= m, n <= 50
// 0 <= image[i][j], newColor < 2¹⁶
// 0 <= sr < m
// 0 <= sc < n
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 👍 433 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	old := image[sr][sc]
	fmt.Println(old)
	y := len(image)
	x := len(image[0])
	image[sr][sc] = color

	m := make(map[string]struct{})
	//x,y  x-1,y  x+1,y x,y+1  x,y-1

	var help func(a, b int)
	help = func(a, b int) {
		if a-1 >= 0 && b < x {
			if image[a-1][b] == old {
				image[a-1][b] = color
				_, ok := m[fmt.Sprintf("%v%v", a-1, b)]
				if ok {
					return
				}

				m[fmt.Sprintf("%v%v", a-1, b)] = struct{}{}
				help(a-1, b)

			}
		}

		if a+1 < y && b < x {
			if image[a+1][b] == old {
				image[a+1][b] = color
				_, ok := m[fmt.Sprintf("%v%v", a+1, b)]
				if ok {
					return
				}

				m[fmt.Sprintf("%v%v", a+1, b)] = struct{}{}
				help(a+1, b)
			}
		}

		if b-1 >= 0 && a < y {
			if image[a][b-1] == old {
				image[a][b-1] = color
				_, ok := m[fmt.Sprintf("%v%v", a, b-1)]
				if ok {
					return
				}

				m[fmt.Sprintf("%v%v", a, b-1)] = struct{}{}
				help(a, b-1)
			}
		}

		if b+1 < x && a < y {
			if image[a][b+1] == old {
				image[a][b+1] = color
				_, ok := m[fmt.Sprintf("%v%v", a, b+1)]
				if ok {
					return
				}

				m[fmt.Sprintf("%v%v", a, b+1)] = struct{}{}
				help(a, b+1)
			}
		}

	}

	help(sr, sc)
	return image

}

//leetcode submit region end(Prohibit modification and deletion)

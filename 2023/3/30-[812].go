package leetcode

import "math"

//给你一个由 X-Y 平面上的点组成的数组 points ，其中 points[i] = [xi, yi] 。从其中取任意三个不同的点组成三角形，返回能组成的
//最大三角形的面积。与真实值误差在 10⁻⁵ 内的答案将会视为正确答案。
//
//
//
// 示例 1：
//
//
//输入：points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
//输出：2.00000
//解释：输入中的 5 个点如上图所示，红色的三角形面积最大。
//
//
// 示例 2：
//
//
//输入：points = [[1,0],[0,0],[0,1]]
//输出：0.50000
//
//
//
//
// 提示：
//
//
// 3 <= points.length <= 50
// -50 <= xi, yi <= 50
// 给出的所有点 互不相同
//
//
// Related Topics 几何 数组 数学 👍 182 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func largestTriangleArea(points [][]int) float64 {
	square2 := 0.0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				a := getSide(points[i][0], points[i][1], points[j][0], points[j][1])
				b := getSide(points[i][0], points[i][1], points[k][0], points[k][1])
				c := getSide(points[k][0], points[k][1], points[j][0], points[j][1])
				p := (a + b + c) / 2
				square := p * (p - a) * (p - b) * (p - c)
				if square > square2 {
					square2 = square
				}

			}
		}
	}
	return math.Sqrt(square2)
}
func getSide(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

//leetcode submit region end(Prohibit modification and deletion)
/*
我们可以枚举所有的三角形，然后计算三角形的面积并找出最大的三角形面积。设三角形三个
顶点的坐标为(x1,y1)、(x2,y2) 和(x;,y3), 则三角形面积S可以用行列式的绝对值表示:
          |x1,y1,1|
S= 1/2 *  |x2,y2,2|  = 1/2 * |x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2|
          |x3,y3,3|




func triangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
    return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
}

func largestTriangleArea(points [][]int) (ans float64) {
    for i, p := range points {
        for j, q := range points[:i] {
            for _, r := range points[:j] {
                ans = math.Max(ans, triangleArea(p[0], p[1], q[0], q[1], r[0], r[1]))
            }
        }
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/largest-triangle-area/solutions/1490629/zui-da-san-jiao-xing-mian-ji-by-leetcode-yefh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

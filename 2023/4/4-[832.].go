package leetcode

//给定一个
// n x n 的二进制矩阵 image ，先 水平 翻转图像，然后 反转 图像并返回 结果 。
//
// 水平翻转图片就是将图片的每一行都进行翻转，即逆序。
//
//
// 例如，水平翻转 [1,1,0] 的结果是 [0,1,1]。
//
//
// 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。
//
//
// 例如，反转 [0,1,1] 的结果是 [1,0,0]。
//
//
//
//
// 示例 1：
//
//
//输入：image = [[1,1,0],[1,0,1],[0,0,0]]
//输出：[[1,0,0],[0,1,0],[1,1,1]]
//解释：首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
//     然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
//
//
// 示例 2：
//
//
//输入：image = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
//输出：[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
//解释：首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
//     然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
//
//
//
//
// 提示：
//
//
//
//
//
// n == image.length
// n == image[i].length
// 1 <= n <= 20
// images[i][j] == 0 或 1.
//
//
// Related Topics 数组 双指针 矩阵 模拟 👍 298 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func flipAndInvertImage(image [][]int) [][]int {
	aa := make([][]int, 0, len(image))
	for _, v := range image {
		a := make([]int, 0, len(v))
		for i := len(v) - 1; i >= 0; i-- {
			if v[i] == 1 {
				a = append(a, 0)
				continue
			}
			a = append(a, 1)

		}
		aa = append(aa, a)
	}
	return aa

}

//leetcode submit region end(Prohibit modification and deletion)

/*

//两边向中间取  位运算

func flipAndInvertImage(image [][]int) [][]int {
    for _, row := range image {
        left, right := 0, len(row)-1
        for left < right {
            if row[left] == row[right] {
                row[left] ^= 1
                row[right] ^= 1
            }
            left++
            right--
        }
        if left == right {
            row[left] ^= 1
        }
    }
    return image
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/flipping-an-image/solutions/617961/fan-zhuan-tu-xiang-by-leetcode-solution-yljd/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

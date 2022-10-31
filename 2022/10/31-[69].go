package leetcode

import "fmt"

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
//
//
// 示例 1：
//
//
//输入：x = 4
//输出：2
//
//
// 示例 2：
//
//
//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
//
//
//
//
// 提示：
//
//
// 0 <= x <= 2³¹ - 1
//
//
// Related Topics 数学 二分查找 👍 1174 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	var left, right float64

	left, right = float64(1), float64(x)

	var mid float64
	for {
		if left*left == float64(x) {
			return int(left)
		}
		if right*right == float64(x) {
			return int(right)
		}

		mid = (right-left)/2 + left

		if mid*mid > float64(x) {
			if right-left < 0.000001 {
				fmt.Println(left, mid, right)
				return int(mid)
			}
			right = mid
		} else if mid*mid < float64(x) {
			if right-left < 0.000001 {
				fmt.Println(left, mid, right)
				return int(right)
			}
			left = mid
		} else {
			return int(mid)
		}

	}

}

//leetcode submit region end(Prohibit modification and deletion)

package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−2³¹, 231 − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
//
//
// 示例 1：
//
//
//输入：x = 123
//输出：321
//
//
// 示例 2：
//
//
//输入：x = -123
//输出：-321
//
//
// 示例 3：
//
//
//输入：x = 120
//输出：21
//
//
// 示例 4：
//
//
//输入：x = 0
//输出：0
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
// Related Topics 数学 👍 3827 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var ok bool
	if x < 0 {
		x = -x
		ok = true
	}
	str := fmt.Sprint(x)
	var a string
	for i := 0; i < len(str); i++ {
		a = string(str[i]) + a
	}

	m, _ := strconv.Atoi(a)
	f := math.Pow(2, 31)
	if ok {

		if float64(m) > f {
			return 0
		}
		m = -m
	}

	if float64(m) > f {
		return 0
	}

	return m

}

//leetcode submit region end(Prohibit modification and deletion)

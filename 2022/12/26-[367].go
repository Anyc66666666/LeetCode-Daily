package leetcode

//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
// 进阶：不要 使用任何内置的库函数，如 sqrt 。
//
//
//
// 示例 1：
//
//
//输入：num = 16
//输出：true
//
//
// 示例 2：
//
//
//输入：num = 14
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= num <= 2^31 - 1
//
//
// Related Topics 数学 二分查找 👍 460 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left := 0
	right := num
	var mid int
	for right-left > 0 {
		mid = (right - left) / 2
		if mid == 0 {
			return false
		}
		if (left+mid)*(left+mid) == num {
			return true
		} else if (left+mid)*(left+mid) > num {
			right = right - mid
		} else {
			left = left + mid
		}

	}
	return false

}

//leetcode submit region end(Prohibit modification and deletion)

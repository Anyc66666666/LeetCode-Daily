package leetcode

import "fmt"

//给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
//
//
//
// 示例 1:
//
//
//输入: num = 100
//输出: "202"
//
//
// 示例 2:
//
//
//输入: num = -7
//输出: "-10"
//
//
//
//
// 提示：
//
//
// -10⁷ <= num <= 10⁷
//
//
// Related Topics 数学 👍 197 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	var a int
	var s string

	if num == 0 {
		return "0"
	}
	var ok bool
	if num < 0 {
		num = -num
		ok = true
	}
	fmt.Println(a)
	for num != 0 {
		a = num % 7
		num = num / 7
		s = fmt.Sprint(a) + s
	}
	if ok {
		return "-" + s
	}
	return s

}

//leetcode submit region end(Prohibit modification and deletion)

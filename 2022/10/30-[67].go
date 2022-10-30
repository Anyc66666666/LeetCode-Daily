package leetcode

//给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
//
//
//
// 示例 1：
//
//
//输入:a = "11", b = "1"
//输出："100"
//
// 示例 2：
//
//
//输入：a = "1010", b = "1011"
//输出："10101"
//
//
//
// 提示：
//
//
// 1 <= a.length, b.length <= 10⁴
// a 和 b 仅由字符 '0' 或 '1' 组成
// 字符串如果不是 "0" ，就不含前导零
//
//
// Related Topics 位运算 数学 字符串 模拟 👍 906 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	var m = map[uint8]int{
		'1': 1,
		'0': 0,
	}
	var n = map[int]string{
		1: "1",
		0: "0",
	}
	var x, y int
	if len(a) < len(b) {
		x, y = len(b), len(a)
	} else {
		x, y = len(a), len(b)
	}

	for j := 0; j < x-y; j++ {
		if len(a) < len(b) {
			a = "0" + a
		}
		if len(b) < len(a) {
			b = "0" + b
		}
	}
	var s string
	var flag int
	for i := x - 1; i >= 0; i-- {

		if m[a[i]]+m[b[i]]+flag == 2 {
			flag = 1
			s = "0" + s
			if i == 0 {
				s = "1" + s
				return s
			}

		} else if m[a[i]]+m[b[i]]+flag == 3 {
			flag = 1
			s = "1" + s
			if i == 0 {
				s = "1" + s
				return s
			}
		} else {
			s = n[m[a[i]]+m[b[i]]+flag] + s
			flag = 0
		}

	}
	return s

}

//leetcode submit region end(Prohibit modification and deletion)

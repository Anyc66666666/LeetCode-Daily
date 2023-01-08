package leetcode

import "fmt"

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
//
//
//
// 示例 1：
//
//
//输入：num1 = "11", num2 = "123"
//输出："134"
//
//
// 示例 2：
//
//
//输入：num1 = "456", num2 = "77"
//输出："533"
//
//
// 示例 3：
//
//
//输入：num1 = "0", num2 = "0"
//输出："0"
//
//
//
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 10⁴
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
//
//
// Related Topics 数学 字符串 模拟 👍 657 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	m := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	var a int
	var s string
	var b bool

	if len(num1) > len(num2) {
		a = len(num1)
		l := len(num1) - len(num2)

		for j := 0; j < l; j++ {
			num2 = "0" + num2
			fmt.Println(num2)

		}
	} else {
		a = len(num2)
		l := len(num2) - len(num1)
		for i := 0; i < l; i++ {
			num1 = "0" + num1
			fmt.Println("----")
		}
	}
	fmt.Println(num1, num2)

	for i := a - 1; i >= 0; i-- {
		n := m[num1[i]] + m[num2[i]]
		if b {
			n++
			b = false
		}
		if n > 9 {
			n = n % 10
			b = true
		}
		s = fmt.Sprint(n) + s

	}
	if b {
		s = "1" + s
	}

	return s

}

//leetcode submit region end(Prohibit modification and deletion)

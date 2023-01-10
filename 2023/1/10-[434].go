package leetcode

import "strings"

//统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
//
// 请注意，你可以假定字符串里不包括任何不可打印的字符。
//
// 示例:
//
// 输入: "Hello, my name is John"
//输出: 5
//解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
//
//
// Related Topics 字符串 👍 194 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func countSegments(s string) int {
	var a int
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	for i := 0; i < len(s); i++ {
		if i > 0 {

			if s[i] == ' ' && s[i-1] != ' ' {
				a++
			}

		}
	}
	return a + 1

}

//leetcode submit region end(Prohibit modification and deletion)

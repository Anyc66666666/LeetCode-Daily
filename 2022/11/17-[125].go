package leetcode

import (
	"fmt"
	"strings"
)

//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
//
//
// 示例 2：
//
//
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
//
//
// 示例 3：
//
//
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// s 仅由可打印的 ASCII 字符组成
//
//
// Related Topics 双指针 字符串 👍 593 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var ss string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			ss = ss + string(s[i])
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			ss = ss + string(s[i])
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			ss = ss + string(s[i])
		}
	}
	fmt.Println(ss)
	//12321   123321
	for i := 0; i < len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

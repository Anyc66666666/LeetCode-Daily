package leetcode

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 3565 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	ma := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			b = append(b, s[i])
		}
		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(b) == 0 {
				return false
			}
			if ma[b[len(b)-1]] == s[i] {
				b = b[:len(b)-1]
			} else {
				return false
			}
		}

	}
	return len(b) == 0

}

//leetcode submit region end(Prohibit modification and deletion)

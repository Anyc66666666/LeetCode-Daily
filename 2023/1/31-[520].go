package leetcode

import "unicode"

//我们定义，在以下情况时，单词的大写用法是正确的：
//
//
// 全部字母都是大写，比如 "USA" 。
// 单词中所有字母都不是大写，比如 "leetcode" 。
// 如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
//
//
// 给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：word = "USA"
//输出：true
//
//
// 示例 2：
//
//
//输入：word = "FlaG"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= word.length <= 100
// word 由小写和大写英文字母组成
//
//
// Related Topics 字符串 👍 227 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func detectCapitalUse(word string) bool {
	var a int
	for i := 0; i < len(word); i++ {
		if i == 0 && unicode.IsUpper(rune(word[i])) {
			a = 1
			continue
		}
		if i == 0 && !unicode.IsUpper(rune(word[i])) {
			a = 2 //leetcode
			continue
		}
		if a == 1 {
			if i == 1 && !unicode.IsUpper(rune(word[i])) {
				a = 3 //Google
				continue
			}
			if i == 1 && unicode.IsUpper(rune(word[i])) {
				a = 4 //USA
				continue
			}
		}

		if a == 3 && unicode.IsUpper(rune(word[i])) {
			return false
		}
		if a == 4 && !unicode.IsUpper(rune(word[i])) {
			return false
		}
		if a == 2 && unicode.IsUpper(rune(word[i])) {
			return false
		}

	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

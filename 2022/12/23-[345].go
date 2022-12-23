package leetcode

//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
//
//
//
// 示例 1：
//
//
//输入：s = "hello"
//输出："holle"
//
//
// 示例 2：
//
//
//输入：s = "leetcode"
//输出："leotcede"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由 可打印的 ASCII 字符组成
//
//
// Related Topics 双指针 字符串 👍 275 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	m := []byte(s)
	p := make(map[byte]struct{})
	p = map[byte]struct{}{
		'a': {},
		'e': {},
		'o': {},
		'i': {},
		'u': {},
		'A': {},
		'E': {},
		'O': {},
		'I': {},
		'U': {},
	}
	var a []int
	for i := 0; i < len(s); i++ {
		_, ok := p[m[i]]
		if ok {
			a = append(a, i)
		}

	}
	for i := 0; i < len(a)/2; i++ {
		m[a[i]], m[a[len(a)-i-1]] = m[a[len(a)-i-1]], m[a[i]]
	}
	return string(m)
}

//leetcode submit region end(Prohibit modification and deletion)

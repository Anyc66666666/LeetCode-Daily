package leetcode

import "unicode"

//给你一个字符串 s ，根据下述规则反转字符串：
//
//
// 所有非英文字母保留在原有位置。
// 所有英文字母（小写或大写）位置反转。
//
//
// 返回反转后的 s 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "ab-cd"
//输出："dc-ba"
//
//
//
//
//
// 示例 2：
//
//
//输入：s = "a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
//
//
//
//
//
// 示例 3：
//
//
//输入：s = "Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
//
//
//
//
// 提示
//
//
// 1 <= s.length <= 100
// s 仅由 ASCII 值在范围 [33, 122] 的字符组成
// s 不含 '\"' 或 '\\'
//
//
// Related Topics 双指针 字符串 👍 190 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(s string) string {
	left, right := 0, len(s)-1
	var sum int
	var l, r string

	for left < right {
		if !unicode.IsLetter(rune(s[left])) {
			l = l + string(s[left])
			left++
			sum++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) {
			r = string(s[right]) + r
			right--
			sum++
			continue
		}
		l = l + string(s[right])
		r = string(s[left]) + r
		left++
		right--

	}
	if (len(s)-sum)%2 == 1 {
		return l + string(s[left]) + r
	}
	return l + r

}

//leetcode submit region end(Prohibit modification and deletion)

/*

func reverseOnlyLetters(s string) string {
    ans := []byte(s)
    left, right := 0, len(s)-1
    for {
        for left < right && !unicode.IsLetter(rune(s[left])) { // 判断左边是否扫描到字母
            left++
        }
        for right > left && !unicode.IsLetter(rune(s[right])) { // 判断右边是否扫描到字母
            right--
        }
        if left >= right {
            break
        }
        ans[left], ans[right] = ans[right], ans[left]
        left++
        right--
    }
    return string(ans)
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/reverse-only-letters/solutions/1280988/jin-jin-fan-zhuan-zi-mu-by-leetcode-solu-db20/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

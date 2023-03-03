package leetcode

//给你一个字符串 s，最多 可以从中删除一个字符。
//
// 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：s = "aba"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "abca"
//输出：true
//解释：你可以删除字符 'c' 。
//
//
// 示例 3：
//
//
//输入：s = "abc"
//输出：false
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 由小写英文字母组成
//
//
// Related Topics 贪心 双指针 字符串 👍 568 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)

/*

考虑最朴素的方法：首先判断原串是否是回文串，如果是，就返回 true；
如果不是，则枚举每一个位置作为被删除的位置，再判断剩下的字符串是否是回文串。这种做法的渐进时间复杂度是 O(n^2) 会超出时间限制。

我们换一种想法。首先考虑如果不允许删除字符，如何判断一个字符串是否是回文串。常见的做法是使用双指针。
定义左右指针，初始时分别指向字符串的第一个字符和最后一个字符，每次判断左右指针指向的字符是否相同，如果不相同，则不是回文串；
如果相同，则将左右指针都往中间移动一位，直到左右指针相遇，则字符串是回文串。

在允许最多删除一个字符的情况下，同样可以使用双指针，通过贪心实现。
初始化两个指针 low 和 high 分别指向字符串的第一个字符和最后一个字符。
每次判断两个指针指向的字符是否相同，如果相同，则更新指针，
将 low 加 1，high 减 1，然后判断更新后的指针范围内的子串是否是回文字符串。
如果两个指针指向的字符不同，则两个字符中必须有一个被删除，
此时我们就分成两种情况：即删除左指针对应的字符，留下子串 s[low+1:high] :，或者删除右指针对应的字符，留下子串 s[low:high−1]。
当这两个子串中至少有一个是回文串时，就说明原始字符串删除一个字符之后就以成为回文串。

作者：力扣官方题解
链接：https://leetcode.cn/problems/valid-palindrome-ii/solutions/251842/yan-zheng-hui-wen-zi-fu-chuan-ii-by-leetcode-solut/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


*/

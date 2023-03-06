package leetcode

//给定一个字符串 s，统计并返回具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是成组连续的。
//
// 重复出现（不同位置）的子串也要统计它们出现的次数。
//
// 示例 1：
//
//
//输入：s = "00110011"
//输出：6
//解释：6 个子串满足具有相同数量的连续 1 和 0 ："0011"、"01"、"1100"、"10"、"0011" 和 "01" 。
//注意，一些重复出现的子串（不同位置）要统计它们出现的次数。
//另外，"00110011" 不是有效的子串，因为所有的 0（还有 1 ）没有组合在一起。
//
// 示例 2：
//
//
//输入：s = "10101"
//输出：4
//解释：有 4 个子串："10"、"01"、"10"、"01" ，具有相同数量的连续 1 和 0 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s[i] 为 '0' 或 '1'
//
//
// Related Topics 双指针 字符串 👍 503 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func countBinarySubstrings(s string) int {
	counts := []int{}
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
	}
	ans := 0
	for i := 1; i < len(counts); i++ {
		ans += min123(counts[i], counts[i-1])
	}
	return ans

}

func min123(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

/*
我们可以将字符串 sss 按照 000 和 111 的连续段分组，存在counts 数组中，
例如 s=00111011，可以得到这样的counts 数组：counts={2,3,1,2}。

这里 counts 数组中两个相邻的数一定代表的是两种不同的字符。
假设 counts 数组中两个相邻的数字为 u 或者 v，
它们对应着 u 个 0 和 v 个 1，或者 u 个1 和 v 个 0。
它们能组成的满足条件的子串数目为 min{u,v}，即一对相邻的数字对答案的贡献。

我们只要遍历所有相邻的数对，求它们的贡献总和，即可得到答案。

*/

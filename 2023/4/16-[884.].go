package leetcode

import "strings"

//句子 是一串由空格分隔的单词。每个 单词 仅由小写字母组成。
//
// 如果某个单词在其中一个句子中恰好出现一次，在另一个句子中却 没有出现 ，那么这个单词就是 不常见的 。
//
// 给你两个 句子 s1 和 s2 ，返回所有 不常用单词 的列表。返回列表中单词可以按 任意顺序 组织。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s1 = "this apple is sweet", s2 = "this apple is sour"
//输出：["sweet","sour"]
//
//
// 示例 2：
//
//
//输入：s1 = "apple apple", s2 = "banana"
//输出：["banana"]
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 200
// s1 和 s2 由小写英文字母和空格组成
// s1 和 s2 都不含前导或尾随空格
// s1 和 s2 中的所有单词间均由单个空格分隔
//
//
// Related Topics 哈希表 字符串 👍 179 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func uncommonFromSentences(s1 string, s2 string) []string {
	var str []string
	ss1 := make(map[string]int)
	ss2 := make(map[string]int)
	str1 := strings.Split(s1, " ")
	str2 := strings.Split(s2, " ")
	for _, v := range str1 {
		ss1[v]++
	}
	for _, v := range str2 {
		ss2[v]++
	}
	for k, v := range ss1 {
		if v == 1 {
			if ss2[k] == 0 {
				str = append(str, k)
			}
		}
	}
	for k, v := range ss2 {
		if v == 1 {
			if ss1[k] == 0 {
				str = append(str, k)
			}
		}
	}
	return str

}

//leetcode submit region end(Prohibit modification and deletion)
/*

func uncommonFromSentences(s1 string, s2 string) []string {
	freq := make(map[string]int)

	insert := func(s string) {
		words := strings.Split(s, " ")
		for _, word := range words {
			freq[word]++
		}
	}

	insert(s1)
	insert(s2)

	ans := []string{}
	for word, occ := range freq {
		if occ == 1 {
			ans = append(ans, word)
		}
	}
	return ans
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/uncommon-words-from-two-sentences/solutions/1237687/liang-ju-hua-zhong-de-bu-chang-jian-dan-a8bmz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

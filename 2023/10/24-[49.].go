package leetcode

import "sort"

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
//
//
// 示例 1:
//
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
//
//
//输入: strs = [""]
//输出: [[""]]
//
//
// 示例 3:
//
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
//
// Related Topics 数组 哈希表 字符串 排序 👍 1685 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {

	l := make([][26]int, len(strs))
	for k, str := range strs {
		for i := 0; i < len(str); i++ {
			l[k][str[i]-'a']++
		}
	}

	// 1 22 333 4444
	sort.Slice(l, func(i, j int) bool {
		return l[i] == l[j]
	})

	var a [][]string
	var b []string
	for i := 0; i < len(l); i++ {
		if i == 0 {
			b = append(b, strs[i])
			continue
		}
		if l[i] == l[i-1] {
			b = append(b, strs[i])
		} else {
			var tmp []string
			a = append(a, append(tmp, b...))
			b = []string{strs[i]}
		}
	}
	return a

}

//leetcode submit region end(Prohibit modification and deletion)

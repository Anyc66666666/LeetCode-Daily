package leetcode

import (
	"strings"
)

//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答
//案。
//
//
//
// 示例 1：
//
//
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
//
//
// 示例 2：
//
//
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] 由小写英文字母组成
//
//
// Related Topics 数组 哈希表 字符串 👍 351 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func commonChars(words []string) []string {
	var a []string
	if len(words) == 1 {
		for _, v := range words[0] {
			a = append(a, string(v))
		}
		return a
	}

	for _, v := range words[0] {
		ok := true
		for i := 1; i < len(words); i++ {
			if !strings.Contains(words[i], string(v)) {
				ok = false
				break
			}
			words[i] = strings.Replace(words[i], string(v), " ", 1)

		}
		if ok {
			a = append(a, string(v))
		}

	}

	return a

}

//leetcode submit region end(Prohibit modification and deletion)

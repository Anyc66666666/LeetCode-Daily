package leetcode

import "strings"

//给定一个段落 (paragraph) 和一个禁用单词列表 (banned)。返回出现次数最多，同时不在禁用列表中的单词。
//
// 题目保证至少有一个词不在禁用列表中，而且答案唯一。
//
// 禁用列表中的单词用小写字母表示，不含标点符号。段落中的单词不区分大小写。答案都是小写字母。
//
//
//
// 示例：
//
// 输入:
//paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
//banned = ["hit"]
//输出: "ball"
//解释:
//"hit" 出现了3次，但它是一个禁用的单词。
//"ball" 出现了2次 (同时没有其他单词出现2次)，所以它是段落里出现次数最多的，且不在禁用列表中的单词。
//注意，所有这些单词在段落里不区分大小写，标点符号需要忽略（即使是紧挨着单词也忽略， 比如 "ball,"），
//"hit"不是最终的答案，虽然它出现次数更多，但它在禁用单词列表中。
//
//
//
//
// 提示：
//
//
// 1 <= 段落长度 <= 1000
// 0 <= 禁用单词个数 <= 100
// 1 <= 禁用单词长度 <= 10
// 答案是唯一的, 且都是小写字母 (即使在 paragraph 里是大写的，即使是一些特定的名词，答案都是小写的。)
// paragraph 只包含字母、空格和下列标点符号!?',;.
// 不存在没有连字符或者带有连字符的单词。
// 单词里只包含字母，不会出现省略号或者其他标点符号。
//
//
// Related Topics 哈希表 字符串 计数 👍 215 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]struct{})
	list := make(map[string]int)
	for i := 0; i < len(banned); i++ {
		bannedMap[banned[i]] = struct{}{}
	}

	strSlice := strings.Split(paragraph, " ")

	for _, v := range strSlice {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		s := strings.ToLower(v)
		if s[len(s)-1] < 'a' || s[len(s)-1] > 'z' {
			s = s[:len(s)-1]
		}
		a := strings.Split(s, ",")
		if len(a) > 0 {
			for _, vv := range strings.Split(s, ",") {
				if _, ok := bannedMap[vv]; !ok {
					list[vv]++
				}
			}
			if _, ok := bannedMap[s]; !ok {
				list[s]++
			}
		}
	}
	var a int
	var ss string
	for k, v := range list {
		if v > a {
			a = v
			ss = k
		}
	}
	return ss

}

//leetcode submit region end(Prohibit modification and deletion)

/*
为了判断给定段落中的每个单词是否在禁用单词列表中，需要使用哈希集合存储禁用单词列表
中的单词。以下将禁用单词列表中的单词称为禁用单词。
遍历段落
paragraph，得到段落中的所有单词，并对每个单词计数，使用哈希表记录每个单词
的计数。由于每个单词由连续的字母组成，因此当遇到一个非字母的字符且该字符的前一-个字:
符是字母时，即为一个单词的结束，如果该单词不是禁用单词，则将该单词的计数加1。如果
段落的最后一个字符是字母，则当遍历结束时需要对段落中的最后-一个单词判断是否为禁用单
词，如果不是禁用单词则将次数加1。
在遍历段落的过程中，对于每个单词都会更新计数，因此遍历结束之后即可得到最大计数，即
出现次数最多的单词的出现次数。
遍历段落之后，遍历哈希表，寻找出现次数等于最大计数的单词，该单词即为最常见的单词。


func mostCommonWord(paragraph string, banned []string) string {
    ban := map[string]bool{}
    for _, s := range banned {
        ban[s] = true
    }
    freq := map[string]int{}
    maxFreq := 0
    var word []byte
    for i, n := 0, len(paragraph); i <= n; i++ {
        if i < n && unicode.IsLetter(rune(paragraph[i])) {
            word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
        } else if word != nil {
            s := string(word)
            if !ban[s] {
                freq[s]++
                maxFreq = max(maxFreq, freq[s])
            }
            word = nil
        }
    }
    for s, f := range freq {
        if f == maxFreq {
            return s
        }
    }
    return ""
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/most-common-word/solutions/1424731/zui-chang-jian-de-dan-ci-by-leetcode-sol-mzjb/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

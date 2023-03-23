package leetcode

//给你一个字符串 licensePlate 和一个字符串数组 words ，请你找出 words 中的 最短补全词 。
//
// 补全词 是一个包含 licensePlate 中所有字母的单词。忽略 licensePlate 中的 数字和空格 。不区分大小写。如果某个字母在
//licensePlate 中出现不止一次，那么该字母在补全词中的出现次数应当一致或者更多。
//
// 例如：licensePlate = "aBc 12c"，那么它的补全词应当包含字母 'a'、'b' （忽略大写）和两个 'c' 。可能的 补全词 有
//"abccdef"、"caaacab" 以及 "cbca" 。
//
// 请返回 words 中的 最短补全词 。题目数据保证一定存在一个最短补全词。当有多个单词都符合最短补全词的匹配条件时取 words 中 第一个 出现的那个
//。
//
//
//
// 示例 1：
//
//
//输入：licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
//输出："steps"
//解释：最短补全词应该包括 "s"、"p"、"s"（忽略大小写） 以及 "t"。
//"step" 包含 "t"、"p"，但只包含一个 "s"，所以它不符合条件。
//"steps" 包含 "t"、"p" 和两个 "s"。
//"stripe" 缺一个 "s"。
//"stepple" 缺一个 "s"。
//因此，"steps" 是唯一一个包含所有字母的单词，也是本例的答案。
//
// 示例 2：
//
//
//输入：licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
//输出："pest"
//解释：licensePlate 只包含字母 "s" 。所有的单词都包含字母 "s" ，其中 "pest"、"stew"、和 "show" 三者最短。答案是
//"pest" ，因为它是三个单词中在 words 里最靠前的那个。
//
//
//
//
// 提示：
//
//
// 1 <= licensePlate.length <= 7
// licensePlate 由数字、大小写字母或空格 ' ' 组成
// 1 <= words.length <= 1000
// 1 <= words[i].length <= 15
// words[i] 由小写英文字母组成
//
//
// Related Topics 数组 哈希表 字符串 👍 123 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func shortestCompletingWord(licensePlate string, words []string) string {
	type Name struct {
		Str string
		Id  int
		Len int
	}

	plate := make(map[byte]int)
	var str []Name
	for i := 0; i < len(licensePlate); i++ {
		if (licensePlate[i] >= '0' && licensePlate[i] <= '9') || licensePlate[i] == ' ' {
			continue
		}
		if licensePlate[i] < 97 {
			plate[licensePlate[i]+32]++
			continue
		}

		plate[licensePlate[i]]++
	}

	for k, v := range words {
		n := make(map[byte]int)
		ok := true
		for i := 0; i < len(v); i++ {
			n[v[i]]++
		}

		for kk, vv := range plate {
			if n[kk] < vv {
				ok = false
			}
		}

		if ok {
			str = append(str, Name{Str: v, Id: k, Len: len(v)})
		}

	}
	minLen := str[0].Len
	minId := str[0].Id
	ss := str[0].Str
	for _, v := range str {
		if v.Len == minLen {
			minLen = v.Len
			if v.Id < minId {
				minId = v.Id
				ss = v.Str
			}
			continue
		}
		if v.Len < minLen {
			minLen = v.Len
			minId = v.Id
			ss = v.Str
		}
	}

	return ss

}

//leetcode submit region end(Prohibit modification and deletion)

/*
func shortestCompletingWord(licensePlate string, words []string) (ans string) {
    cnt := [26]int{}
    for _, ch := range licensePlate {
        if unicode.IsLetter(ch) {
            cnt[unicode.ToLower(ch)-'a']++
        }
    }

next:
    for _, word := range words {
        c := [26]int{}
        for _, ch := range word {
            c[ch-'a']++
        }
        for i := 0; i < 26; i++ {
            if c[i] < cnt[i] {
                continue next
            }
        }
        if ans == "" || len(word) < len(ans) {
            ans = word
        }
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/shortest-completing-word/solutions/1147475/zui-duan-bu-quan-ci-by-leetcode-solution-35pt/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

package leetcode

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//
//
// 示例 1:
//
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "rat", t = "car"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s 和 t 仅包含小写字母
//
//
//
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
// Related Topics 哈希表 字符串 排序 👍 855 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	//if len(s)!=len(t){
	//	return false
	//}
	//m:=make(map[byte]int)
	//for i:=0;i<len(s);i++{
	//	m[s[i]]++
	//	m[t[i]]--
	//}
	//for _,v:=range m{
	//	if v!=0{
	//		return false
	//	}
	//}
	//return true

	//------------------

	//s1, s2 := []byte(s), []byte(t)
	//sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	//sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	//return string(s1) == string(s2)

	//------------------------
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2

}

//leetcode submit region end(Prohibit modification and deletion)

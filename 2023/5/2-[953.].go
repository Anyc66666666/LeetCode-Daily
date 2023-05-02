package leetcode

//某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
//
// 给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回
//false。
//
//
//
// 示例 1：
//
//
//输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
//输出：true
//解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。
//
// 示例 2：
//
//
//输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
//输出：false
//解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。
//
// 示例 3：
//
//
//输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
//输出：false
//解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，因为 'l' > '∅'，其中 '∅
//' 是空白字符，定义为比任何其他字符都小（更多信息）。
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 20
// order.length == 26
// 在 words[i] 和 order 中的所有字符都是英文小写字母。
//
//
// Related Topics 数组 哈希表 字符串 👍 242 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isAlienSorted(words []string, order string) bool {
	index := [26]int{}
	for i, c := range order {
		index[c-'a'] = i //保存每个字母对位置
	}
next:
	for i := 1; i < len(words); i++ { //前后两个进行比较
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			pre, cur := index[words[i-1][j]-'a'], index[words[i][j]-'a']
			if pre > cur {
				return false
			}
			if pre < cur { // b < c
				continue next //比完后，比下一个字符串的相同位置的字符
			}
			//ab ac      //这里a都是相同的，所有开始比a后面的字符直到不相同为止
			//如果per==cur，则比下一个字符
		}

		if len(words[i-1]) > len(words[i]) {
			return false
		}
	}
	//先比较所有字符串中的前后两个字符串的第一个字符，然后再比较第二个字符
	//eg:先比较第一个和第二个字符串的第一个字符，满足则然后比较第二个和第三个字符串的第一个字符..
	//第一个字符全部被比较后，再从第一个字符串和第二个字符串开始比较第二个字符..

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

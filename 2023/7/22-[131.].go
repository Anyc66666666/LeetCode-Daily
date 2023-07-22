package leetcode

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 回溯 👍 1567 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	var ans [][]string
	var a []string
	var dfs func(str string, now int)
	dfs = func(str string, now int) {
		length := len(str)
		if now == length {
			//如果起始位置等于s的大小，说明已经找到了一组分割方案了
			tmp := make([]string, len(a))
			copy(tmp, a)
			ans = append(ans, tmp)
			return
		}

		for i := now; i < length; i++ {
			ss := str[now : i+1]
			if help(ss) {
				a = append(a, ss)
				dfs(str, i+1)
				a = a[:len(a)-1]
			}
		}

	}

	dfs(s, 0)
	return ans

}

func help(str string) bool {
	length := len(str)
	if length == 1 {
		return true
	}
	// a b b a
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)

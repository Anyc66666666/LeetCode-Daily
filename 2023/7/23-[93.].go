package leetcode

import "strconv"
import "strings"

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
//和 "192.168@1.1" 是 无效 IP 地址。
//
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
//排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
//
// 示例 2：
//
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//
//
// 示例 3：
//
//
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// s 仅由数字组成
//
//
// Related Topics 字符串 回溯 👍 1254 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	//	//输入：s = "101023"
	//	//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
	//	var ans []string
	//	var a []string
	//	length:=len(s)
	//	var dfs func(str string,now int)
	//	dfs= func(str string,now int) {
	//		if len(a)==4{
	//
	//			x:=strings.Join(a,".")
	//			if len(x)!=length+3{
	//				return
	//			}
	//
	//			ans=append(ans,x)
	//			return
	//		}
	//
	//		for i:=now;i<length&&i<now+3;i++{
	//			m:=str[now:i+1]
	//			if len(m)!=1&&m[0]=='0'{
	//				break
	//			}
	//			n,err:=strconv.Atoi(m)
	//			if err != nil {
	//				continue
	//			}
	//			if n>255{
	//				continue
	//			}
	//			a=append(a,m)
	//			dfs(str,i+1)
	//			a=a[:len(a)-1]
	//		}
	//
	//	}
	//
	//	dfs(s,0)
	//	return ans

	//------------
	var path, res []string
	var dfs func(s string, start int)

	dfs = func(s string, start int) {
		if len(path) == 4 {
			if start == len(s) {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}

		for i := start; i < len(s); i++ {
			if i != start && s[start] == '0' {
				break
			}
			str := s[start : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				path = append(path, str)
				dfs(s, i+1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}

	}

	dfs(s, 0)
	return res

}

//leetcode submit region end(Prohibit modification and deletion)

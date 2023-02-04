package leetcode

import "fmt"

//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abcdefg", k = 2
//输出："bacdfeg"
//
//
// 示例 2：
//
//
//输入：s = "abcd", k = 2
//输出："bacd"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= k <= 10⁴
//
//
// Related Topics 双指针 字符串 👍 409 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	//1 2 3 4   5 6 7 8
	var m int
	var ss []string
	for _, v := range s {
		ss = append(ss, string(v))
	}
	for i := 0; i < len(s); i++ {
		m++

		if len(s) < 2*k {
			if len(s) < k {
				for j := 0; j <= (len(s)-1)/2; j++ {
					ss[j], ss[len(s)-1-j] = ss[len(s)-1-j], ss[j]
				}
				break
			}
			if len(s) >= k && len(s) <= 2*k {
				fmt.Println("12312312312312")
				for j := 0; j <= (k-1)/2; j++ {
					ss[j], ss[k-1-j] = ss[k-1-j], ss[j]
				}
				//for j:=m-k+1;j<len(s);j++{
				//	s=s+string(s[j])
				//}
				break
			}

		}

		if m%(2*k) == 0 { // [m-2k,m-k-1}
			for j := m - 2*k; j <= (2*m-(3*k)-1)/2; j++ {
				ss[j], ss[2*m-3*k-1-j] = ss[2*m-3*k-1-j], ss[j]
			}
			// 4 5 6 7 8
			if len(s)-m < k { // [m,len(s)-1]
				for j := m; j <= (len(s)-1+m)/2; j++ {
					ss[j], ss[len(s)+m-1-j] = ss[len(s)+m-1-j], ss[j]
				}
				fmt.Println("6666")
				break
			}
			if len(s)-m >= k && len(s)-m < 2*k {
				fmt.Println("12312312312312")
				// m m+k-1
				for j := m; j <= (2*m+k-1)/2; j++ {
					ss[j], ss[2*m+k-1-j] = ss[2*m+k-1-j], ss[j]
				}
				//for j:=m-k+1;j<len(s);j++{
				//	s=s+string(s[j])
				//}
				break
			}

			continue
		}

	}

	var sss string
	for _, v := range ss {
		sss = sss + v
	}
	return sss

}

//leetcode submit region end(Prohibit modification and deletion)

/*
func reverseStr(s string, k int) string {
    t := []byte(s)
    for i := 0; i < len(s); i += 2 * k {
        sub := t[i:min(i+k, len(s))]
        for j, n := 0, len(sub); j < n/2; j++ {
            sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
        }
    }
    return string(t)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

*/

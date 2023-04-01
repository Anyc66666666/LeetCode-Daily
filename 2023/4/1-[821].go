package leetcode

//给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。
//
// 返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的
//字符 c 的 距离 。
//
// 两个下标 i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。
//
//
//
// 示例 1：
//
//
//输入：s = "loveleetcode", c = "e"
//输出：[3,2,1,0,1,0,0,1,2,2,1,0]
//解释：字符 'e' 出现在下标 3、5、6 和 11 处（下标从 0 开始计数）。
//距下标 0 最近的 'e' 出现在下标 3 ，所以距离为 abs(0 - 3) = 3 。
//距下标 1 最近的 'e' 出现在下标 3 ，所以距离为 abs(1 - 3) = 2 。
//对于下标 4 ，出现在下标 3 和下标 5 处的 'e' 都离它最近，但距离是一样的 abs(4 - 3) == abs(4 - 5) = 1 。
//距下标 8 最近的 'e' 出现在下标 6 ，所以距离为 abs(8 - 6) = 2 。
//
//
// 示例 2：
//
//
//输入：s = "aaab", c = "b"
//输出：[3,2,1,0]
//
//
//
//提示：
//
//
// 1 <= s.length <= 10⁴
// s[i] 和 c 均为小写英文字母
// 题目数据保证 c 在 s 中至少出现一次
//
//
// Related Topics 数组 双指针 字符串 👍 332 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func shortestToChar(s string, c byte) []int {
	var cc []int
	var l []int
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cc = append(cc, i)
		}
	}
	for i := 0; i < len(s); i++ {
		m := getMin(i, cc)
		l = append(l, m)
	}
	return l

}
func getMin(a int, cc []int) int {
	min := abs(a, cc[0])
	for _, v := range cc {
		d := abs(a, v)
		if d < min {
			min = d
		}

	}
	return min
}
func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

//leetcode submit region end(Prohibit modification and deletion)

/*

问题可以转换成，对s的每个下标i,求
●s[i]到其左侧最近的字符c的距离
●s[i] 到其右侧最近的字符c的距离
这两者的最小值。
对于前者，我们可以从左往右遍历s，若s[i]=c则记录下此时字符c的的下标idx。遍历的
同时更新answer[i] = i - idx。
对于后者，我们可以从右往左遍历s，若s[i]=c则记录下此时字符c的的下标idx。遍历的
同时更新answer[i] = min(answer[i, idx - i)。
代码实现时，在开始遍历的时候idx 可能不存在，为了简化逻辑，我们可以用-n或2n表
示，这里n是s的长度。


func shortestToChar(s string, c byte) []int {
    n := len(s)
    ans := make([]int, n)

    idx := -n
    for i, ch := range s {
        if byte(ch) == c {
            idx = i
        }
        ans[i] = i - idx
    }

    idx = n * 2
    for i := n - 1; i >= 0; i-- {
        if s[i] == c {
            idx = i
        }
        ans[i] = min(ans[i], idx-i)
    }
    return ans
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/shortest-distance-to-a-character/solutions/1429810/zi-fu-de-zui-duan-ju-chi-by-leetcode-sol-2t49/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

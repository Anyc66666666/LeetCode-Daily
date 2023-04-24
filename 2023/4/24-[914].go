package leetcode

//给定一副牌，每张牌上都写着一个整数。
//
// 此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
//
//
// 每组都有 X 张牌。
// 组内所有的牌上都写着相同的整数。
//
//
// 仅当你可选的 X >= 2 时返回 true。
//
//
//
// 示例 1：
//
//
//输入：deck = [1,2,3,4,4,3,2,1]
//输出：true
//解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
//
//
// 示例 2：
//
//
//输入：deck = [1,1,1,2,2,2,3,3]
//输出：false
//解释：没有满足要求的分组。
//
//
// 提示：
//
//
// 1 <= deck.length <= 10⁴
// 0 <= deck[i] < 10⁴
//
//
// Related Topics 数组 哈希表 数学 计数 数论 👍 282 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)

	for _, v := range deck {
		m[v]++

	}
	a := m[deck[0]]
	for _, v := range m {
		if v < 2 {
			return false
		}
		if v != a {
			p := getPublic(a, v)
			if p < 2 {
				return false
			}
			a = v

			//return false
		}

	}
	return true

}

func getPublic(a, b int) int {
	//4 6
	if a < b {
		a, b = b, a
	}

	//a>=b
	if a%b == 0 {
		return b
	}

	if a%2 == 0 && b%2 == 0 {
		return 2
	}

	if a%3 == 0 && b%3 == 0 {
		return 3
	}

	if a%5 == 0 && b%5 == 0 {
		return 3
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

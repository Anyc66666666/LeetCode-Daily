package leetcode

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2619 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	length := len(nums)
	var ans [][]int
	var a []int
	used := make([]bool, length)
	var dfs func()
	dfs = func() {
		if len(a) == length {
			//
			//n := make(map[int]struct{})
			//for _, v := range a {
			//	if _, ok := n[v]; ok {
			//		return
			//	} else if !ok {
			//		n[v] = struct{}{}
			//	}
			//}
			m := make([]int, length)
			copy(m, a)
			ans = append(ans, m)
			return

		}

		for i := 0; i < length; i++ {
			if !used[i] {

				used[i] = true
				a = append(a, nums[i])
				dfs()
				used[i] = false
				a = a[:len(a)-1]
			}

		}

	}

	dfs()
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

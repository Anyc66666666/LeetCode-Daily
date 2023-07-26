package leetcode

import "sort"

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
//
//
// Related Topics 位运算 数组 回溯 👍 1125 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	//sort.Ints(nums)
	//var ans [][]int
	//var a []int
	//var dfs func(now int)
	//n := make(map[string]struct{})
	//dfs = func(now int) {
	//	m := make([]int, len(a))
	//	copy(m, a)
	//	var s string
	//	for _, v := range m {
	//		s = s + string(v)
	//	}
	//	if _, ok := n[s]; ok {
	//		s = ""
	//		return
	//	} else {
	//		n[s] = struct{}{}
	//		ans = append(ans, m)
	//		s = ""
	//	}
	//
	//	for i := now; i < len(nums); i++ {
	//		//if i!=0&&nums[i]==nums[i-1]{
	//		//
	//		//}
	//
	//		a = append(a, nums[i])
	//		dfs(i + 1)
	//		a = a[:len(a)-1]
	//
	//	}
	//
	//}
	//
	//dfs(0)
	//return ans

	sort.Ints(nums)
	var ans [][]int
	var a []int
	var dfs func(now int)

	dfs = func(now int) {
		m := make([]int, len(a))
		copy(m, a)

		ans = append(ans, m)

		for i := now; i < len(nums); i++ {
			if i != now && nums[i] == nums[i-1] {
				continue
			}

			a = append(a, nums[i])
			dfs(i + 1)
			a = a[:len(a)-1]

		}

	}

	dfs(0)
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

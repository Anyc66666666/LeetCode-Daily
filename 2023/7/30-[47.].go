package leetcode

import "sort"

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 👍 1409 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var a []int
	use := make([]bool, len(nums))
	sort.Ints(nums)
	//n:=make(map[string]struct{})

	var dfs func()
	dfs = func() {
		if len(a) == len(nums) {
			//var str string
			//for _,v:=range a{
			//	str=str+fmt.Sprint(v)
			//}
			//if _,ok:=n[str];ok{
			//	return
			//}else if !ok{
			//	n[str]= struct{}{}
			//	m:=make([]int,len(nums))
			//	copy(m,a)
			//	ans=append(ans,m)
			//}

			m := make([]int, len(nums))
			copy(m, a)
			ans = append(ans, m)

		}

		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] && use[i-1] {
				continue
			}

			if !use[i] {
				use[i] = true
				a = append(a, nums[i])
				dfs()
				a = a[:len(a)-1]
				use[i] = false
			}

		}

	}

	dfs()
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

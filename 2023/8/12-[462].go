package leetcode

import "sort"

//给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最小操作数。
//
// 在一次操作中，你可以使数组中的一个元素加 1 或者减 1 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：2
//解释：
//只需要两次操作（每次操作指南使一个元素加 1 或减 1）：
//[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
//
//
// 示例 2：
//
//
//输入：nums = [1,10,2,9]
//输出：16
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 数组 数学 排序 👍 292 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minMoves2(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	mid := nums[length/2]
	var a int

	for _, v := range nums {
		if mid > v {
			a = a + mid - v
			continue
		}
		a = a + v - mid

	}

	return a

}

//leetcode submit region end(Prohibit modification and deletion)

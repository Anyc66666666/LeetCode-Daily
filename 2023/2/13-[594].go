package leetcode

import "sort"

//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
//
// 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
//
// 数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,2,2,5,2,3,7]
//输出：5
//解释：最长的和谐子序列是 [3,2,2,2,3]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,4]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [1,1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 数组 哈希表 排序 👍 354 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findLHS(nums []int) int {
	var a int
	sort.Ints(nums)
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1]+1 {
			if m[nums[i]]+m[nums[i-1]] > a {
				a = m[nums[i]] + m[nums[i-1]]
			}
		}
	}
	return a

}

//leetcode submit region end(Prohibit modification and deletion)

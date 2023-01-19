package leetcode

//给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,0,1,1,1]
//输出：3
//解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
//
//
// 示例 2:
//
//
//输入：nums = [1,0,1,1,0,1]
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 不是 0 就是 1.
//
//
// Related Topics 数组 👍 364 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
	var a int
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if m[1] > a {
				a = m[1]
			}
			m[1] = 0
			continue
		}
		m[1]++
	}
	if m[1] > a {
		a = m[1]
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

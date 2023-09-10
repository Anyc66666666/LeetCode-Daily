package leetcode

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 3：
//
//
//输入：nums = [1,2,3]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1447 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func rob2(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	var a, b []int
	a = append(a, nums[1:]...)
	b = append(b, nums[:len(nums)-1]...)

	length := len(a)
	dp1 := make([]int, length+1)
	dp2 := make([]int, length+1)
	dp1[1] = a[0]
	dp2[1] = b[0]
	for i := 2; i <= length; i++ {
		dp1[i] = max7(dp1[i-1], dp1[i-2]+a[i-1])
	}
	for i := 2; i <= length; i++ {
		dp2[i] = max7(dp2[i-1], dp2[i-2]+b[i-1])
	}

	return max7(dp1[length], dp2[length])
}
func max7(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

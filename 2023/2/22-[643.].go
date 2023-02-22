package leetcode

//给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//
// 任何误差小于 10⁻⁵ 的答案都将被视为正确答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//
//
// 示例 2：
//
//
//输入：nums = [5], k = 1
//输出：5.00000
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= k <= n <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 滑动窗口 👍 282 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum = sum + v
	}
	maxNum := sum
	//只关注数组区间的最左边和最右边，左边少一个，右边多一个。直到右边到达最后一个
	//滑动窗口
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxNum = max1(maxNum, sum)
	}
	return float64(maxNum) / float64(k)

}
func max11(a, b int) int {
	if a > b {
		return a
	}
	return b

}

//leetcode submit region end(Prohibit modification and deletion)

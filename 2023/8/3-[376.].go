package leetcode

//如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也
//视作摆动序列。
//
//
// 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
// 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一
//个差值为零。
//
//
// 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
//
// 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,7,4,9,2,5]
//输出：6
//解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。
//
//
// 示例 2：
//
//
//输入：nums = [1,17,5,10,13,15,10,5,16,8]
//输出：7
//解释：这个序列包含几个长度为 7 摆动序列。
//其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。
//
//
// 示例 3：
//
//
//输入：nums = [1,2,3,4,5,6,7,8,9]
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 1000
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度完成此题?
//
// Related Topics 贪心 数组 动态规划 👍 981 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func wiggleMaxLength(nums []int) int {

	//--贪心
	//n:=len(nums)
	//if n<2{
	//	return n
	//}
	//ans:=1
	//pre:=nums[1]-nums[0]
	//if pre!=0{
	//	ans=2
	//}
	//
	//for i:=2;i<len(nums);i++{
	//	now:=nums[i]-nums[i-1]
	//	if now<0&&pre>=0||now>0&&pre<=0{
	//		ans++
	//		pre=now
	//	}
	//}
	//
	//return ans

	//--动态规划
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([][2]int, n)
	//i 0 波峰最大长度
	//i 1 波谷最大长度

	dp[0][0] = 1
	dp[0][1] = 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				//nums[i]为波谷
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}

			if nums[j] < nums[i] {
				//nums[i]为波峰
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
			if nums[j] == nums[i] {
				dp[i][0] = max(dp[i][0], dp[j][0]) //波峰
				dp[i][1] = max(dp[i][1], dp[j][1]) //波谷
			}
		}
	}

	return max(dp[n-1][0], dp[n-1][1])

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
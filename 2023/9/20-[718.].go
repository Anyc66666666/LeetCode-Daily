package leetcode

//给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//输出：3
//解释：长度最长的公共子数组是 [3,2,1] 。
//
//
// 示例 2：
//
//
//输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
//输出：5
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 100
//
//
// Related Topics 数组 二分查找 动态规划 滑动窗口 哈希函数 滚动哈希 👍 998 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findLength(nums1 []int, nums2 []int) int {
	//var length int
	//for i:=1;i<len(nums1)+len(nums2);i++{
	//	if i<=len(nums2){
	//		le:=0
	//		for a:=0;a<i;a++{
	//			if nums2[a]==nums1[len(nums1)-i+a]{
	//				le++
	//			}else {
	//				le=0
	//			}
	//		}
	//		length=max991(length,le)
	//	}else if i <=len(nums1){
	//		// 2 3 4 5 6 // 1 2 3
	//		// 1 2 3         1 2 3
	//		// 2 3 4 5 6   2 3 4 5 6
	//		le:=0
	//		for a:=i-len(nums2);a<len(nums2);a++{
	//			if nums1[a]==nums2[i-len(nums2)+a]{
	//				le++
	//			}else {
	//				le=0
	//			}
	//		}
	//		length=max991(length,le)
	//
	//	}
	//}
	//
	//return length

	n, m := len(nums1), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans

}
func max991(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

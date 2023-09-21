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
	n, m := len(nums1), len(nums2)
	ret := 0
	for i := 0; i < n; i++ {
		length := min991(m, n-i)
		maxLen := maxLength(nums1, nums2, i, 0, length)
		ret = max991(ret, maxLen)
	}
	for i := 0; i < n; i++ {
		length := min991(n, m-i)
		maxLen := maxLength(nums1, nums2, 0, i, length)
		ret = max991(ret, maxLen)
	}
	return ret

	//我们可以枚举 A 和 B 所有的对齐方式。
	//对齐的方式有两类：第一类为 A 不变，B 的首元素与 A 中的某个元素对齐；
	//第二类为 B 不变，A 的首元素与 B 中的某个元素对齐。
	//对于每一种对齐方式，我们计算它们相对位置相同的重复子数组即可。

}
func maxLength(a, b []int, addA, addB, length int) int {
	ret, k := 0, 0
	for i := 0; i < length; i++ {
		if a[addA+i] == b[addB+i] {
			k++
		} else {
			k = 0
		}
		ret = max991(ret, k)
	}
	return ret
}
func max991(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min991(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

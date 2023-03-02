package leetcode

//给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那
//么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,5,4,7]
//输出：3
//解释：最长连续递增序列是 [1,3,5], 长度为3。
//尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2]
//输出：1
//解释：最长连续递增序列是 [2], 长度为1。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 数组 👍 360 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {

	//1 3 5 4 7
	//8 9 1 3 5 4
	var a int
	for i := 0; i < len(nums); i++ {
		var b int

		for j := i; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[j] > nums[j-1] {
				b++
			} else {
				break
			}
		}
		b++
		a = maxX(a, b)

	}
	return a

}

func maxX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
/*
 贪心
func findLengthOfLCIS(nums []int) (ans int) {
    start := 0
    for i, v := range nums {
        if i > 0 && v <= nums[i-1] {
            start = i
        }
        ans = max(ans, i-start+1)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/longest-continuous-increasing-subsequence/solutions/573383/zui-chang-lian-xu-di-zeng-xu-lie-by-leet-dmb8/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

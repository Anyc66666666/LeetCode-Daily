package leetcode

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
// 示例 2：
//
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 已按 非递减顺序 排序
//
//
//
//
// 进阶：
//
//
// 请你设计时间复杂度为 O(n) 的算法解决本问题
//
//
// Related Topics 数组 双指针 排序 👍 797 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	//a := make([]int, 0, len(nums))
	//
	//for _, v := range nums {
	//	a = append(a, v*v)
	//}
	//sort.Ints(a)
	//return a

	n := len(nums)
	a := make([]int, n)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if l, r := nums[left]*nums[left], nums[right]*nums[right]; l > r {
			a[i] = l
			left++
		} else {
			a[i] = r
			right--
		}
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
/*
因为原切片是非递减的，
所以平方之后，两端的值会大于中间的值，且最大的值在两端，最小的值在中间。
可以考虑逆序放入答案
总是去比较最左边和最右边的值，拿到最大的逆序放入切片中，相当于左右双指针
放入之后，就改变该指针的位置

*/

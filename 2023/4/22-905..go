package leetcode

//给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
//
// 返回满足此条件的 任一数组 作为答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,1,2,4]
//输出：[2,4,3,1]
//解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// 0 <= nums[i] <= 5000
//
//
// Related Topics 数组 双指针 排序 👍 353 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParity(nums []int) []int {
	var a, b []int
	for _, v := range nums {
		if v%2 == 0 {
			a = append(a, v)
			continue
		}
		b = append(b, v)

	}
	a = append(a, b...)
	return a

}

//leetcode submit region end(Prohibit modification and deletion)

/*

func sortArrayByParity(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    left, right := 0, n-1
    for _, num := range nums {
        if num%2 == 0 {
            ans[left] = num
            left++
        } else {
            ans[right] = num
            right--
        }
    }
    return ans
}


*/

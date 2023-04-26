package leetcode

//给定一个非负整数数组 nums， nums 中一半整数是 奇数 ，一半整数是 偶数 。
//
// 对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。
//
// 你可以返回 任何满足上述条件的数组作为答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,2,5,7]
//输出：[4,5,2,7]
//解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
//
//
// 示例 2：
//
//
//输入：nums = [2,3]
//输出：[2,3]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 2 * 10⁴
// nums.length 是偶数
// nums 中一半是偶数
// 0 <= nums[i] <= 1000
//
//
//
//
// 进阶：可以不使用额外空间解决问题吗？
//
// Related Topics 数组 双指针 排序 👍 263 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(nums []int) []int {
	length := len(nums)

	o := make([]int, 0, length/2)
	j := make([]int, 0, length/2)
	for _, v := range nums {
		if v%2 == 0 {
			o = append(o, v)
			continue
		}
		j = append(j, v)
	}

	for i := 0; i < length; i = i + 2 {
		nums[i] = o[i/2]
		nums[i+1] = j[i/2]
	}
	return nums

}

//leetcode submit region end(Prohibit modification and deletion)

/*

func sortArrayByParityII(nums []int) []int {
    ans := make([]int, len(nums))
    i := 0
    for _, v := range nums {
        if v%2 == 0 {
            ans[i] = v
            i += 2
        }
    }
    i = 1
    for _, v := range nums {
        if v%2 == 1 {
            ans[i] = v
            i += 2
        }
    }
    return ans
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/sort-array-by-parity-ii/solutions/481450/an-qi-ou-pai-xu-shu-zu-ii-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package leetcode

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
//字，并以数组的形式返回结果。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1]
//输出：[2]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
//
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
//
// Related Topics 数组 哈希表 👍 1138 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	a := make([]int, 0)
	m := make(map[int]struct{}, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			a = append(a, i)
		}
	}
	return a

}

//leetcode submit region end(Prohibit modification and deletion)
/*
func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
	v = (v - 1) % n
	nums[v] += n
	    }


	for i, v := range nums {
	if v <= n {
	ans = append(ans, i+1)
	       }
	    }

	return

	}

由于 nums 的数字范围均在 [1,n] 中，我们可以利用这一范围之外的数字，来表达「是否存在」的含义。

具体来说，遍历 nums，每遇到一个数 x，就让 nums[x−1] 增加 n。由于 nums 中所有数均在 [1,n] 中，增加以后，这些数必然大于 n。最后我们遍历 nums，若 nums[i] 未大于 n，就说明没有遇到过数 i+1。这样我们就找到了缺失的数字。

注意，当我们遍历到某个位置时，其中的数可能已经被增加过，因此需要对 n 取模来还原出它本来的值。

//如果该元素不大于n，说明该值没有被加到，即该元素不存在，即该下标也不存在
*/

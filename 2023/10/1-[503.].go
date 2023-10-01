package leetcode

//给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素
// 。
//
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
//。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
//
// 示例 2:
//
//
//输入: nums = [1,2,3,4,3]
//输出: [2,3,4,-1,4]
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 栈 数组 单调栈 👍 869 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
// 我们可以使用单调栈解决本题。单调栈中保存的是下标，从栈底到栈顶的下标在数组 nums 中对应的值是单调不升的。
// 每次我们移动到数组中的一个新的位置 i，我们就将当前单调栈中所有对应值小于 nums[i] 的下标弹出单调栈，
// 这些值的下一个更大元素即为 nums[i]（证明很简单：如果有更靠前的更大元素，那么这些位置将被提前弹出栈）。
// 随后我们将位置 i 入栈。
// 但是注意到只遍历一次序列是不够的，例如序列 [2,3,1]，
// 最后单调栈中将剩余 [3,1] ,其中元素 [1] 的下一个更大元素还是不知道的。
// 一个朴素的思想是，我们可以把这个循环数组「拉直」，即复制该序列的前 n−1 个元素拼接在原序列的后面。
// 这样我们就可以将这个新序列当作普通序列，用上文的方法来处理。
// 而在本题中，我们不需要显性地将该循环数组「拉直」，而只需要在处理时对下标取模即可。
func nextGreaterElements(nums []int) []int {
	var stack []int
	length := len(nums)
	ans := make([]int, length)
	for i := range ans {
		ans[i] = -1
	}

	for i := 0; i < 2*length-1; i++ {
		for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums[i%length]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%length)
	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

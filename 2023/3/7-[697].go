package leetcode

//给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
//
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2,3,1]
//输出：2
//解释：
//输入数组的度是 2 ，因为元素 1 和 2 的出现频数最大，均为 2 。
//连续子数组里面拥有相同度的有如下所示：
//[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
//最短连续子数组 [2, 2] 的长度为 2 ，所以返回 2 。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,2,3,1,4,2]
//输出：6
//解释：
//数组的度是 3 ，因为元素 2 重复出现 3 次。
//所以 [2,2,3,1,4,2] 是最短子数组，因此返回 6 。
//
//
//
//
// 提示：
//
//
// nums.length 在 1 到 50,000 范围内。
// nums[i] 是一个在 0 到 49,999 范围内的整数。
//
//
// Related Topics 数组 哈希表 👍 460 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findShortestSubArray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	type Length struct {
		Left  int
		Right int
		Sum   int
	}
	m := make(map[int]Length)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; !ok {
			m[nums[i]] = Length{Left: i, Sum: 1}
		} else if ok {
			m[nums[i]] = Length{
				Left:  v.Left,
				Right: i,
				Sum:   v.Sum + 1,
			}
		}

	}
	var a, b int
	for _, v := range m {
		if v.Right == 0 {
			continue
		}
		if v.Sum > a {
			a = v.Sum
			b = v.Right - v.Left + 1
		} else if v.Sum == a {
			b = min9(b, v.Right-v.Left+1)
		}

	}
	if b == 0 {
		return 1
	}
	return b

}
func min9(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

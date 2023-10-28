package leetcode

import "sort"

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//
//
// 示例 3：
//
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 双指针 排序 👍 6472 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	// 先从小到大排序
	sort.Ints(nums)
	// 接收结果
	var res [][]int
	// 获取数组长度
	length := len(nums)
	// 边界处理，数字不足三个直接返回空
	if len(nums) < 3 {
		return res
	}
	// 开始循环第一个固定值
	for index, value := range nums {
		// 如果固定位的值已经大于0，因为已经排好序了，后面的两个指针对应的值也肯定大于0，则和不可能为0，所以返回
		if nums[index] > 0 {
			return res
		}
		// 排除值重复的固定位
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		// 指针初始位置，固定位右边第一个和数组最后一个
		l := index + 1
		r := length - 1
		// 开始移动两个指针
		for l < r {
			// 判断三个数字之和的三种情况
			sum := value + nums[l] + nums[r]
			switch {
			case sum == 0:
				// 将结果加入二元组
				res = append(res, []int{nums[index], nums[l], nums[r]})
				// 去重，如果l < r且下一个数字一样，则继续挪动
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				// 同理
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			case sum > 0:
				// 如果和大于 0，那就说明 right 的值太大，需要左移
				r -= 1
				// 如果和小于 0，那就说明 left 的值太小，需要右移
			case sum < 0:
				l += 1
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

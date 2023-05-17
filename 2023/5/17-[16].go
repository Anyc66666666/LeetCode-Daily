package leetcode

import "math"
import "sort"

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 双指针 排序 👍 1388 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	// -4 -1 1 2
	sort.Ints(nums)
	min := math.MaxInt
	var sum1 int
	length := len(nums)
	for i := 0; i < length-2; i++ {
		l, r := i+1, length-1
		for j := 0; j < length-i-2; j++ {
			if l >= r {
				break
			}
			sum := nums[i] + nums[l] + nums[r]
			//fmt.Println(sum)
			if sum == target {
				return target
			} else if sum < target {
				l++
				if min > target-sum {
					min = target - sum
					sum1 = sum
				}
			} else {
				r--
				if min > sum-target {
					min = sum - target
					sum1 = sum
				}
			}
		}
	}
	return sum1
}

//leetcode submit region end(Prohibit modification and deletion)

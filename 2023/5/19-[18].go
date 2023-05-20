package leetcode

import "sort"

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 双指针 排序 👍 1594 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	// -2 -1 0 0 1 2
	length := len(nums)
	var a [][]int

	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			//fmt.Println("jjj",nums[i],nums[j])

			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, length-1
			for m := 0; m < length-j-2; m++ {
				if l != j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r != length-1 && nums[r] == nums[r+1] {
					r--
					continue
				}

				if l >= r {
					break
				}
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					b := []int{nums[i], nums[j], nums[l], nums[r]}
					a = append(a, b)
					l++
				} else if sum < target {
					l++
				} else {
					r--
				}
			}

		}

	}
	return a

}

//leetcode submit region end(Prohibit modification and deletion)

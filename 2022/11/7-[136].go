package leetcode

import "fmt"

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 说明：
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
// 示例 1:
//
// 输入: [2,2,1]
//输出: 1
//
//
// 示例 2:
//
// 输入: [4,1,2,1,2]
//输出: 4
//
// Related Topics 位运算 数组 👍 2646 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0

}

//leetcode submit region end(Prohibit modification and deletion)

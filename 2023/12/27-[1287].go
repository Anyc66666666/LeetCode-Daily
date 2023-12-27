package leetcode

//给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。
//
// 请你找到并返回这个整数
//
//
//
// 示例：
//
//
//输入：arr = [1,2,2,6,6,6,6,7,10]
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10^4
// 0 <= arr[i] <= 10^5
//
//
// Related Topics 数组 👍 79 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findSpecialInteger(arr []int) int {
	m := make(map[int]int)
	count := len(arr) / 4
	for _, v := range arr {
		m[v]++
		if m[v] > count {
			return v
		}
	}
	return 0

}

//leetcode submit region end(Prohibit modification and deletion)

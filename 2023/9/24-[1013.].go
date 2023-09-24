package leetcode

//给你一个整数数组 arr，只有可以将其划分为三个和相等的 非空 部分时才返回 true，否则返回 false。
//
// 形式上，如果可以找出索引 i + 1 < j 且满足 (arr[0] + arr[1] + ... + arr[i] == arr[i + 1] +
//arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1])
//就可以将数组三等分。
//
//
//
// 示例 1：
//
//
//输入：arr = [0,2,1,-6,6,-7,9,1,2,0,1]
//输出：true
//解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
//
//
// 示例 2：
//
//
//输入：arr = [0,2,1,-6,6,7,9,-1,2,0,1]
//输出：false
//
//
// 示例 3：
//
//
//输入：arr = [3,3,6,5,-2,2,5,1,-9,4]
//输出：true
//解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
//
//
//
//
// 提示：
//
//
// 3 <= arr.length <= 5 * 10⁴
// -10⁴ <= arr[i] <= 10⁴
//
//
// Related Topics 贪心 数组 👍 201 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func canThreePartsEqualSum(arr []int) bool {
	n := len(arr)
	preSum := make([]int, n)
	sufSum := make([]int, n)
	preSum[0] = arr[0]
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1] + arr[i]
	}
	if preSum[n-1]%3 != 0 {
		return false
	}
	sufSum[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		sufSum[i] = sufSum[i+1] + arr[i]
	}

	for i := 0; i < n; i++ {
		if preSum[i]*3 != preSum[n-1] {
			continue
		}
		for j := i + 2; j < n; j++ {
			if preSum[i] == sufSum[j] {
				return true
			}
		}
	}
	return false

	//两边往中间出发，且中间至少留一个
}

//leetcode submit region end(Prohibit modification and deletion)

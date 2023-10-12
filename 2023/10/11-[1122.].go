package leetcode

//给你两个数组，arr1 和 arr2，arr2 中的元素各不相同，arr2 中的每个元素都出现在 arr1 中。
//
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末
//尾。
//
//
//
// 示例 1：
//
//
//输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//输出：[2,2,2,1,4,3,3,9,6,7,19]
//
//
// 示例 2:
//
//
//输入：arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
//输出：[22,28,8,6,17,44]
//
//
//
//
// 提示：
//
//
// 1 <= arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// arr2 中的元素 arr2[i] 各不相同
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中
//
//
// Related Topics 数组 哈希表 计数排序 排序 👍 276 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	//rank := map[int]int{}
	//for i, v := range arr2 {
	//	rank[v] = i
	//}
	//sort.Slice(arr1, func(i, j int) bool {
	//	x, y := arr1[i], arr1[j]
	//	rankX, hasX := rank[x]
	//	rankY, hasY := rank[y]
	//	if hasX && hasY {
	//		return rankX < rankY
	//	}
	//	if hasX || hasY {
	//		return hasX
	//	}
	//	return x < y
	//})
	//return arr1

	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)

	for _, v := range arr1 {
		frequency[v]++
	}

	ans := make([]int, 0, len(arr1))

	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}

	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

/*
注意到本题中元素的范围为 [0,1000]，这个范围不是很大，我们也可以考虑不基于比较的排序，例如「计数排序」。
具体地，我们使用一个长度为 1001（下标从 0 到 10000）的数组frequency，
记录每一个元素在数组 arr1 中出现的次数。随后我们遍历数组 arr2,
当遍历到元素 x 时，我们将 frequency[x] 个 x 加入答案中，
并将 frequency[x] 清零。当遍历结束后，所有在 arr2 中出现过的元素就已经有序了。
此时还剩下没有在 arr2 中出现过的元素，
因此我们还需要对整个数组 frequency 进行一次遍历。
当遍历到元素 x 时，如果 frequency[x] 不为 0，我们就将 frequency[x] 个 x 加入答案中。

细节
我们可以对空间复杂度进行一个小优化。实际上，我们不需要使用长度为 1001 的数组，
而是可以找出数组 arr1中的最大值 upper，
使用长度为 upper+1 的数组即可。

*/

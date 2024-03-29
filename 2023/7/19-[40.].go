package leetcode

import (
	"sort"
)

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
// 示例 2:
//
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
//
// Related Topics 数组 回溯 👍 1390 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	length := len(candidates)
	res, path := make([][]int, 0), make([]int, 0, length)
	sort.Ints(candidates) //排序，为剪枝做准备
	var dfs func(start, target int)

	dfs = func(start, target int) {
		if target == 0 { //达到目标值,path符合题意
			//但不能把path直接追加，因为后续还要用，不然底层会修改
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		// 1 1 2 2 5 6 7 10
		for i := start; i < length; i++ {
			if candidates[i] > target { //剪枝 提前返回
				//不然target变负数了，不符合题意
				return
			}

			if i != start && candidates[i] == candidates[i-1] {
				continue //去重
			}
			path = append(path, candidates[i])
			dfs(i+1, target-candidates[i])
			path = path[:len(path)-1]

		}

	}

	dfs(0, target)
	return res

}

//leetcode submit region end(Prohibit modification and deletion)

//func min1(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

//sort.Ints(candidates)
//var ans [][]int
//var freq [][2]int
//for _, num := range candidates {
//if freq == nil || num != freq[len(freq)-1][0] {
//freq = append(freq, [2]int{num, 1})
//} else {
//freq[len(freq)-1][1]++
//}
//}

//var sequence []int
//var dfs func(pos, rest int)
//dfs = func(pos, rest int) {
//	if rest == 0 {
//		ans = append(ans, append([]int(nil), sequence...))
//		return
//	}
//	if pos == len(freq) || rest < freq[pos][0] {
//		return
//	}
//
//	dfs(pos+1, rest)
//
//	//most := min1(rest/freq[pos][0], freq[pos][1])
//	var most int
//	if rest/freq[pos][0]<freq[pos][1]{
//		most=rest/freq[pos][0]
//	}else {
//		most=freq[pos][1]
//	}
//	for i := 1; i <= most; i++ {
//		sequence = append(sequence, freq[pos][0])
//		dfs(pos+1, rest-i*freq[pos][0])
//	}
//	sequence = sequence[:len(sequence)-most]
//}
//dfs(0, target)
//return ans

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/combination-sum-ii/solution/zu-he-zong-he-ii-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

package leetcode

//给你由 n 个小写字母字符串组成的数组 strs，其中每个字符串长度相等。
//
// 这些字符串可以每个一行，排成一个网格。例如，strs = ["abc", "bce", "cae"] 可以排列为：
//
//
//abc
//bce
//cae
//
// 你需要找出并删除 不是按字典序升序排列的 列。在上面的例子（下标从 0 开始）中，列 0（'a', 'b', 'c'）和列 2（'c', 'e',
//'e'）都是按升序排列的，而列 1（'b', 'c', 'a'）不是，所以要删除列 1 。
//
// 返回你需要删除的列数。
//
//
//
// 示例 1：
//
//
//输入：strs = ["cba","daf","ghi"]
//输出：1
//解释：网格示意如下：
//  cba
//  daf
//  ghi
//列 0 和列 2 按升序排列，但列 1 不是，所以只需要删除列 1 。
//
//
// 示例 2：
//
//
//输入：strs = ["a","b"]
//输出：0
//解释：网格示意如下：
//  a
//  b
//只有列 0 这一列，且已经按升序排列，所以不用删除任何列。
//
//
// 示例 3：
//
//
//输入：strs = ["zyx","wvu","tsr"]
//输出：3
//解释：网格示意如下：
//  zyx
//  wvu
//  tsr
//所有 3 列都是非升序排列的，所以都要删除。
//
//
//
//
// 提示：
//
//
// n == strs.length
// 1 <= n <= 100
// 1 <= strs[i].length <= 1000
// strs[i] 由小写英文字母组成
//
//
// Related Topics 数组 字符串 👍 105 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minDeletionSize(strs []string) int {
	var sum int
	m := make(map[int]string)
	for _, v := range strs {
		for i := 0; i < len(v); i++ {
			m[i] = m[i] + string(v[i])
		}
	}
	for _, v := range m {
		if !isIncreasing(v) {
			sum++
		}
	}
	return sum

}
func isIncreasing(str string) bool {
	if len(str) == 1 {
		return true
	}
	for i := 1; i < len(str); i++ {
		if str[i] < str[i-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

/*
func minDeletionSize(strs []string) (ans int) {
    for j := range strs[0] {
        for i := 1; i < len(strs); i++ {
            if strs[i-1][j] > strs[i][j] {
                ans++
                break
            }
        }
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/delete-columns-to-make-sorted/solutions/1483008/shan-lie-zao-xu-by-leetcode-solution-bqyy/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
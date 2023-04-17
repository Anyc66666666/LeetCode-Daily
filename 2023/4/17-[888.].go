package leetcode

//爱丽丝和鲍勃拥有不同总数量的糖果。给你两个数组 aliceSizes 和 bobSizes ，aliceSizes[i] 是爱丽丝拥有的第 i 盒糖果中的糖
//果数量，bobSizes[j] 是鲍勃拥有的第 j 盒糖果中的糖果数量。
//
// 两人想要互相交换一盒糖果，这样在交换之后，他们就可以拥有相同总数量的糖果。一个人拥有的糖果总数量是他们每盒糖果数量的总和。
//
// 返回一个整数数组 answer，其中 answer[0] 是爱丽丝必须交换的糖果盒中的糖果的数目，answer[1] 是鲍勃必须交换的糖果盒中的糖果的数目
//。如果存在多个答案，你可以返回其中 任何一个 。题目测试用例保证存在与输入对应的答案。
//
//
//
// 示例 1：
//
//
//输入：aliceSizes = [1,1], bobSizes = [2,2]
//输出：[1,2]
//
//
// 示例 2：
//
//
//输入：aliceSizes = [1,2], bobSizes = [2,3]
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：aliceSizes = [2], bobSizes = [1,3]
//输出：[2,3]
//
//
// 示例 4：
//
//
//输入：aliceSizes = [1,2,5], bobSizes = [2,4]
//输出：[5,4]
//
//
//
//
// 提示：
//
//
// 1 <= aliceSizes.length, bobSizes.length <= 10⁴
// 1 <= aliceSizes[i], bobSizes[j] <= 10⁵
// 爱丽丝和鲍勃的糖果总数量不同。
// 题目数据保证对于给定的输入至少存在一个有效答案。
//
//
// Related Topics 数组 哈希表 二分查找 排序 👍 221 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var sum1, sum2 int
	var a = []int{0, 0}
	for _, v := range aliceSizes {
		sum1 = sum1 + v
	}
	for _, v := range bobSizes {
		sum2 = sum2 + v
	}
	for _, v := range aliceSizes {
		for _, k := range bobSizes {
			if sum1-v+k == sum2-k+v {
				a[0] = v
				a[1] = k
				return a
			}
		}
	}
	return a

}

//leetcode submit region end(Prohibit modification and deletion)

/*

记爱丽丝的糖果棒的总大小为sumA,鲍勃的糖果棒的总大小为sumB。设答案为{x,y}， 即
爱丽丝的大小为x的糖果棒与鲍勃的大小为y的糖果棒交换，则有如下等式:
sumA-x+y=sumB+x-y
化简，得:

x=y+(sumA-sumB)/2
2
即对于bobSizes 中的任意一个数 y'，只要aliceSizes 中存在-个数x',满足x'=y'+(sumA-sumB)/2

那么{x',y'} 即为一组可行解。
2
为了快速查询aliceSizes 中是否存在某个数，我们可以先将aliceSizes中的数字存入哈希表
中。然后遍历bobSizes 序列中的数y'，在哈希表中查询是否有对应的x'。



func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sumA := 0
    setA := map[int]struct{}{}
    for _, v := range aliceSizes {
        sumA += v
        setA[v] = struct{}{}
    }
    sumB := 0
    for _, v := range bobSizes {
        sumB += v
    }
    delta := (sumA - sumB) / 2
    for i := 0; ; i++ {
        y := bobSizes[i]
        x := y + delta
        if _, has := setA[x]; has {
            return []int{x, y}
        }
    }
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/fair-candy-swap/solutions/585529/gong-ping-de-tang-guo-jiao-huan-by-leetc-tlam/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

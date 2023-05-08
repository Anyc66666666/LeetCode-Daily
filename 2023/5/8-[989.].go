package leetcode

import "strconv"
import "fmt"

//整数的 数组形式 num 是按照从左到右的顺序表示其数字的数组。
//
//
// 例如，对于 num = 1321 ，数组形式是 [1,3,2,1] 。
//
//
// 给定 num ，整数的 数组形式 ，和整数 k ，返回 整数 num + k 的 数组形式 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：num = [1,2,0,0], k = 34
//输出：[1,2,3,4]
//解释：1200 + 34 = 1234
//
//
// 示例 2：
//
//
//输入：num = [2,7,4], k = 181
//输出：[4,5,5]
//解释：274 + 181 = 455
//
//
// 示例 3：
//
//
//输入：num = [2,1,5], k = 806
//输出：[1,0,2,1]
//解释：215 + 806 = 1021
//
//
//
//
// 提示：
//
//
// 1 <= num.length <= 10⁴
// 0 <= num[i] <= 9
// num 不包含任何前导零，除了零本身
// 1 <= k <= 10⁴
//
//
// Related Topics 数组 数学 👍 231 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func addToArrayForm(num []int, k int) []int {
	strK := fmt.Sprint(k)
	lensK := len(strK)
	lensN := len(num)
	if lensN < lensK {
		for i := 0; i < lensK-lensN; i++ {
			zero := []int{0}
			num = append(zero, num[:]...)
		}
	} else {
		for i := 0; i < lensN-lensK; i++ {
			strK = "0" + strK
		}
	}
	fmt.Println(num, strK)

	for i := len(strK) - 1; i >= 0; i-- {
		p, _ := strconv.Atoi(string(strK[i]))
		sum := num[i] + p
		if sum >= 10 {
			num[i] = sum % 10
			if i == 0 {
				one := []int{1}
				num = append(one, num...)
			} else {
				num[i-1] = num[i-1] + 1
			}
			continue
		}
		num[i] = sum

	}
	return num

}

//leetcode submit region end(Prohibit modification and deletion)
/*

func addToArrayForm(num []int, k int) (ans []int) {
    for i := len(num) - 1; i >= 0; i-- {
        sum := num[i] + k%10
        k /= 10
        if sum >= 10 {
            k++
            sum -= 10
        }
        ans = append(ans, sum)
    }
    for ; k > 0; k /= 10 {
        ans = append(ans, k%10)
    }
    reverse(ans)
    return
}

func reverse(num []int) {
    for i, n := 0, len(num); i < n/2; i++ {
        num[i], num[n-1-i] = num[n-1-i], num[i]
    }
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/add-to-array-form-of-integer/solutions/570434/shu-zu-xing-shi-de-zheng-shu-jia-fa-by-l-jljp/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

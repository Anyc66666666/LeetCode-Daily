package leetcode

import "math"

//给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
//
// 整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
//
// 返回被除数 dividend 除以除数 divisor 得到的 商 。
//
// 注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2³¹, 231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 2
//31 − 1 ；如果商 严格小于 -2³¹ ，则返回 -2³¹ 。
//
//
//
// 示例 1:
//
//
//输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
//
// 示例 2:
//
//
//输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。
//
//
//
// 提示：
//
//
// -2³¹ <= dividend, divisor <= 2³¹ - 1
// divisor != 0
//
//
// Related Topics 位运算 数学 👍 1122 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { //被除数最小
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}

	if divisor == math.MinInt32 { //除数最小
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	if dividend == 0 { //被除数为0，则恒为0
		return 0
	}

	//二分 //将所有的正数取相反数
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}
	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 {
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans

}

// 快速乘 -- 乘法分配律 a*11=a* 1011  =a*(1*2零次方+1*2一次方+0*2二次方+1*2三次方)
// x,y 负数 ，z正数
// 判断z*y >= x 是否成立
// y=a z=11  x=?
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 {
		if z&1 > 0 { //二进制最右边的数为1      //如果是0的话，积也是0了，没必要算进去再加
			//需要保证result+add>=x
			if result+add < x { //因为z是正整数,y是负数，所以z*y<=y  >=x
				return false //最大值都比x小，那肯定全小于，返回false
			}
			result += add //最开始是a 2的0   //不断累加，如果是积0就不加
		}
		if z != 1 { //二进制最右边的数为0且z不为1
			//需要保证add+add>=x
			if add+add < x { //add会不断累加，如果还没加完就已经小了，直接false
				return false
			}
			add += add //这里相当于不断乘以2 //2的0 2的1 2的2 2的3
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
//result 0 a   3a  3a  11a
//add    a 2a  4a  8a

package leetcode

//给你两个整数 left 和 right ，在闭区间 [left, right] 范围内，统计并返回 计算置位位数为质数 的整数个数。
//
// 计算置位位数 就是二进制表示中 1 的个数。
//
//
// 例如， 21 的二进制表示 10101 有 3 个计算置位。
//
//
//
//
// 示例 1：
//
//
//输入：left = 6, right = 10
//输出：4
//解释：
//6 -> 110 (2 个计算置位，2 是质数)
//7 -> 111 (3 个计算置位，3 是质数)
//9 -> 1001 (2 个计算置位，2 是质数)
//10-> 1010 (2 个计算置位，2 是质数)
//共计 4 个计算置位为质数的数字。
//
//
// 示例 2：
//
//
//输入：left = 10, right = 15
//输出：5
//解释：
//10 -> 1010 (2 个计算置位, 2 是质数)
//11 -> 1011 (3 个计算置位, 3 是质数)
//12 -> 1100 (2 个计算置位, 2 是质数)
//13 -> 1101 (3 个计算置位, 3 是质数)
//14 -> 1110 (3 个计算置位, 3 是质数)
//15 -> 1111 (4 个计算置位, 4 不是质数)
//共计 5 个计算置位为质数的数字。
//
//
//
//
// 提示：
//
//
// 1 <= left <= right <= 10⁶
// 0 <= right - left <= 10⁴
//
//
// Related Topics 位运算 数学 👍 141 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimeSetBits(left int, right int) int {
	var m int
	for i := left; i <= right; i++ {
		if isPrime(toBinary(i)) {
			m++
		}
	}
	return m

}

func isPrime(a int) bool {
	var m int
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			m++
		}
	}
	//fmt.Println("isPrime:",a,m==2)
	return m == 2
}
func toBinary(a int) int {
	//5
	var d int
	for a != 0 {
		if a%2 == 1 {
			d++
			a = a / 2
			continue
		}
		a = a / 2
	}
	//fmt.Println("count 1:",d)
	return d
}

//leetcode submit region end(Prohibit modification and deletion)
/*

func isPrime(x int) bool {
    if x < 2 {
        return false
    }
    for i := 2; i*i <= x; i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}

func countPrimeSetBits(left, right int) (ans int) {
    for x := left; x <= right; x++ {
        if isPrime(bits.OnesCount(uint(x))) {
            ans++
        }
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/prime-number-of-set-bits-in-binary-representation/solutions/1389365/er-jin-zhi-biao-shi-zhong-zhi-shu-ge-ji-jy35g/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/*

注意到right≤ 106 < 220，因此二-进制中1的个数不会超过19，而不超过19的质数只有
2,3,5, 7,11, 13,17,19
我们可以用一个二进制数mask = 665772 = 101000101000101011002来存储这些质数，其q
mask二进制的从低到高的第i位为1表示i是质数，为0表示i不是质数。
设整数x
的二进制中1的个数为c,若mask按位与2C不为0，则说明c是一个质数。

func countPrimeSetBits(left, right int) (ans int) {
    for x := left; x <= right; x++ {
        if 1<<bits.OnesCount(uint(x))&665772 != 0 {
            ans++
        }
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/prime-number-of-set-bits-in-binary-representation/solutions/1389365/er-jin-zhi-biao-shi-zhong-zhi-shu-ge-ji-jy35g/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

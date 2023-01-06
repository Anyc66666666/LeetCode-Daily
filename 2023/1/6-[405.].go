package leetcode

import "strings"

//给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。
//
// 注意:
//
//
// 十六进制中所有字母(a-f)都必须是小写。
// 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
// 给定的数确保在32位有符号整数范围内。
// 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
//
//
// 示例 1：
//
//
//输入:
//26
//
//输出:
//"1a"
//
//
// 示例 2：
//
//
//输入:
//-1
//
//输出:
//"ffffffff"
//
//
// Related Topics 位运算 数学 👍 260 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func toHex(num int) string {
	//10 -> 16  转化为 10 -> 2 -> 16

	/*给定的数确保在32位有符号整数范围内
	一位十六进制 对应 四位二进制数
	32位相当于8组四位二进制数，四位一组
	*/
	if num == 0 {
		return "0"
	}
	sb := strings.Builder{}

	for i := 7; i >= 0; i-- { //32位相当于8组四位二进制数，四位一组  分8组
		val := num >> (4 * i) & 0xf // >>右移 右边移出的弃掉，左边补0 // 4*7先拿到最左边到四位， // 0xf的二进制是1111(8+4+2+1)  &0xf相当于限制长度为4位
		//跟 1做与运算 仍然是等于原来本身 不变
		if val > 0 || sb.Len() > 0 {
			var digit byte
			if val < 10 { //此时val<10
				digit = '0' + byte(val) //0-9
			} else { //此时val>10
				digit = 'a' + byte(val-10) //a-f
			}

			sb.WriteByte(digit)
		}

	}

	return sb.String()

}

//leetcode submit region end(Prohibit modification and deletion)

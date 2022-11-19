package leetcode

import "fmt"

//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
//
// 例如：
//
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//
//
// 示例 1：
//
//
//输入：columnNumber = 1
//输出："A"
//
//
// 示例 2：
//
//
//输入：columnNumber = 28
//输出："AB"
//
//
// 示例 3：
//
//
//输入：columnNumber = 701
//输出："ZY"
//
//
// 示例 4：
//
//
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
//
//
//
//
// 提示：
//
//
// 1 <= columnNumber <= 2³¹ - 1
//
//
// Related Topics 数学 字符串 👍 580 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	//A1 Z26 AA26+1=27  AZ 52 BA 53 26*2+1  ZA26*26+1  ZZ 26*26+26
	//AAA 26*26+26+1
	var m = map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}

	var d = columnNumber
	var n int
	var c []int
	for {
		n = d % 26
		d = d / 26
		if n == 0 {
			n = 26
			d--
		}
		if d == 0 {
			c = append(c, n)
			break
		}
		c = append(c, n)
	}
	fmt.Println(c)
	var s string
	for i := 0; i < len(c); i++ {
		s = m[c[i]] + s
	}

	return s

}

//leetcode submit region end(Prohibit modification and deletion)

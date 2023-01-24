package leetcode

import "fmt"

//在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄。他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。
//
// 当提莫攻击艾希，艾希的中毒状态正好持续 duration 秒。
//
// 正式地讲，提莫在 t 发起发起攻击意味着艾希在时间区间 [t, t + duration - 1]（含 t 和 t + duration - 1）处于中毒
//状态。如果提莫在中毒影响结束 前 再次攻击，中毒状态计时器将会 重置 ，在新的攻击之后，中毒影响将会在 duration 秒后结束。
//
// 给你一个 非递减 的整数数组 timeSeries ，其中 timeSeries[i] 表示提莫在 timeSeries[i] 秒时对艾希发起攻击，以及一
//个表示中毒持续时间的整数 duration 。
//
// 返回艾希处于中毒状态的 总 秒数。
//
// 示例 1：
//
//
//输入：timeSeries = [1,4], duration = 2
//输出：4
//解释：提莫攻击对艾希的影响如下：
//- 第 1 秒，提莫攻击艾希并使其立即中毒。中毒状态会维持 2 秒，即第 1 秒和第 2 秒。
//- 第 4 秒，提莫再次攻击艾希，艾希中毒状态又持续 2 秒，即第 4 秒和第 5 秒。
//艾希在第 1、2、4、5 秒处于中毒状态，所以总中毒秒数是 4 。
//
// 示例 2：
//
//
//输入：timeSeries = [1,2], duration = 2
//输出：3
//解释：提莫攻击对艾希的影响如下：
//- 第 1 秒，提莫攻击艾希并使其立即中毒。中毒状态会维持 2 秒，即第 1 秒和第 2 秒。
//- 第 2 秒，提莫再次攻击艾希，并重置中毒计时器，艾希中毒状态需要持续 2 秒，即第 2 秒和第 3 秒。
//艾希在第 1、2、3 秒处于中毒状态，所以总中毒秒数是 3 。
//
//
//
//
// 提示：
//
//
// 1 <= timeSeries.length <= 10⁴
// 0 <= timeSeries[i], duration <= 10⁷
// timeSeries 按 非递减 顺序排列
//
//
// Related Topics 数组 模拟 👍 354 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findPoisonedDuration(timeSeries []int, duration int) int {

	// 1 3 4 5    2   1-2 3-4 4-5 5-6 1-6
	// 1 3 4 5    3   1-3 3-5 4-6 5-7  1-7
	// 1 3 4 5    1   1   3  4  5
	// 1-7 9-10  12-13
	// 1 3 4 5 6 8 11 15 3  1-3 3-5 4-6 6-8 8-10 11-13 15-17	  1-13 15-17
	var a []int
	var n int
	a = append(a, timeSeries[0], timeSeries[0])
	for i := 0; i < len(timeSeries); i++ {

		if timeSeries[i] <= a[1]+1 {
			a[1] = timeSeries[i] + duration - 1
			continue
		}
		fmt.Println(a)
		n = n + a[len(a)-1] - a[0] + 1
		a[0] = timeSeries[i]
		a[1] = timeSeries[i] + duration - 1

	}
	n = n + a[len(a)-1] - a[0] + 1

	return n

}

//leetcode submit region end(Prohibit modification and deletion)

/*
我们只需要对数组进行一次扫描就可以计算出总的中毒持续时间。我们记录艾希恢复为未中毒的起始时间 expired，设艾希遭遇第 i 次的攻击的时间为 timeSeries[i]。当艾希遭遇第 i 攻击时：

如果当前他正处于未中毒状态，则此时他的中毒持续时间应增加 duration，同时更新本次中毒结束时间 expired等于 timeSeries[i]+duration ；
如果当前他正处于中毒状态，由于中毒状态不可叠加，我们知道上次中毒后结束时间为 expired，本次中毒后结束时间为 timeSeries[i]+duration，因此本次中毒增加的持续中毒时间为 timeSeries[i]+duration−expired；
我们将每次中毒后增加的持续中毒时间相加即为总的持续中毒时间。


func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
    expired := 0
    for _, t := range timeSeries {
        if t >= expired {
            ans += duration
        } else {
            ans += t + duration - expired
        }
        expired = t + duration
    }
    return
}

*/

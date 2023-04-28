package leetcode

//写一个 RecentCounter 类来计算特定时间范围内最近的请求。
//
// 请你实现 RecentCounter 类：
//
//
// RecentCounter() 初始化计数器，请求数为 0 。
// int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求
//）。确切地说，返回在 [t-3000, t] 内发生的请求数。
//
//
// 保证 每次对 ping 的调用都使用比之前更大的 t 值。
//
//
//
// 示例 1：
//
//
//输入：
//["RecentCounter", "ping", "ping", "ping", "ping"]
//[[], [1], [100], [3001], [3002]]
//输出：
//[null, 1, 2, 3, 3]
//
//解释：
//RecentCounter recentCounter = new RecentCounter();
//recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
//recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
//recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
//recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是 [2,3002]，返回
//3
//
//
//
//
// 提示：
//
//
// 1 <= t <= 10⁹
// 保证每次对 ping 调用所使用的 t 值都 严格递增
// 至多调用 ping 方法 10⁴ 次
//
//
// Related Topics 设计 队列 数据流 👍 214 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type RecentCounter struct {
	T []int
}

func Constructor() RecentCounter {
	var t []int
	return RecentCounter{T: t}
}

func (this *RecentCounter) Ping(t int) int {
	this.T = append(this.T, t)
	var a int
	for _, v := range this.T {
		if v < t-3000 {
			a++
		}
	}
	return len(this.T) - a

}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
//leetcode submit region end(Prohibit modification and deletion)

/*

type RecentCounter []int

func Constructor() (_ RecentCounter) { return }

func (q *RecentCounter) Ping(t int) int {
    *q = append(*q, t)
    for (*q)[0] < t-3000 {
        *q = (*q)[1:]
    }
    return len(*q)
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/number-of-recent-calls/solutions/1467662/zui-jin-de-qing-qiu-ci-shu-by-leetcode-s-ncm1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

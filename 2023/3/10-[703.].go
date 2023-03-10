package leetcode

//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
// 请实现 KthLargest 类：
//
//
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
//
//
//
//
// 示例：
//
//
//输入：
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//输出：
//[null, 4, 5, 5, 8, 8]
//
//解释：
//KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//kthLargest.add(3);   // return 4
//kthLargest.add(5);   // return 5
//kthLargest.add(10);  // return 5
//kthLargest.add(9);   // return 8
//kthLargest.add(4);   // return 8
//
//
//
//提示：
//
//
// 1 <= k <= 10⁴
// 0 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// -10⁴ <= val <= 10⁴
// 最多调用 add 方法 10⁴ 次
// 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
//
//
// Related Topics 树 设计 二叉搜索树 二叉树 数据流 堆（优先队列） 👍 414 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type KthLargest struct {
	pq      []int
	size    int
	maxSize int
}

func Constructor(k int, nums []int) KthLargest {
	heap := KthLargest{
		size:    0,
		maxSize: k,
		pq:      make([]int, k+1),
	}
	for _, num := range nums {
		heap.Add(num)
	}
	return heap

}

// 1 2 4 5 6 7

func (this *KthLargest) Add(val int) int {
	if this.size < this.maxSize {
		this.size++
		this.pq[this.size] = val
		this.Swim(this.size) //上浮
	} else {
		if this.pq[1] >= val {
			return this.pq[1]
		} else {
			this.pq[1] = val
			this.Sink(1) //下沉
		}
	}
	return this.pq[1]

}

/* 上浮第 k 个元素，以维护最小堆性质 */
func (this *KthLargest) Swim(k int) {
	for k > 1 && this.Less(k, this.Parent(k)) {
		this.Exch(k, this.Parent(k))
		k = this.Parent(k)
	}
}

/* 下沉第 k 个元素，以维护最小堆性质 */
func (this *KthLargest) Sink(k int) {
	for this.Left(k) <= this.maxSize {
		older := this.Left(k)
		if this.Right(k) <= this.maxSize && this.Less(this.Right(k), older) {
			older = this.Right(k)
		}
		if this.Less(k, older) {
			break
		}
		this.Exch(older, k)
		k = older
	}
}

/* 交换数组的两个元素 */
func (this *KthLargest) Exch(i, j int) {
	tmp := this.pq[i]
	this.pq[i] = this.pq[j]
	this.pq[j] = tmp
}

/* pq[i] 是否比 pq[j] 小？ */
func (this *KthLargest) Less(i, j int) bool {
	return this.pq[i] < this.pq[j]
}

// 父节点的索引
func (this *KthLargest) Parent(root int) int {
	return root / 2
}

// 左孩子的索引
func (this *KthLargest) Left(root int) int {
	return root * 2
}

// 右孩子的索引
func (this *KthLargest) Right(root int) int {
	return root*2 + 1
}

//作者：刘新
//链接：https://leetcode.cn/problems/kth-largest-element-in-a-stream/solutions/38494/golandtong-guo-zui-xiao-dui-shi-xian-by-liuxin95/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)

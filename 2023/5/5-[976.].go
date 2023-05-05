package leetcode

import "sort"

//给定由一些正数（代表长度）组成的数组 nums ，返回 由其中三个长度组成的、面积不为零的三角形的最大周长 。如果不能形成任何面积不为零的三角形，返回 0。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [2,1,2]
//输出：5
//解释：你可以用三个边长组成一个三角形:1 2 2。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,1,10]
//输出：0
//解释：
//你不能用边长 1,1,2 来组成三角形。
//不能用边长 1,1,10 来构成三角形。
//不能用边长 1、2 和 10 来构成三角形。
//因为我们不能用任何三条边长来构成一个非零面积的三角形，所以我们返回 0。
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10⁴
// 1 <= nums[i] <= 10⁶
//
//
// Related Topics 贪心 数组 数学 排序 👍 240 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(nums []int) int {
	//// 1 2 3 4 5  死解 太慢了
	//	sort.Ints(nums)
	//	var perimeter int
	//	for i:=len(nums)-1;i>1;i--{
	//		for j:=i-1;j>0;j--{
	//			for k:=j-1;k>=0;k--{
	//				if nums[i]+nums[j]<=nums[k]||
	//					nums[i]+nums[k]<=nums[j]||
	//					nums[k]+nums[j]<=nums[i]{
	//					continue
	//				}
	//				if nums[i]+nums[j]+nums[k]>perimeter{
	//					perimeter=nums[i]+nums[j]+nums[k]
	//					return perimeter
	//				}
	//			}
	//		}
	//	}
	//	return perimeter

	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {

		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i-1] + nums[i-2] + nums[i]
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
/*
方法--:贪心+排序
不失一般性，我们假设三角形的边长a,b,c满足a≤b≤c，那么这三条边组成面积不为零的
三角形的充分必要条件为a+b> c。
基于此，我们可以选择枚举三角形的最长边c，而从贪心的角度考虑，我们一定是选「小于c
的最大的两个数」作为边长a和b，此时最有可能满足a+b> c，使得三条边能够组成一个
三角形，且此时的三角形的周长是最大的。
因此，我们先对整个数组排序，倒序枚举第i个数作为最长边，那么我们只要看其前两个数
A[i-2]和A[i-1],判断A[i-2]+A[i- 1]是否大于A[i] 即可，如果能组成三角形我们就
找到了最大周长的三角形，返回答案A[i- 2]+ A[i-1]+ A[i]即可。如果对于任何数作为最
长边都不存在面积不为零的三角形，则返回答案0。


为什么只需要判断a+b>c ?
正常情况下，对于三条边长，判断能否组成三角形需要判断任何两条边长相加都大于其余的一条边长，即:
a+b>c&&a+C>b&&b+c>a
而如果已知a<=b<=c,那么必然有:
1.a+c>b，因为c>=b,那c加上一个正数一定就比b大了。而题目里说所数都>=1,所以
c加上a-定比b大。
2. b+c>a,因为b和c至少跟a-样大(b>=a, c>=a)， 加起来的结果至少有2a,即b+c>=
2a> a
所以最终只需要判断a+b> c即可。
为什么只需要判断数组中相邻的三个数?
在固定最后一个数A[i]时，前两个数需不需要再往前找呢?
如果A[i-2]+ A[i-1] <=A[i]，这三个数一定不能构成三角形，而A[i-3]以及 更往前的数，都小于等于A[i-
2],所以再往前取任何两个数只会让相加的值更小，就更不能满足A[j] + A[k] > A[]了(i<i-2, k<i-1, j<k)。
所以如果相邻的数构不成三角形，就不需要再固定第三个数并往前找两个数了。
如果A[i-2] + A[i-1] > A[i]j,这三个数可以构成三角形，再往前找只会让周长变短，所以也不用再往前了。
综上，只需要判断相邻的三个数。
*/

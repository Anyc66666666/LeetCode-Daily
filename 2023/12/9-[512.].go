if len(nums) == 0 {
return nil
}
best := 0
for i, num := range nums {
if num > nums[best] {
best = i
}
}
return &TreeNode{nums[best], constructMaximumBinaryTree(nums[:best]), constructMaximumBinaryTree(nums[best+1:])}

作者：力扣官方题解
链接：https://leetcode.cn/problems/maximum-binary-tree/solutions/1759348/zui-da-er-cha-shu-by-leetcode-solution-lbeo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
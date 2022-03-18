package hot100

// 关键: 有序 代表着双指针
// 推荐解法: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/solution/gong-shui-san-xie-guan-yu-shan-chu-you-x-glnq/
// 关键:
func removeDuplicates2(nums []int) int {
	var process func(k int) int
	process = func(k int) int {
		ret := 0
		for _, v := range nums {
			// ret<k 代表着,先收集前面k个元素
			// nums[ret-k]!=v 代表着 只有当 当前位置和前面k个元素的下标的位置不同的时候才能保留
			// 如 k=2 ,nums=1,1,1,1,2 ,则 ret<k的时候会先将前面2个1 填充, 然后第三个元素的时候,因为nums[2-2]=1 ,和当前的1匹配,则不能赋值
			if ret < k || nums[ret-k] != v {
				nums[ret] = v
				ret++
			}
		}
		return ret
	}
	return process(2)
}

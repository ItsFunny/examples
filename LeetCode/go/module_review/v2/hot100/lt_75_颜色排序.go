package hot100

// 关键:
// 1. 2个for循环
// 2. 用map 统计个数
func sortColors(nums []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v]+1
	}
	index:=0
	for i:=0;i<3;i++{
		count:=m[i]
		for j:=0;j<count;j++{
			nums[index]=i
			index++
		}
	}
}

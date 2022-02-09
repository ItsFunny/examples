package hot100

// 题目的意思是,将同为`字母异位词`的放在一起返回
// 字母异位词是字母出现的元素都相同,顺序不同而已
// 如 abc 与 acb 是同一组异位词,此时可以发现,abc,acb出现的字符的个数都是相同的,因此对于这样的数,可以用map来做映射
func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	m := make(map[[26]int][]string)
	for _, str := range strs {
		// 计算这个str 中出现的字符的次数,异位词必然都是相同的
		count := [26]int{}
		for _, v := range str {
			count[v-'a']++
		}
		m[count] = append(m[count], str)
	}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

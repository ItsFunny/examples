package hot100

import "sort"

type arrSorts [][]int

func (a arrSorts) Len() int {
	return len(a)
}

func (a arrSorts) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a arrSorts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 关键:
// 1. 排序
// 2. for 循环,如果该index的数组的最后一个值比result中的最后一个的数组的最后一个小,则合并
// 3. 合并的时候还有一个注意点,就是必须判断合并的最大值
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return nil
	}
	sort.Sort(arrSorts(intervals))
	ret := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		node := intervals[i]
		// 初始化
		if len(ret)==0{
			ret=append(ret,node)
			continue
		}
		last:=ret[len(ret)-1]
		// 判断初始值与last中的最后值
		if node[0]<=last[len(last)-1]{
			// 表明可以合并,此时还需要判断谁的值更大
			if last[len(last)-1]<node[len(node)-1]{
				last[len(last)-1]=node[len(node)-1]
			}
		}else{
			// 直接append
			ret=append(ret,node)
		}
	}

	return ret
}

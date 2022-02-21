package hot100

// 关键: 判断什么时候需要合并: 当只有有交叉的时候需要合并
// 1. 当interval 更小的时候,直接append即可
// 1. 遍历原始区间, 判断是否在区间内 : 当left< val0 || right>val1 的时候 不在区间内
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	ret := make([][]int, 0)

	merged := false
	for _, interval := range intervals {
		val0, val1 := interval[0], interval[1]
		// 先if 不在区间的情况
		if val0 > right {
			// 说明当前interval 在右侧 ,并且是没有交集
			if !merged {
				ret = append(ret, []int{left, right})
				merged = true
			}
			ret = append(ret, interval)
		} else if val1 < left {
			// 说明当前interval在左侧,并且没有交集 (既 这个interval更小)
			// 此时不需要对left,right 进行处理,因为还没有匹配到更大的
			ret = append(ret, interval)
		} else {
			// 此时表明是有重叠空间了
			// 则开始计算并集,并集的计算,左边选更小的,右边选更大的
			left = insertMin(val0, left)
			right = insertMax(val1, right)
			// 当合并了区间之后,就可以拿新的值去继续for循环了
		}
	}

	// 注意: 最后一步,可能这个新的interval,或者是合并区间后的left,right 是最右边(既最大),可能会在for循环中一直没有append
	// 所以最后还需要继续if
	if !merged {
		ret = append(ret, []int{left, right})
	}

	return ret
}
func insertMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func insertMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

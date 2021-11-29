/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/15 9:10 上午
# @File : lt_134_加油站.go
# @Description :
# @Attention :
*/
package offer

// 解题关键:
// 遍历所有的加油站,判断是否都能够到达
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(cost)
	for i := 0; i < n; {
		count := 0     // 当前走过的加油站
		sumOfCas := 0  // 总共加的油
		sumOfCost := 0 // 总共要使用的油
		for count < n {
			index := (i + count) % n // 因为是环形,所以要取余
			sumOfCas += gas[index]
			sumOfCost += cost[index]
			if sumOfCost > sumOfCas {
				break // 代表的是油不够,所以直接break即可
			}
			count++ // 满足,则行驶到下一个站点
		}
		if count == n {
			// 代表着,这个加油站出发,能够回到原来位置
			return i
		}
		// 则从下一个未探测的站点出发,这一步就是可以避免重复计算的
		i = i + count + 1
	}
	// 都不满足
	return -1
}
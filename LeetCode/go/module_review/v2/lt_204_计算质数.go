/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 8:33 下午
# @File : lt_204_计算质数.go
# @Description :
# @Attention :
*/
package v2

// 质数的定义: 乘法 只有 1和 它自己本身,1 不是质数
// 如果一个数v是质数,那么2v,3v,4v,5v ....nv 必然不是质数
func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	ret := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			ret++
			// 则 2v,3v等必然不是质数
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return ret
}

/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/19 9:29 下午
# @File : lt_166_分数到小数.go
# @Description :
# @Attention :
*/
package offer

import (
	"fmt"
	"strconv"
)

// 参考:  https://leetcode-cn.com/problems/fraction-to-recurring-decimal/solution/gong-shui-san-xie-mo-ni-shu-shi-ji-suan-kq8c4/
// 关键: 用草稿比比划下草稿纸上运算的过程
// 1. 先直接整除,能除下来的,则是实数部分,剩余的则是小数部分
// 2. 然后取余对这个数整除的时候,肯定是都需要*10的
// 3. 额外关注无限小数的这种,如 10/3 =3.33333 ,并且会是一直都是以小数结尾,这时候,对于这种是通过缓存来判断的
// 以 10/3 为例 10/3 不整除,最后会余一个1 去除3 ,1/3 的时候,会不停的*10 然后继续除
func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	ret := make([]byte, 0)
	if numerator*denominator < 0 {
		ret = append(ret, '-')
	}
	if numerator < 0 {
		numerator *= -1
	}
	if denominator < 0 {
		denominator *= -1
	}

	first := numerator / denominator
	left := numerator % denominator
	ret = append(ret, byte(first))
	ret = append(ret, '.')
	lastCache := make(map[int]int)
	for left != 0 {
		lastCache[left] = len(ret)
		left *= 10
		// 计算结果 : 1*10 /3
		ret = append(ret, byte(left/denominator))
		left %= denominator
		if v, exist := lastCache[left]; exist {
			// 表明是重复的,如  1/3 之后  10/3 依旧有了(1.3) ,
			// 则 抽离非重复的+ 重复的
			nonRepeat := ret[:v]
			repeated := append([]byte{'{'}, ret[v:]...)
			repeated = append(repeated, '}')
			return fmt.Sprintf("%s(%s)", string(nonRepeat), string(repeated))
		}
	}

	return string(ret)
}

/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/12 8:55 上午
# @File : lt_43_字符串相乘.go
# @Description :
# @Attention :
*/
package hot100

import "strconv"

/*
// 解题关键:
// 采用加法:
// num1的数 * num2 的每个数, 然后和再累加 (这里唯一需要注意的点是,记得结果跟上0)
// 如 num1=1234 ,num2=456
// 1. 1234 *6
		1234
		   6
		7404
	2. 1234 *5  // 因为5 为倒数第2个数,后面还有个数,所以最后要加上一个0
	   1234
		 5
	  6170+0=> 61700
	3. 1234*4 // 因为4为倒数第3个数,后面有2个数,所以要加上2个0
	   1234
	    4
	 493600
	4. 最终将所有结果相加: 7304+61700+49360 =562704
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)

	ret := ""
	// 遍历num2 的每个数
	for i := l2 - 1; i >= 0; i-- {

		cur := ""
		// 表明在这个位置上该追加几个0 ,
		for j := l2 - 1; j > i; j-- {
			cur += "0"
		}
		// 代表的是乘 多于的值,如 3*4=12 则 more为1
		more := 0
		// 遍历num1的每个数,每个数去乘 num2的每个数
		y := int(num2[i] - '0')
		for j := l1 - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			mux := x*y + more
			// 然后取余 获取得到最后一个数
			cur = strconv.Itoa(mux%10) + cur
			more = mux / 10
		}
		// 此时more 可能不为0 ,所以还需要加上more
		for ; more != 0; more /= 10 {
			cur = strconv.Itoa(more%10) + cur
		}
		// 最后再讲结果相加,字符串相加
		ret = addStrings(cur, ret)
	}
	return ret
}

func addStrings(num1, num2 string) string {
	i1, i2 := len(num1)-1, len(num2)-1
	add := 0
	ret := ""
	for ; i1 >= 0 || i2 >= 0 || add != 0; i1, i2 = i1-1, i2-1 {
		x, y := 0, 0
		if i1 >= 0 {
			x = int(num1[i1] - '0')
		}
		if i2 >= 0 {
			y = int(num2[i2] - '0')
		}
		result := x + y + add
		// 然后计算多的值
		ret = strconv.Itoa(result%10) + ret
		add = result / 10

	}

	return ret
}

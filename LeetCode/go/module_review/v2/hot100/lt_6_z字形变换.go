/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/8 9:18 下午
# @File : lt_6_z字形变换.go
# @Description :
# @Attention :
*/
package hot100

// 关键:
// 找规律,发现这个图形的构建是通过 从上到下,然后从下到上排列的
/*
 */
/*
	将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
 

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"
*/

func convert(s string, numRows int) string {
	if numRows==1{
		return s
	}
	// 如: PAYPALISHIRING
	/*
		P   A   H   N
		A P L S I I G
		Y   I   R
		P,A,Y 到达三个之后 ,上移 然后继续 P ,再继续上移 => A ,发现到了临界点,变成了L
	*/
	ret := ""
	sbs := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		sbs[i] = make([]byte, 0)
	}
	down := false
	currentRow := 0
	for i := 0; i < len(s); i++ {
		sbs[currentRow] = append(sbs[currentRow], s[i])
		if currentRow == 0 || currentRow == numRows-1 {
			down = !down
		}
		if down {
			currentRow++
		} else {
			currentRow--
		}
	}
	for _, v := range sbs {
		ret += string(v)
	}
	return ret
}

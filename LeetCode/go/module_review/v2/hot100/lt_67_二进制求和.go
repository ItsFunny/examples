package hot100

import (
	"strconv"
	"strings"
)

// 关键:
// 类似于两个 数组合并
// 问题: carry的作用
func addBinary(a string, b string) string {
	l1, l2 := len(a)-1, len(b)-1
	carry := 0

	ret := strings.Builder{}
	for l1 >= 0 && l2 >= 0 {
		sum := carry
		sum += int(a[l1] - '0')
		sum += int(b[l2] - '0')
		carry = sum / 2
		ret.WriteString(strconv.Itoa(sum % 2))
		l1--
		l2--
	}
	// 如果l1更长
	for l1 >= 0 {
		sum := carry + int(a[l1]-'0')
		carry = sum / 2
		ret.WriteString(strconv.Itoa(sum % 2))
		l1--
	}

	// 如果l2 更长
	for l2 >= 0 {
		sum := carry + int(b[l2]-'0')
		carry = sum / 2
		ret.WriteString(strconv.Itoa(sum % 2))
		l2--
	}
	// 还有个进位数没加进去，需要补充
	if carry == 1 {
		ret.WriteString("1")
	}
	return addBinaryReverse(ret.String())
}

func addBinaryReverse(str string) string {
	ret := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		ret = append(ret, str[i])
	}
	return string(ret)
}

/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-20 14:29 
# @File : common.go
# @Description : 常规参数校验
*/
package utils

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 调用fft之前: 将音频字节码转换为复数数组
// func ConvT

func GetUUID() string {
	uid:= uuid.NewV4()
	return strings.Replace(uid.String(), "-", "", -1)
}

func GetStringLen(str string) int {
	return len([]rune(str))
}

// 返回以分为单位的值
func GetMoney(str string) (int, error) {
	isNegative := false
	var returnErr = errors.New("输入不合法")
	if strings.HasPrefix(str, "-") {
		isNegative = true
		str = str[1:]
	}

	moneys := strings.Split(str, ".")
	if len(moneys) > 2 {
		return 0, returnErr
	} else if len(moneys) == 2 {
		// 含有小数
		m1, err := strconv.Atoi(moneys[0])
		if err != nil {
			return 0, returnErr
		}
		m2, err := strconv.Atoi(moneys[1])
		if err != nil {
			return 0, returnErr
		}
		if len(moneys[1]) == 1 {
			m2 *= 10
		}

		if len(moneys[1]) > 2 {
			return 0, returnErr
		}

		if isNegative {
			return (m1*100 + m2) * -1, nil
		} else {
			return m1*100 + m2, nil
		}
	} else {
		// 不含小数
		i, err := strconv.Atoi(moneys[0])
		if err != nil {
			return 0, returnErr
		}

		if isNegative {
			return i * 100 * -1, nil
		} else {
			return i * 100, nil
		}
	}
}

// 获取两个时间之间差多少天
func GetTimeSubDays(t1, t2 time.Time) int {
	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()
	if hours <= 0 {
		// return -1
		t1, t2 = t2, t1
		hours = t1.Sub(t2).Hours()
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0
		} else {
			return 1
		}
	} else {
		// equal or more than 24
		if (hours/24)-float64(int(hours/24)) == 0 {
			// just 24's times
			return int(hours / 24)
		} else {
			// more than 24 hours
			return int(hours/24) + 1
		}
	}
}

// 获取两个时间之间差多少小时
func GetTimeSubHours(t1, t2 time.Time) int {
	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()
	if hours <= 0 {
		// return -1
		t1, t2 = t2, t1
		hours = t1.Sub(t2).Hours()
	}

	var hoursInt int
	hoursInt = int(hours)
	hours *= 10
	if int(hours)%10 > 0 {
		hoursInt += 1
	}
	return hoursInt
}

func FormatStruct(info interface{}) string {
	result := ""

	bytes, err := json.MarshalIndent(info, " ", " ")
	if err != nil {
		result = fmt.Sprintf("解析interface错误:%s", err.Error())
		return result
	}

	result = "\n" + string(bytes)
	return result
}

// 获取金钱的大写返回
// 入参以分为单位,出参
func ConvertNumToCny(num int) string {
	strnum := strconv.Itoa(num)
	// 判断正数负数
	if num < 0 {
		strnum = string([]rune(strnum)[1:])
	}
	sliceUnit := []string{"亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum[:] {
		str = str + upperDigitUnit[string(v)] + s[k]
	}
	reg, _ := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, _ = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")

	reg, _ = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, _ = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")

	reg, _ = regexp.Compile(`零{2,}`)
	str = reg.ReplaceAllString(str, "零")

	reg, _ = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")

	reg, _ = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")

	reg, _ = regexp.Compile(`零*元`)
	str = reg.ReplaceAllString(str, "元")

	reg, _ = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^元")

	reg, _ = regexp.Compile(`零元`)
	str = reg.ReplaceAllString(str, "零")

	if num < 0 {
		str = "负" + str
	}
	return str
}

// 转换次数
// 入参是数字,出参 是 一\二\三 ...
func ConvertNumToCny2(num int) string {
	strnum := strconv.Itoa(num)
	str := ""

	upperDigitUnit := map[string]string{"0": "零", "1": "一", "2": "二", "3": "三", "4": "四", "5": "五", "6": "六", "7": "七", "8": "八", "9": "九"}

	sliceUnit := []string{"千", "百", "十", "亿", "千", "百", "十", "万", "千", "百", "十", "亿", "千", "百", "十", "万", "千", "百", "十", ""}
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]

	// 从最高位开始,逐级往后拼接
	for i, c := range strnum[:] {
		temp := string(c)

		// 如果第一位是1 并且单位是十,则不写数字,只写十
		if len(str) == 0 && temp == "1" && s[i] == "十" {
			str += s[i]
			continue
		}

		// 如果碰上为0 的情况,如果str已经有零做结尾,则不需要拼接,直接跳过
		if temp == "0" && strings.HasSuffix(str, upperDigitUnit[temp]) {
			continue
		}

		// 如果碰上为0 的情况,如果str没有零做结尾,则只拼一个零就行了,不用带单位
		if temp == "0" && !strings.HasSuffix(str, upperDigitUnit[temp]) {
			str += upperDigitUnit["0"]
		}

		// 其它情况直接拼接
		str += upperDigitUnit[temp] + s[i]
	}

	// 去掉结尾的所有零
	str = strings.TrimSuffix(str, upperDigitUnit["0"])

	return str
}

// 把 钱(分为单位) 转换为 3位 分解的 字符串
// 如 把  1 分解为 1.00
// 如 把  10 分解为 1.00
// 如 把  100 分解为 1.00
// 12345678900 返回 123,456,789.00
func ConvertMoney(money int) string {
	isNegative := false
	if money < 0 {
		isNegative = true
		if money < -10000 {
			money *= -1
		}
	}

	// 如果小于 10000 ,返回小数
	if money <= 10000 {
		return strconv.FormatFloat(float64(money)/100, 'f', 2, 64)
	}

	subMoneyStr := strconv.Itoa(money % 100)
	for len(subMoneyStr) < 2 {
		subMoneyStr = "0" + subMoneyStr
	}
	subMoneyStr = "." + subMoneyStr

	result := ""
	moneyStr := ReverseStr(strconv.Itoa(money / 100))
	for i := 0; i < GetStringLen(moneyStr); i += 3 {
		j := i + 3
		if j >= GetStringLen(moneyStr) {
			j = GetStringLen(moneyStr)
		}

		result += moneyStr[i:j] + ","
	}
	result = strings.TrimSuffix(result, ",")

	if isNegative {
		return "-" + ReverseStr(result) + subMoneyStr
	} else {
		return ReverseStr(result) + subMoneyStr
	}
}

// 反转字符串
func ReverseStr(str string) string {
	runes := []rune(str)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// []string->[]int数组
func ConvertArrayStringToInt(strs ...string) ([]int, error) {
	var ids []int
	for _, v := range strs {
		// 如果为空串，则默认初值为0
		if len(v) == 0 || strings.TrimSpace(v) == "" {
			v = "0"
		}
		if id, err := strconv.Atoi(v); err == nil {
			ids = append(ids, id)
		} else {
			return ids, fmt.Errorf("[]string转换为[]int出错")
		}
	}
	return ids, nil
}

// []int ->[]string的转换
func ConvertArrayIntToString(nums ...int) ([]string, error) {
	var ids []string
	for _, v := range nums {
		ids = append(ids, strconv.Itoa(v))
	}
	return ids, nil
}

// MM->CM的转换
func MMTOCM(numMM float64) float64 {

	return numMM / 10
}

// 用户去掉打印时的页眉页脚
// Save the File to an xlsx file at the provided path.
func Save(path string, f *xlsx.File) (err error) {
	target, err := os.Create(path)
	if err != nil {
		return err
	}
	err = Write(target, f)
	if err != nil {
		return err
	}
	return target.Close()
}

// Write the File to io.Writer as xlsx
func Write(writer io.Writer, f *xlsx.File) (err error) {
	parts, err := f.MarshallParts()
	vs := parts["xl/worksheets/sheet1.xml"]

	start := `<headerFooter differentFirst="false" differentOddEven="false">`
	end := `</headerFooter>`
	newStr := "<oddHeader></oddHeader><oddFooter></oddFooter>"

	s1 := strings.Index(vs, start)
	e1 := strings.Index(vs, end)

	startStr := string(vs[0 : s1+len(start)])
	endStr := string(vs[e1:])
	result := startStr + newStr + endStr

	parts["xl/worksheets/sheet1.xml"] = result
	if err != nil {
		return
	}
	zipWriter := zip.NewWriter(writer)
	for partName, part := range parts {
		w, err := zipWriter.Create(partName)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(part))
		if err != nil {
			return err
		}
	}
	return zipWriter.Close()
}

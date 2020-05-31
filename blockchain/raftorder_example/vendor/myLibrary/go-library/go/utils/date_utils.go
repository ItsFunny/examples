package utils

import "time"

type ByteTime []uint8

// 将[]uint8 类型的数据转为日期
func (t ByteTime) Parse() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", string(t))
}

func GetCurrentTimeDefault() string {
	return GetCurrentTimeStrByTemplate("2006-01-02 15:04:05")
}

func GetCurrentTimeStrByTemplate(tempalte string) string {
	return time.Now().Format(tempalte)
}

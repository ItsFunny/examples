package utils

import (
	"fmt"
	"strings"
)

func DebugPrintDetail(decorate string, valueName string, value interface{}) {
	fmt.Println(strings.Repeat(decorate, 10))
	fmt.Printf("%s的值为:%v \n \n", valueName, value)
	fmt.Println(strings.Repeat(decorate, 10))
}
func DebugPrint(decorate string, name string) {
	fmt.Println(strings.Repeat(decorate, 10))
	fmt.Printf("进入了:[ %v ]此方法 \n ", name)
	fmt.Println(strings.Repeat(decorate, 10))
}

func DebugPrintSignal(decorate string, signale interface{}) {
	DebugDecorateShowSignal(decorate, 10, signale)
}

func DebugDecorateShowSignal(decorate string, len int, signal interface{}) {
	fmt.Println(strings.Repeat(decorate, 10))
	fmt.Printf("信号:%v \n ", signal)
	fmt.Println(strings.Repeat(decorate, 10))
}

func DefaultDebugDecorateShowSignal(signal interface{}) {
	fmt.Println(strings.Repeat("==", 20))
	fmt.Printf("信号:%v \n ", signal)
	fmt.Println(strings.Repeat("==", 20))
}

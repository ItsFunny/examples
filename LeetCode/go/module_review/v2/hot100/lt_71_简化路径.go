package hot100

import "strings"

// 关键:
// 第一想法: 用栈来处理, 结合strings的api
func simplifyPath(path string) string {
	stack := make([]string, 0)
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			// 则弹出上一个
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "" && v != "." {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

package hot100

import (
	"strconv"
	"strings"
)

// 关键
// 回溯算法 dfs
// 并且,注意 current ,当append 之后是不可以重新初始化的,因为后续的递归dfs 依赖了这个
func restoreIpAddresses(s string) []string {
	current:=make([]string,4)
	ret:=make([]string,0)
	var dfs func(index int,ipIndex int)

	dfs= func(index int,ipIndex int) {
		// dfs: 先考虑退出条件
		// 当当前长度为4的时候,并且 当前index 到了最后,则代表是一个结果集
		if ipIndex==4{
			if len(current)==4 && index== len(s){
				ret=append(ret,strings.Join(current,"."))
			}
			return
		}

		// 如果已经有了4元组,但是 index 还没到长度,直接return
		if index==len(s){
			return
		}
		if s[index]=='0'{
			current[ipIndex]="0"
			dfs(index+1,ipIndex+1)
		}
		// 开始给每个下标进行赋值
		add:=0
		for i:=index;i<len(s);i++{
			// add *10 是因为,每次在上一次移动到下一个的时候,都需要扩大10倍
			add=add*10+int(s[i]-'0')
			if add>0 && add<=255{
				current[ipIndex]=strconv.Itoa(add)
				dfs(i+1,ipIndex+1)
			}else{
				// 说明这个值已经不符合要求了
				break
			}

		}
	}
	dfs(0,0)

	return ret
}

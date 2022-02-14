package hot100


// 1. 从右往左和从上往下是需要单独的if判断的(并且是与条件,而不是或)
// 2. 注意if的边界条件,从左往右和从上往下都是<= 的边界条件,但是剩下的2个都是 <
// 3. 还是注意边界if
func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0{
		return nil
	}
	ret:=make([]int,0)

	left,right,top,bottom:=0, len(matrix[0])-1,0, len(matrix)-1

	for ;left<=right && top<=bottom;{
		// 开始遍历: 从左往右
		for i:=left;i<=right;i++{
			ret=append(ret,matrix[top][i])
		}

		// 继续遍历: 从上往下
		for i:=top+1;i<=bottom;i++{
			ret=append(ret,matrix[i][right])
		}

		// 最关键的是这一步,从右往左和从下往上是需要经过if判断
		if right>left && bottom>top{
			// 从右往左
			for i:=right-1;i>left;i--{
				ret=append(ret,matrix[bottom][i])
			}

			// 从下往上
			for i:=bottom;i>top;i--{
				ret=append(ret,matrix[i][left])
			}
		}
		left++
		top++
		right--
		bottom--
	}


	return  ret
}

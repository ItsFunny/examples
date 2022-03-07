package hot100


// 关键:
// 1. 二分法
// 当输入为8的时候,mid值为4 ,因为4*4=16 > 8 ,所以往右移动, 只有当 mid*mid< x 的时候,才代表这是一个可行的解
// 求中间值,然后不停的平方,
func mySqrt(x int) int {
	l,r:=0,x
	ret:=0
	for l<=r{
		mid:=l+(r-l)>>1
		if mid*mid<=x{
			ret=mid
			l=mid+1
		}else{
			r=mid-1
		}
	}
	return ret
}

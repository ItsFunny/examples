package hot100


// 关键:
// 递归判断,左右节点
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return sameTree(p,q)
}
func sameTree(l1,l2 *TreeNode)bool{
	if l1==nil && l2==nil{
		return true
	}
	if (l1==nil && l2!=nil) || (l2==nil && l1!=nil)|| (l1.Val!=l2.Val){
		return false
	}
	return sameTree(l1.Left,l2.Left) && sameTree(l1.Right,l2.Right)
}
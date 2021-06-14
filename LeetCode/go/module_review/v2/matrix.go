/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/4 9:13 上午
# @File : matrix.go
# @Description :
# @Attention :
*/
package v2

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 {
		return nil
	}
	// 多源 BFS ,0 先入队
	queue := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		x, y := q[0], q[1]
		for i := 0; i < 4; i++ {
			newX := q[0] + dx[i]
			newY := q[1] + dy[i]
			// 如果这个元素没有访问过,则标识为访问过
			// 因为取出来的元素,要么是之前为0 的元素,要么为 1然后被标识了最小距离的
			// 就是逐渐向外扩散,先把外层的1 的最小距离算出来
			if newX >= 0 && newX < len(mat) && newY >= 0 && newY < len(mat[q[0]]) && mat[newX][newY] == -1 {
				mat[newX][newY] = mat[x][y] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}

	return mat
}

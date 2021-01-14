/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 15:56 
# @File : heap_sort.go
# @Description : 
# @Attention : 
*/
package sort

func HeapSort(data []int) {
	for i := len(data) >> 1; i >= 0; i-- {
		buildHeap(data, i, len(data))
	}
	// 当最大堆构建完毕之后,我们只需要不断的将最大的移动到最后然后重新建堆即可
	// 并且因为最大的一直都是在0位,所以我们只需要从后往前更换元素即可
	for i := len(data) - 1; i >= 0; i-- {
		data[i], data[0] = data[0], data[i]
		buildHeap(data, 0, i)
	}
}

// 构建完全二叉树
func buildHeap(data []int, index, limit int) {
	leftChild := (index << 1) + 1
	rightChild := (index << 1) + 2
	// 因为是最大堆,所以要获取到最大的那个下标,进行重新构造树
	maxIndex := index
	if leftChild < limit && data[leftChild] > data[maxIndex] {
		maxIndex = leftChild
	}
	if rightChild < limit && data[rightChild] > data[maxIndex] {
		maxIndex = rightChild
	}
	if index == maxIndex {
		return
	}
	// 如果根节点不是最大值,则需要交换位置,重新构建子树
	data[maxIndex], data[index] = data[index], data[maxIndex]
	buildHeap(data, maxIndex, limit)
}

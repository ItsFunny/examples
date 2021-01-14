/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-16 09:45 
# @File : lt_1052_Grumpy_Bookstore_Owner.go
# @Description : 
# @Attention : 
*/
package slide_window

/*
	问题关键:
	连续的X分钟不会生气,而不是单独的n个
 */

func maxSatisfied(customers []int, grumpy []int, X int) int {
	res := 0
	for i := 0; i <len(customers); i++{
		if grumpy[i] == 0{
			res+= customers[i]
		}

		customers[i] *= grumpy[i]
	}

	max, sum := 0,0
	for i :=0; i < len(customers); i++{
		sum +=customers[i]
		if i >= X{
			sum -= customers[i-X]
		}

		max = maxF(max, sum)
	}

	return max+res
}

func maxF(a,b int) int{
	if a >b {
		return a
	}

	return b
}


//
// func maxSatisfied(customers []int, grumpy []int, X int) int {
// 	angryCustomers := make([]int, 0)
// 	result := 0
// 	for i := 0; i < len(customers); i++ {
// 		if grumpy[i] == 1 {
// 			angryCustomers = append(angryCustomers, customers[i])
// 		} else {
// 			result += customers[i]
// 			angryCustomers = append(angryCustomers, 0)
// 		}
// 	}
//
// 	max, left, right, count := 0, 0, 0, 0
// 	for ; right < len(grumpy); right++ {
// 		max += angryCustomers[right]
// 		if angryCustomers[right] > 0 {
// 			count++
// 		}
// 		if right-left+1> X{
// 			prev:=max-angryCustomers[left]
// 			max=maxSatisfiedMax(max,prev)
// 			left++
// 		}
// 	}
//
// 	return result + max
// }
//
// func maxSatisfiedMax(a,b int )int{
// 	if a>b{
// 		return a
// 	}
// 	return b
// }
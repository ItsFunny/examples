package com.basic.sort;

import lombok.Data;

/**
 * @author joker
 * @When
 * @Description 堆排序
 * @Detail
 * @date 创建时间：2019-01-17 22:01
 */

/*
    总结:
    注意点: 1.无论是建堆((length>>1)-1)还是排序(length-1),都是从大的往前推,因而是i-- ,而不是i++
           2.当为一个堆更换为(大根堆或者小根堆之后,子树叶要重新建堆),切记
           3.当建堆完毕排序的时候,从末尾开始固定与0更换位置即可
    其实核心就是不停的建堆:


    步骤:
        1. 建堆 ,如何建:从中(length/2-1)往前,左右孩子与根节点大小判断,交换,然后重复建堆,退出条件为根节点是最大(小)的值
        2. 排序 ,如何排:从尾往前,与0号元素交换位置,然后重新建堆
 */
public class HeapSort
{


    // 堆排序:
    // 分为2个步骤: 建堆,排序 (其实排序就是重新建堆的过程,因为构建的是堆,只需要将首位移到最后,剩余的继续建堆即可)
    // 堆排序是建立在完全二叉树上的,堆又分为最大堆和最小堆,如果我们想升序的话则需要构建的是最大堆,并且最大值在第一位
    public void heapSort(Integer[] arr)
    {
        // 建堆
        for (int i = (arr.length >> 1) - 1; i >= 0; i--)
        {
            // 为什么我们要从父节点的上限开始,原因在于当我们从后往前的时候,子节点所处的树较父节点肯定是小的
            // 也就可以省去很多无用的操作 ,想象一下就行,当根节点与某个子节点发生了变化之后,子节点需要重新排序
            // 这时候的树肯定是较当这个节点为父节点时候的子节点的树要大的
            buildHeap(arr, i, arr.length - 1);
        }
        // 当最大堆构建完毕之后,我们只需要不断的将最大的移动到最后然后重新建堆即可
        // 并且因为最大的一直都是在0位,所以我们只需要从后往前更换元素即可
        for (int i = arr.length - 1; i >= 0; i--)
        {
            int temp = arr[i];
            arr[i] = arr[0];
            arr[0] = temp;
            sort(arr, 0, i);
        }

    }

    // 建堆有2种方式,第一种是通过递归建堆的方式,如下:
    public void buildHeap(Integer[] arr, Integer rootIndex, Integer limitIndex)
    {
        // leftChild index
        Integer leftChildIndex = (rootIndex << 1) + 1;
        // rightChind index
        Integer rightChindIndex = (rootIndex << 1) + 2;
        Integer maxIndex = rootIndex;
        if (leftChildIndex < limitIndex && arr[leftChildIndex] > arr[maxIndex])
        {
            maxIndex = leftChildIndex;
        }
        if (rightChindIndex < limitIndex && arr[rightChindIndex] > arr[maxIndex])
        {
            maxIndex = rightChindIndex;
        }
        // 如果根节点不是最小值则交换位置
        if (maxIndex == rootIndex)
        {
            return;
        }
        Integer temp = arr[maxIndex];
        arr[maxIndex] = arr[rootIndex];
        arr[rootIndex] = temp;
        // 因为我们是对整个堆进行排序,所以当更换了值之后,所在的树也基本上变了,所以我们需要重新建堆
        // 既子节点的树很可能是发生了变化
        // 这里就是可以优化的地方,既然变更的只是子节点,大可抽出成为一个for循环实现
        buildHeap(arr, maxIndex, limitIndex);
    }

    public void buildHeapNoRecursion(Integer[] arr)
    {
        for(Integer i=(arr.length>>1)-1;i>=0;i--)
        {
            buildHeapWithOutRecursion(arr, i, arr.length-1);
        }

        // sort
        for (int i =arr.length-1; i >0 ; i--)
        {
            Integer temp=arr[0];
            arr[0]=arr[i];
            arr[i]=temp;
            buildHeapWithOutRecursion(arr, 0, i);
        }
    }

    // 非递归排序的思路与递归排序的思路是一样的;
    // 选取左右孩子的最大值,然后对交换位置的孩子作为根节点的树继续调整树
    public void buildHeapWithOutRecursion(Integer[] arr, Integer rootIndex, Integer limitIndex)
    {
        Integer temp=arr[rootIndex];
        for(Integer i=(rootIndex<<1)+1;i<limitIndex;i=(rootIndex<<1)+1)
        {
            // 右孩子就是+1所处的位置
            // 选取左右孩子的最大值与根节点进行比较
            if (i+1<=limitIndex&&arr[i+1]>arr[i])
            {
                i++;
            }
            // 这里可能会有疑问,为什么是不变的temp值去遍历比较底下的叶子节点的值:
            // 因为当我们不满足条件的时候(既根节点与孩子节点更换之后),此时孩子节点的root值就是temp了
            //
            //     5                    7
            //   3   7   -->         3      5    对   5  进行重新建树
            //  2 1 0 6           2    1  0  6       0 6
            //
            if (temp>=arr[i])
            {
                break;
            }
            // 否则就是孩子节点的值更加大,则需要更换位置,将孩子节点提到root节点上,
            // 然后对子节点所处的进行再建树(rootIndex=i 就使得这个孩子节点变成了根节点),对应上面的就是递归
            arr[rootIndex]=arr[i];
            rootIndex=i;
        }
        // 上面是直接复制的,最先的rootIndex节点的值就丢失了,因而我们需要将其补回,这里其实有点像插入排序
        arr[rootIndex]=temp;
    }


    public void sort(Integer[] arr, Integer index, Integer limit)
    {
        buildHeap(arr, index, limit);
    }


}

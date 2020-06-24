package com.basic.structure.tree;

import lombok.Data;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Stack;
import java.util.*;

/**
 * @author joker
 * @When
 * @Description 完全二叉树
 * @Detail
 * @date 创建时间：2019-01-16 08:57
 */
/*

 总结: 注意节点的判断是以数组的下标为0开始的

 完全二叉树的定义:
    1. 叶子节点都在左边
    2. 有右孩子必然有左孩子
    3. 按顺序满足满二叉树
    4. 树高度差为[0,1]
    5. 数组中之后[0,length>>1-1]
    6. 左孩子的节点为2*index+1   右孩子的节点为: 2*index+2        注意这里的index都是以0为起始的

 构建完全二叉树的注意点:
    1. 通过数组来构建,而数组的下标是从0开始

 判断是否是完全二叉树只需要抓住两个条件即可:
    1. 当左孩子存在的时候右孩子必然存在
    2. 所有的叶子节点都在[k-1,k]层
    实际上我们可以通过设定一个标志位,判断这个节点之前的那个节点是否是完整的,注意判断的是之前的节点,
    如果是完整的,而这个节点又有孩子则一定不是完全二叉树
 */
@Data
public class CompleteTree
{

    private TreeNode root;

    @Data
    private class TreeNode
    {
        private Integer data;
        private TreeNode leftChild;
        private TreeNode rightChild;

        public TreeNode(Integer data)
        {
            this.data = data;
        }
    }


    // 构建一颗完全二叉树
    // 构建树的时候是要注意考虑到数组是从0开始的
    // 构件完全二叉树的时候我们需要判断,左孩子节点 2*i+1 是否超过长度  <length(也可以是<=length-1 既最后一个下标),右孩子是否超过长度:2*i+2<length
    // 同时,当左孩子不存在的时候,右孩子就没必要判断了
    public void buildCompleteBinaryTree(Integer[] arr)
    {

        List<TreeNode> nodeList = new ArrayList<>();
        this.root = new TreeNode(arr[0]);
        nodeList.add(this.root);
        for (int i = 1; i < arr.length; i++)
        {
            nodeList.add(new TreeNode(arr[i]));
        }
        Integer length = (arr.length >> 1)-1;
        for (int i = 0; i <= length; i++)
        {
            if (i * 2 + 1 < arr.length)
            {
                nodeList.get(i).setLeftChild(nodeList.get(i * 2 + 1));
                if (i * 2 + 2 < arr.length)
                {
                    nodeList.get(i).setRightChild(nodeList.get(i * 2 + 2));
                }
            }

        }
    }

    public void buildTreeByStack(Integer[] arr)
    {
        if (arr.length == 0)
        {
            throw new RuntimeException("error");
        }
        Stack<TreeNode> stack = new Stack<>();
        this.root = new TreeNode(arr[0]);
        stack.push(this.root);
        boolean left = true;

        for (int i = 1; i < arr.length; i++)
        {

            if (arr[i] == -1)
            {
                if (left)
                {
                    left = false;

                } else if (!stack.isEmpty())
                {
                    stack.pop();
                }
                continue;
            }

            TreeNode newNode = new TreeNode(arr[i]);
            if (left)
            {
                stack.peek().leftChild = newNode;

            } else
            {
                left = true;
                TreeNode popNode = stack.pop();
                popNode.rightChild = newNode;
            }
            stack.push(newNode);
        }

    }


    // 注意点也是相同的,就是基于数组的特殊性,从1开始(从0开始会栈溢出),真正操作要减去1
    public void inIteratorByArray(Integer[] arr, Integer index, List<Integer> resultList)
    {
        if (index <= arr.length-1)
        {
            // ROOT
            resultList.add(arr[index]);
            // LEFT
            this.inIteratorByArray(arr, 2 * index + 1, resultList);
            // RIGHT
            this.inIteratorByArray(arr, 2 * index + 2, resultList);
        }
    }

    public List<Integer> BFSTree()
    {
        LinkedList<TreeNode> queue = new LinkedList<TreeNode>();
        List<Integer> resultList = new ArrayList<Integer>();
        queue.add(this.root);
        TreeNode temp = null;
        for (temp = queue.pop(); temp != null; temp = queue.pop())
        {
            resultList.add(temp.data);
            if (null != temp.leftChild)
            {
                queue.add(temp.leftChild);
            }
            if (null != temp.rightChild)
            {
                queue.add(temp.rightChild);
            }

            if (queue.isEmpty())
            {
                return resultList;
            }
        }
        return resultList;
    }


    //  注意点:我们需要通过BFS来遍历每一层
    //  完全二叉树的定义:
    //  叶子节点都在左边(既:有右节点必然有左节点)
    //  所有的叶子节点都在k层或者k-1层,也就是说[k-1,k]层的节点具有这样的特性:
    //  既如果存在左孩子不为空,右孩子为空,则这层这个节点之后的所有节点都必须为子节点才行
//    public boolean validIsCompleteTree(TreeNode root)
//    {
//
//        if (null == root)
//        {
//            return false;
//        }
//
//        LinkedList<TreeNode> queue = new LinkedList<>();
//        queue.add(root);
//        TreeNode temp = null;
//        for (temp = queue.pop(); temp != null; temp = queue.pop())
//        {
//
//            // left !=null && rith==null || left!=null && right!=null
//            if (temp.leftChild == null && temp.rightChild != null)
//            {
//                return false;
//                // leftChild!=null || rightChild==null
//            } else if (temp.rightChild == null)
//            {
//                // 确保这层的其他元素不是父节点
//                // 通过判断队列是否为空来判断是否需要进一步判断
//                while (!queue.isEmpty())
//                {
//                }
//                // 下面的做法是错误的,因为当深度为5的时候,第三层就是满的了,而此时是不仅仅只有2个节点的
////                if (queue.isEmpty())
////                {
////                    queue.push(temp.leftChild);
////                    continue; // 说明没有兄弟节点,或者说temp是最后一个兄弟节点
////                }
//                // 这里直接pop是没关系,因为如果条件满足,则直接退出了,如果不满足则为子节点,也没必要接着遍历
////                temp = queue.pop();
////                if (temp.leftChild != null || temp.rightChild != null)
////                {
////                    return false;
////                }
//
//            } else
//            {
//                queue.push(temp.leftChild);
//                queue.push(temp.rightChild);
//            }
//        }
//
//    }


    // 采取标志位的方式:
    // 如果一棵树只有左节点,则标志位为true,代表不完整
    // 若后续的树不完整(意味着left+right|left|right)则不是完全二叉树,既一旦有节点就不是完全二叉树
    // 非二叉树的条件:
    //  1. 当存在右孩子,左孩子却为空
    //  2. 当存在左孩子,右孩子为空,而同层的后续节点有孩子(既可以认为是之前的节点是不完整的)
    public boolean validIfCompleteTree()
    {
        if (this.root == null)
        {
            return false;
        }
        LinkedList<TreeNode> queue = new LinkedList<>();
        boolean previousCompleted = true;
        queue.add(this.root);
        for (TreeNode temp = queue.pop(); temp != null; temp = queue.pop())
        {
            // 1.对第一种条件的判断 和对第二种条件的判断
            if (temp.leftChild == null && temp.rightChild != null ||
                    !previousCompleted && (temp.leftChild != null || temp.rightChild != null))
            {
                return false;
            } else
            {

                // left=null &right=null
                // left!=null & right=null || left!=null & right!=null
                // previousCompleted
                if (temp.leftChild != null)
                {
                    queue.add(temp.leftChild);
                }
                if (temp.rightChild != null)
                {
                    previousCompleted = true;
                    queue.add(temp.rightChild);
                } else
                {
                    previousCompleted = false;
                }
            }
            if (queue.isEmpty())
            {
                break;
            }
        }
        return true;
    }


}

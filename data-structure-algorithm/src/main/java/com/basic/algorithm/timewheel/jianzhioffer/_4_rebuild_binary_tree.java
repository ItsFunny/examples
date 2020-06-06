package com.basic.algorithm.timewheel.jianzhioffer;

/**
 * @author joker
 * @When
 * @Description 根据先序和中序重新构建二叉树
 * 历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
 * @Detail          此题不会!!
 * @date 创建时间：2019-05-17 19:56
 */
public class _4_rebuild_binary_tree
{
    static class TreeNode
    {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x)
        {
            val = x;
        }

        private void interatorTree()
        {
        }
    }

    public static TreeNode rebuildTree(int[] pres, int[] ins)
    {

        return reConBTree(pres, 0, pres.length - 1, ins, 0, ins.length - 1);
    }

    public static TreeNode reConBTree(int[] pre, int preleft, int preright, int[] in, int inleft, int inright)
    {
        if (preleft > preright || inleft > inright)//当到达边界条件时候返回null
        {
            return null;
        }
        //新建一个TreeNode
        TreeNode pRootOfTree = new TreeNode(pre[preleft]);
        //对中序数组进行输入边界的遍历
        for (int i = inleft; i <= inright; i++)
        {
            if (pre[preleft] == in[i])
            {
                //重构左子树，注意边界条件
                pRootOfTree.left = reConBTree(pre, preleft + 1, preleft + i - inleft, in, inleft, i - 1);
                //重构右子树，注意边界条件
                pRootOfTree.right = reConBTree(pre, preleft + i + 1 - inleft, preright, in, i + 1, inright);
            }
        }
        return pRootOfTree;
    }

    public static void main(String[] args)
    {
        int[] pres = {1, 2, 4, 7, 3, 5, 6, 8};
        int[] ins = {4, 7, 2, 1, 5, 3, 8, 6};
        System.out.println(_4_rebuild_binary_tree.rebuildTree(pres, ins));

    }


}

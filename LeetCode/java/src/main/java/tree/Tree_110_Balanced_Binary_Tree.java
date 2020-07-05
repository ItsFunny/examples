package tree;

import common.TreeNode;

/**
 * @author Charlie
 * @When
 * @Description Given a binary tree, determine if it is height-balanced.
 * <p>
 * For this problem, a height-balanced binary tree is defined as:
 * <p>
 * a binary tree in which the left and right subtrees of
 * every node differ in height by no more than 1.
 * 判断是否是平衡二叉树
 * @Detail 是否是平衡二叉树的关键在于, 高度差最大为1
 * @Attention:
 * @Date 创建时间：2020-03-14 16:57
 */
public class Tree_110_Balanced_Binary_Tree
{
    public boolean isBalanced(TreeNode root)
    {
        if (root == null)
        {
            return true;
        }
        return helper(root) != -1;
    }

    private int helper(TreeNode root)
    {
        if (root == null)
        {
            return 0;
        }
        int left = helper(root.left);
        int right = helper(root.right);
        if (left == -1 || right == -1 || Math.abs(left - right) > 1)
        {
            return -1;
        }
        return Math.max(left, right) + 1;
    }

}

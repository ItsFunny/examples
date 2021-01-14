package string;

import common.TreeNode;

import java.util.concurrent.locks.AbstractQueuedSynchronizer;

/**
 * @author Charlie
 * @When
 * @Description You need to construct a string consists of
 * parenthesis and integers from a binary tree with the preorder traversing way.
 * <p>
 * The null node needs to be represented by empty parenthesis pair "()".
 * And you need to omit all the empty parenthesis pairs that don't affect
 * the one-to-one mapping relationship between the string and the original binary tree.
 * @Detail
 * 刚开始题意没有理解, 需要注意的就是当左右都为空的时候, 这个限制条件直接插入空的即可
 * 4种情况:
 * 1. 左右都为空,此时直接返回即可
 * 2. 左为空,右不为空,根据先序的规则: 根左右,所以需要 先加上()
 * 3. 左不为空,但是右为空: 则 根据题意,此时直接(+left+)
 * 4. 左右都不为空: 则 ( + left + ) + ( + right + )
 * @Attention:
 * @Date 创建时间：2020-03-13 16:02
 */
public class String_606_Construct_String_from_Binary_Tree
{
    public String tree2str(TreeNode t)
    {
        return preorderTree(t);
    }

    private String preorderTree(TreeNode t)
    {
        if (t == null)
        {
            return "";
        }
        StringBuilder sb = new StringBuilder();
        sb.append(t.val);
        if (t.left == null && t.right == null) return sb.toString();
        if (t.left != null && t.right != null)
        {
            sb.append('(');
            sb.append(preorderTree(t.left));
            sb.append(')');
            sb.append('(');
            sb.append(preorderTree(t.right));
            sb.append(')');
        } else if (t.left != null)
        {
            sb.append('(');
            sb.append(preorderTree(t.left));
            sb.append(')');
        } else
        {
            sb.append("()(");
            sb.append(preorderTree(t.right));
            sb.append(')');
        }

        return sb.toString();
    }

    public static void main(String[] args)
    {
        String_606_Construct_String_from_Binary_Tree s = new String_606_Construct_String_from_Binary_Tree();

    }
}
